package billing

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"github.com/pjover/sam/internal/domain/services/loader"
	"log"
	"sort"
	"strconv"
	"strings"
)

type billingService struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewBillingService(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) ports.BillingService {
	return billingService{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (b billingService) InsertConsumptions(childId int, consumptions map[string]float64, note string) (string, error) {
	return b.insertConsumptions(childId, consumptions, note, false)
}

func (b billingService) insertConsumptions(childId int, consumptions map[string]float64, note string, isRectification bool) (string, error) {
	var buffer bytes.Buffer

	child, err := b.dbService.FindChild(childId)
	if err != nil {
		return "", err
	}
	if !child.Active() {
		return "", fmt.Errorf("l'infant %s no està activat, edita'l per activar-lo abans d'insertar consums", child.NameWithId())
	}

	customerId := childId / 10
	customer, err := b.dbService.FindCustomer(customerId)
	if err != nil {
		return "", err
	}
	if !customer.Active() {
		return "", fmt.Errorf("el client %s no està activat, edita'l per activar-lo abans d'insertar consums", customer.FirstAdultNameWithId())
	}

	bulkLoader := loader.NewBulkLoader(b.configService, b.dbService)
	products, err := bulkLoader.LoadProducts()
	if err != nil {
		return "", err
	}

	completeConsumptions, err := b.completeConsumptions(consumptions, childId, note, isRectification, products)
	if err != nil {
		return "", err
	}

	err = b.dbService.InsertConsumptions(completeConsumptions)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	text, _ := model.ConsumptionListFormatValues(completeConsumptions, child, products, "")
	buffer.WriteString(text)

	return buffer.String(), nil
}

func (b billingService) completeConsumptions(consumptions map[string]float64, childId int, note string, isRectification bool, products map[string]model.Product) ([]model.Consumption, error) {
	yearMonth := b.configService.GetCurrentYearMonth()
	var first = true
	var completeConsumptions []model.Consumption
	for id, units := range consumptions {
		if err := b.checkIfProductExists(id, products); err != nil {
			return nil, err
		}

		var c model.Consumption
		if first {
			c = model.NewConsumption(
				common.RandString(model.ConsumptionIdLength),
				childId,
				id,
				units,
				yearMonth,
				note,
				isRectification,
				"NONE",
			)
			first = false
		} else {
			c = model.NewConsumption(
				common.RandString(model.ConsumptionIdLength),
				childId,
				id,
				units,
				yearMonth,
				"",
				isRectification,
				"NONE",
			)
		}
		completeConsumptions = append(completeConsumptions, c)
	}

	return completeConsumptions, nil
}

func (b billingService) checkIfProductExists(productId string, products map[string]model.Product) error {
	_, exists := products[productId]
	if !exists {
		return fmt.Errorf("el producte amb codi '%s' no existeix a la base de dades", productId)
	}
	return nil
}

func (b billingService) RectifyConsumptions(childId int, consumptions map[string]float64, note string) (string, error) {
	return b.insertConsumptions(childId, consumptions, note, true)
}

func (b billingService) BillConsumptions() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	consumptions, err := b.readConsumptionsFromDatabase()
	if err != nil {
		return "", err
	}
	if len(consumptions) == 0 {
		return "No hi han consums pendents de facturar", nil
	}

	invoices, customers, err := b.consumptionsToInvoices(consumptions)
	if err != nil {
		return "", err
	}

	newInvoices, sequences, err := b.addSequencesToInvoices(invoices, customers)
	if err != nil {
		return "", err
	}

	newConsumptions := b.addInvoiceIdToConsumptions(consumptions, newInvoices)

	err = b.updateDatabase(newConsumptions, newInvoices, sequences)

	return b.formatInvoicesGroupedByPaymentType(newInvoices, customers)
}

func (b billingService) readConsumptionsFromDatabase() ([]model.TransientConsumption, error) {
	consumptions, err := b.dbService.FindAllActiveConsumptions()
	if err != nil {
		return []model.TransientConsumption{}, err
	}
	var transientConsumptions []model.TransientConsumption
	for _, consumption := range consumptions {
		transientConsumption := model.TransientConsumption{
			Id:              consumption.Id(),
			ChildId:         consumption.ChildId(),
			ProductId:       consumption.ProductId(),
			Units:           consumption.Units(),
			YearMonth:       consumption.YearMonth(),
			Note:            consumption.Note(),
			IsRectification: consumption.IsRectification(),
		}
		transientConsumptions = append(transientConsumptions, transientConsumption)
	}
	return transientConsumptions, nil
}

func (b billingService) consumptionsToInvoices(consumptions []model.TransientConsumption) (invoices []model.TransientInvoice, customers map[string]model.Customer, err error) {
	groupedByCustomer := b.groupConsumptions(b.groupConsumptionsByCustomer, consumptions)
	customers = make(map[string]model.Customer)
	for customerIdStr, cons := range groupedByCustomer {
		customerId, _ := strconv.Atoi(customerIdStr)

		customer, err := b.dbService.FindCustomer(customerId)
		if err != nil {
			return nil, nil, err
		}

		consumptionsWithoutRectification, consumptionsWithRectification := b.splitConsumptions(cons)

		invoices, err = b.addInvoiceIfHasConsumptions(invoices, customer, consumptionsWithoutRectification, false)
		if err != nil {
			return nil, nil, err
		}

		invoices, err = b.addInvoiceIfHasConsumptions(invoices, customer, consumptionsWithRectification, true)
		if err != nil {
			return nil, nil, err
		}

		customers[customerIdStr] = customer

	}
	return invoices, customers, nil
}

func (b billingService) addInvoiceIfHasConsumptions(invoices []model.TransientInvoice, customer model.Customer, consumptions []model.TransientConsumption, isRectification bool) ([]model.TransientInvoice, error) {
	if len(consumptions) > 0 {
		invoice, err := b.consumptionsToInvoice(customer, consumptions, isRectification)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (b billingService) splitConsumptions(consumptions []model.TransientConsumption) (consumptionsWithoutRectification []model.TransientConsumption, consumptionsWithRectification []model.TransientConsumption) {
	for _, consumption := range consumptions {
		if consumption.IsRectification {
			consumptionsWithRectification = append(consumptionsWithRectification, consumption)
		} else {
			consumptionsWithoutRectification = append(consumptionsWithoutRectification, consumption)
		}
	}
	return consumptionsWithoutRectification, consumptionsWithRectification
}

func (b billingService) consumptionsToInvoice(customer model.Customer, consumptions []model.TransientConsumption, isRectification bool) (model.TransientInvoice, error) {
	yearMonth := b.configService.GetCurrentYearMonth()
	today := b.osService.Now()

	lines, childrenIds, err := b.childrenLines(consumptions)
	if err != nil {
		return model.TransientInvoice{}, err
	}

	return model.TransientInvoice{
		IsRectification: isRectification,
		CustomerId:      customer.Id(),
		Date:            today,
		YearMonth:       yearMonth,
		ChildrenIds:     childrenIds,
		Lines:           lines,
		PaymentType:     customer.InvoiceHolder().PaymentType(),
		Note:            b.notes(consumptions),
	}, nil
}

func (b billingService) childrenLines(consumptions []model.TransientConsumption) (lines []model.InvoiceLine, childrenIds []int, err error) {
	groupedByChild := b.groupConsumptions(b.groupConsumptionsByChild, consumptions)
	for childId, cons := range groupedByChild {
		cid, _ := strconv.Atoi(childId)
		childrenIds = append(childrenIds, cid)
		productLines, err := b.productLines(cons)
		if err != nil {
			return nil, nil, err
		}
		lines = append(lines, productLines...)
	}
	return lines, childrenIds, nil
}

func (b billingService) productLines(consumptions []model.TransientConsumption) (lines []model.InvoiceLine, err error) {
	groupedByProduct := b.groupConsumptions(b.groupConsumptionsByProduct, consumptions)
	for productId, cons := range groupedByProduct {
		product, err := b.dbService.FindProduct(productId)
		if err != nil {
			return nil, err
		}

		var units float64
		for _, con := range cons {
			units += con.Units
		}
		if units == 0 {
			continue
		}

		line := model.NewInvoiceLine(
			productId,
			units,
			product.Price(),
			product.TaxPercentage(),
			cons[0].ChildId,
		)
		lines = append(lines, line)
	}
	return lines, nil
}

func (b billingService) notes(consumptions []model.TransientConsumption) string {
	var notes []string
	for _, consumption := range consumptions {
		if consumption.Note == "" {
			continue
		}
		notes = append(notes, consumption.Note)
	}
	return strings.Join(notes, ", ")
}

func (b billingService) groupConsumptions(groupBy func(consumption model.TransientConsumption) string, consumptions []model.TransientConsumption) map[string][]model.TransientConsumption {
	var auxMap = make(map[string][]model.TransientConsumption)
	for _, consumption := range consumptions {
		var group = groupBy(consumption)
		groupedConsumptions := auxMap[group]
		groupedConsumptions = append(groupedConsumptions, consumption)
		auxMap[group] = groupedConsumptions
	}
	return auxMap
}

func (b billingService) groupConsumptionsByCustomer(consumption model.TransientConsumption) string {
	return strconv.Itoa(consumption.ChildId / 10)
}

func (b billingService) groupConsumptionsByChild(consumption model.TransientConsumption) string {
	return strconv.Itoa(consumption.ChildId)
}

func (b billingService) groupConsumptionsByProduct(consumption model.TransientConsumption) string {
	return consumption.ProductId
}

func (b billingService) addSequencesToInvoices(invoices []model.TransientInvoice, customers map[string]model.Customer) ([]model.Invoice, []model.Sequence, error) {
	sequencesMap, err := b.getSequences()
	if err != nil {
		return nil, nil, err
	}

	sort.Slice(invoices, func(i, j int) bool {
		return invoices[i].CustomerId < invoices[j].CustomerId
	})

	var outInvoices []model.Invoice
	for _, invoice := range invoices {
		customerIdStr := strconv.Itoa(invoice.CustomerId)
		customer := customers[customerIdStr]
		sequenceType := b.getSequenceType(invoice, customer)
		sequence := sequencesMap[sequenceType.Format()]
		newSequence := model.NewSequence(sequenceType, sequence.Counter()+1)
		newInvoiceId := fmt.Sprintf("%s-%d", newSequence.Id().Prefix(), newSequence.Counter())
		newInvoice := b.newInvoice(invoice, newInvoiceId)
		outInvoices = append(outInvoices, newInvoice)
		sequencesMap[sequenceType.Format()] = newSequence
	}

	var outSequences []model.Sequence
	for _, sequence := range sequencesMap {
		outSequences = append(outSequences, sequence)
	}
	return outInvoices, outSequences, nil
}

func (b billingService) newInvoice(invoice model.TransientInvoice, id string) model.Invoice {
	return model.NewInvoice(
		id,
		invoice.CustomerId,
		invoice.Date,
		invoice.YearMonth,
		invoice.ChildrenIds,
		invoice.Lines,
		invoice.PaymentType,
		invoice.Note,
		false,
		false,
		false,
	)
}

func (b billingService) getSequenceType(invoice model.TransientInvoice, customer model.Customer) sequence_type.SequenceType {
	if invoice.IsRectification {
		return sequence_type.RectificationInvoice
	} else {
		return customer.InvoiceHolder().PaymentType().SequenceType()
	}
}

func (b billingService) getSequences() (sequences map[string]model.Sequence, err error) {
	allSequences, err := b.dbService.FindAllSequences()
	if err != nil {
		return nil, err
	}

	sequences = make(map[string]model.Sequence)
	for _, sequence := range allSequences {
		sequences[sequence.Id().Format()] = sequence
	}
	return sequences, nil
}

func (b billingService) groupInvoices(groupBy func(invoice model.Invoice) string, invoices []model.Invoice) map[string][]model.Invoice {
	var auxMap = make(map[string][]model.Invoice)
	for _, invoice := range invoices {
		var group = groupBy(invoice)
		groupedInvoices := auxMap[group]
		groupedInvoices = append(groupedInvoices, invoice)
		auxMap[group] = groupedInvoices
	}
	return auxMap
}

func (b billingService) groupInvoicesByPaymentType(invoice model.Invoice) string {
	return invoice.PaymentType().Format()
}

func (b billingService) groupInvoicesByCustomer(invoice model.Invoice) string {
	return strconv.Itoa(invoice.CustomerId())
}

func (b billingService) addInvoiceIdToConsumptions(consumptions []model.TransientConsumption, invoices []model.Invoice) []model.Consumption {
	var outConsumptions []model.Consumption
	for _, consumption := range consumptions {
		invoiceId := b.findInvoiceId(consumption, invoices)
		if invoiceId == "" {
			log.Fatalf(
				"no s'ha trobat cap factura per al consum %d  %s  %4.1f  %s  %s",
				consumption.ChildId,
				consumption.YearMonth,
				consumption.Units,
				consumption.ProductId,
				consumption.Note,
			)
		}
		newConsumption := model.NewConsumption(
			consumption.Id,
			consumption.ChildId,
			consumption.ProductId,
			consumption.Units,
			consumption.YearMonth,
			consumption.Note,
			consumption.IsRectification,
			invoiceId,
		)
		outConsumptions = append(outConsumptions, newConsumption)
	}
	return outConsumptions
}

func (b billingService) findInvoiceId(consumption model.TransientConsumption, invoices []model.Invoice) string {
	customerId := consumption.ChildId / 10
	for _, invoice := range invoices {
		if invoice.CustomerId() == customerId {
			return invoice.Id()
		}
	}
	return ""
}

func (b billingService) formatInvoicesGroupedByPaymentType(invoices []model.Invoice, customers map[string]model.Customer) (string, error) {
	var buffer bytes.Buffer
	total := 0.0
	for paymentType, groupedInvoices := range b.groupInvoices(b.groupInvoicesByPaymentType, invoices) {
		subtotal := 0.0
		sort.Slice(groupedInvoices, func(i, j int) bool {
			return groupedInvoices[i].CustomerId() > groupedInvoices[j].CustomerId()
		})
		for i, invoice := range groupedInvoices {
			customerId := strconv.Itoa(invoice.CustomerId())
			customer := customers[customerId]
			buffer.WriteString(fmt.Sprintf(" %d. %s %s\n", i+1, customer.FirstAdultName(), invoice.String()))
			subtotal += invoice.Amount()
		}
		total += subtotal
		buffer.WriteString(fmt.Sprintf("Total %d %s: %.02f €\n", len(groupedInvoices), paymentType, subtotal))
	}
	buffer.WriteString(fmt.Sprintf("TOTAL: %.02f €\n", total))
	return buffer.String(), nil
}

func (b billingService) updateDatabase(consumptions []model.Consumption, invoices []model.Invoice, sequences []model.Sequence) error {

	err := b.dbService.InsertInvoices(invoices)
	if err != nil {
		return err
	}

	err = b.dbService.UpdateConsumptions(consumptions)
	if err != nil {
		// TODO Delete recently inserted invoices
		return err
	}

	err = b.dbService.UpdateSequences(sequences)
	if err != nil {
		// TODO Delete recently inserted invoices
		// TODO Remove invoiceId from recently updated consumptions
		return err
	}
	return nil
}

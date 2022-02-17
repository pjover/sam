package billing

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

type BillingService interface {
	InsertConsumptions(id int, consumptions map[string]float64, note string) (string, error)
	BillConsumptions() (string, error)
}

type billingService struct {
	cfgService  ports.ConfigService
	dbService   ports.DbService
	osService   ports.OsService
	postManager hobbit.HttpPostManager
}

func NewBillingService(cfgService ports.ConfigService, dbService ports.DbService, osService ports.OsService, postManager hobbit.HttpPostManager) BillingService {
	return billingService{
		cfgService:  cfgService,
		dbService:   dbService,
		osService:   osService,
		postManager: postManager,
	}
}

func (b billingService) InsertConsumptions(childId int, consumptions map[string]float64, note string) (string, error) {
	var buffer bytes.Buffer

	child, err := b.dbService.FindChild(childId)
	if err != nil {
		return "", err
	}
	if !child.Active {
		return "", fmt.Errorf("l'infant %s no està activat, edita'l per activar-lo abans d'insertar consums", child.NameWithId())
	}

	customerId := childId / 10
	customer, err := b.dbService.FindCustomer(customerId)
	if err != nil {
		return "", err
	}
	if !customer.Active {
		return "", fmt.Errorf("el client %s no està activat, edita'l per activar-lo abans d'insertar consums", customer.FirstAdultNameWithId())
	}

	products, err := b.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	yearMonth := viper.GetString("yearMonth")
	var first = true
	var completeConsumptions []model.Consumption
	for id, units := range consumptions {
		c := model.Consumption{
			Id:              common.RandString(model.ConsumptionIdLength),
			ChildId:         childId,
			ProductId:       id,
			Units:           units,
			YearMonth:       yearMonth,
			IsRectification: false,
			InvoiceId:       "NONE",
		}
		if first {
			c.Note = note
			first = false
		}
		completeConsumptions = append(completeConsumptions, c)
	}

	err = b.dbService.InsertConsumptions(completeConsumptions)
	if err != nil {
		return "", err
	}

	buffer.WriteString(model.ConsumptionListToString(completeConsumptions, child, products))

	return buffer.String(), nil
}

func (b billingService) BillConsumptions() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	//consumptions, err := b.dbService.FindAllActiveConsumptions()
	//if err != nil {
	//	return "", err
	//}
	//
	//invoices, customers, err := b.consumptionsToInvoices(consumptions)
	//if err != nil {
	//	return "", err
	//}

	// TODO Build invoices sequences

	// TODO Save invoices

	// TODO Display invoices grouped by PaymentType with totals

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.postManager.PrettyJson(url, data)
}

func (b billingService) consumptionsToInvoices(consumptions []model.Consumption) (invoices []model.Invoice, customers []model.Customer, err error) {
	groupedByCustomer := b.groupConsumptionsByCustomer(consumptions)
	for customerId, cons := range groupedByCustomer {
		cid, _ := strconv.Atoi(customerId)

		customer, err := b.dbService.FindCustomer(cid)
		if err != nil {
			return nil, nil, err
		}

		invoice, err := b.consumptionsToInvoice(customer, cons)
		if err != nil {
			return nil, nil, err
		}
		customers = append(customers, customer)
		invoices = append(invoices, invoice)
	}
	return invoices, customers, nil
}

func (b billingService) groupConsumptionsByCustomer(consumptions []model.Consumption) map[string][]model.Consumption {
	return b.groupConsumptions(b.groupByCustomer, consumptions)
}

func (b billingService) consumptionsToInvoice(customer model.Customer, consumptions []model.Consumption) (model.Invoice, error) {
	yearMonth := b.cfgService.GetString("yearMonth")
	today := b.osService.Now()

	lines, childrenIds, err := b.childrenLines(consumptions)
	if err != nil {
		return model.Invoice{}, err
	}

	return model.Invoice{
		CustomerId:  customer.Id,
		Date:        today,
		YearMonth:   yearMonth,
		ChildrenIds: childrenIds,
		Lines:       lines,
		PaymentType: customer.InvoiceHolder.PaymentType,
		Note:        b.notes(consumptions),
	}, nil
}

func (b billingService) childrenLines(consumptions []model.Consumption) (lines []model.Line, childrenIds []int, err error) {
	groupedByChild := b.groupConsumptions(b.groupByChild, consumptions)
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

func (b billingService) productLines(consumptions []model.Consumption) (lines []model.Line, err error) {
	groupedByProduct := b.groupConsumptions(b.groupByProduct, consumptions)
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

		line := model.Line{
			ProductId:     productId,
			Units:         units,
			ProductPrice:  product.Price,
			TaxPercentage: product.TaxPercentage,
			ChildId:       cons[0].ChildId,
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func (b billingService) notes(consumptions []model.Consumption) string {
	var notes []string
	for _, consumption := range consumptions {
		if consumption.Note == "" {
			continue
		}
		notes = append(notes, consumption.Note)
	}
	return strings.Join(notes, ", ")
}

func (b billingService) groupConsumptions(groupBy func(consumption model.Consumption) string, consumptions []model.Consumption) map[string][]model.Consumption {
	var auxMap = make(map[string][]model.Consumption)
	for _, con := range consumptions {
		var group = groupBy(con)
		cons := auxMap[group]
		cons = append(cons, con)
		auxMap[group] = cons
	}
	return auxMap
}

func (b billingService) groupByCustomer(consumption model.Consumption) string {
	return strconv.Itoa(consumption.ChildId / 10)
}

func (b billingService) groupByChild(consumption model.Consumption) string {
	return strconv.Itoa(consumption.ChildId)
}

func (b billingService) groupByProduct(consumption model.Consumption) string {
	return consumption.ProductId
}

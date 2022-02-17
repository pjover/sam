package billing

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"github.com/spf13/viper"
)

type BillingService interface {
	InsertConsumptions(Code int, Consumptions map[string]float64, Note string) (string, error)
	BillConsumptions() (string, error)
}

type billingService struct {
	dbService   ports.DbService
	osService   ports.OsService
	postManager hobbit.HttpPostManager
}

func NewBillingService(dbService ports.DbService, osService ports.OsService, postManager hobbit.HttpPostManager) BillingService {
	return billingService{
		dbService:   dbService,
		osService:   osService,
		postManager: postManager,
	}
}

func (b billingService) InsertConsumptions(childCode int, consumptions map[string]float64, note string) (string, error) {
	var buffer bytes.Buffer

	child, err := b.dbService.FindChild(childCode)
	if err != nil {
		return "", err
	}
	if !child.Active {
		return "", fmt.Errorf("l'infant %s no està activat, edita'l per activar-lo abans d'insertar consums", child.NameWithCode())
	}

	customerCode := childCode / 10
	customer, err := b.dbService.FindCustomer(customerCode)
	if err != nil {
		return "", err
	}
	if !customer.Active {
		return "", fmt.Errorf("el client %s no està activat, edita'l per activar-lo abans d'insertar consums", customer.FirstAdultNameWithCode())
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
			Code:            common.RandString(model.ConsumptionCodeLength),
			ChildCode:       childCode,
			ProductId:       id,
			Units:           units,
			YearMonth:       yearMonth,
			IsRectification: false,
			InvoiceCode:     "NONE",
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

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.postManager.PrettyJson(url, data)
}

func (b billingService) consumptionsToInvoices(consumptions []model.Consumption) []model.Invoice {

	var invoices []model.Invoice
	groupedByCustomer := b.groupConsumptionsByCustomer(consumptions)
	for customerCode, cons := range groupedByCustomer {
		invoice := b.consumptionsToInvoice(customerCode, cons)
		invoices = append(invoices, invoice)
	}
	return invoices
}

func (b billingService) groupConsumptionsByCustomer(consumptions []model.Consumption) map[int][]model.Consumption {
	var auxMap = make(map[int][]model.Consumption)
	for _, con := range consumptions {
		var customerCode = con.ChildCode / 10
		cons := auxMap[customerCode]
		cons = append(cons, con)
		auxMap[customerCode] = cons
	}
	return auxMap
}

func (b billingService) consumptionsToInvoice(customerCode int, consumptions []model.Consumption) model.Invoice {

	return model.Invoice{
		CustomerId: customerCode,
	}
}

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
	InsertConsumptions(id int, consumptions map[string]float64, note string) (string, error)
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

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.postManager.PrettyJson(url, data)
}

func (b billingService) consumptionsToInvoices(consumptions []model.Consumption) []model.Invoice {

	var invoices []model.Invoice
	groupedByCustomer := b.groupConsumptionsByCustomer(consumptions)
	for customerId, cons := range groupedByCustomer {
		invoice := b.consumptionsToInvoice(customerId, cons)
		invoices = append(invoices, invoice)
	}
	return invoices
}

func (b billingService) groupConsumptionsByCustomer(consumptions []model.Consumption) map[int][]model.Consumption {
	var auxMap = make(map[int][]model.Consumption)
	for _, con := range consumptions {
		var customerId = con.ChildId / 10
		cons := auxMap[customerId]
		cons = append(cons, con)
		auxMap[customerId] = cons
	}
	return auxMap
}

func (b billingService) consumptionsToInvoice(customerId int, consumptions []model.Consumption) model.Invoice {

	return model.Invoice{
		CustomerId: customerId,
	}
}

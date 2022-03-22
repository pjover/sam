package bdd

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"time"
)

type InvoiceToBddConverter interface {
	ToBdd(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd
}

type invoiceToBddConverter struct {
	configService ports.ConfigService
	osService     ports.OsService
}

func NewInvoiceToBddConverter(configService ports.ConfigService, osService ports.OsService) InvoiceToBddConverter {
	return invoiceToBddConverter{
		configService: configService,
		osService:     osService,
	}
}

func (i invoiceToBddConverter) ToBdd(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd {
	now := i.osService.Now()
	return Bdd{
		messageIdentification:   i.getMessageIdentification(now),
		creationDateTime:        i.getCreationDateTime(now),
		numberOfTransactions:    i.getNumberOfTransactions(invoices),
		controlSum:              i.getControlSum(invoices),
		name:                    i.configService.GetString("business.name"),
		identification:          i.configService.GetString("bdd.id"),
		requestedCollectionDate: i.getRequestedCollectionDate(now),
		country:                 i.configService.GetString("bdd.country"),
		addressLine1:            i.configService.GetString("business.addressLine1"),
		addressLine2:            i.configService.GetString("business.addressLine2"),
		iban:                    i.configService.GetString("bdd.iban"),
		bic:                     i.configService.GetString("bdd.bankBic"),
		details:                 i.getDetails(now, invoices, customers, products),
	}
}

func (i invoiceToBddConverter) getMessageIdentification(now time.Time) string {
	bddPrefix := i.configService.GetString("bdd.prefix")
	datetime := now.Format("20060102150405000")
	suffix := "" // TODO WIP
	return fmt.Sprintf("%s-%s-%s", bddPrefix, datetime, suffix)
}

func (i invoiceToBddConverter) getCreationDateTime(now time.Time) string {
	return "" // TODO WIP
}

func (i invoiceToBddConverter) getNumberOfTransactions(invoices []model.Invoice) int {
	return 0 // TODO WIP
}

func (i invoiceToBddConverter) getControlSum(invoices []model.Invoice) string {
	return "" // TODO WIP
}

func (i invoiceToBddConverter) getRequestedCollectionDate(now time.Time) string {
	return "" // TODO WIP
}

func (i invoiceToBddConverter) getDetails(now time.Time, invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) []BddDetail {
	return nil // TODO WIP
}

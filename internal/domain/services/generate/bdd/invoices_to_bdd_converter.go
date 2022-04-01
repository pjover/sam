package bdd

import (
	"fmt"
	"github.com/Masterminds/goutils"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"strconv"
	"strings"
	"time"
)

type InvoicesToBddConverter interface {
	Convert(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd
}

type invoicesToBddConverter struct {
	configService ports.ConfigService
	osService     ports.OsService
}

func NewInvoicesToBddConverter(configService ports.ConfigService, osService ports.OsService) InvoicesToBddConverter {
	return invoicesToBddConverter{
		configService: configService,
		osService:     osService,
	}
}

func (i invoicesToBddConverter) Convert(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd {
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

func (i invoicesToBddConverter) getMessageIdentification(now time.Time) string {
	bddPrefix := i.configService.GetString("bdd.prefix")
	datetime := now.Format("20060102150405000")
	checkDigits := common.NewMod9710(bddPrefix, datetime).Checksum()
	return fmt.Sprintf("%s-%s-%s", bddPrefix, datetime, checkDigits)
}

func (i invoicesToBddConverter) getCreationDateTime(now time.Time) string {
	return now.Format("2006-01-02T15:04:05")
}

func (i invoicesToBddConverter) getNumberOfTransactions(invoices []model.Invoice) int {
	return len(invoices)
}

func (i invoicesToBddConverter) getControlSum(invoices []model.Invoice) string {
	var controlSum float64
	for _, invoice := range invoices {
		controlSum += invoice.Amount()
	}
	return fmt.Sprintf("%.02f", controlSum)
}

func (i invoicesToBddConverter) getRequestedCollectionDate(now time.Time) string {
	return now.Format(domain.YearMonthDayLayout)
}

func (i invoicesToBddConverter) getDetails(now time.Time, invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) []BddDetail {
	var details []BddDetail
	for _, invoice := range invoices {
		detail := i.getDetail(now, invoice, customers[invoice.CustomerId], products)
		details = append(details, detail)
	}
	return details
}

func (i invoicesToBddConverter) getDetail(now time.Time, invoice model.Invoice, customer model.Customer, products map[string]model.Product) BddDetail {
	return BddDetail{
		endToEndIdentifier:    i.getDetailEndToEndIdentifier(now, invoice),
		instructedAmount:      i.getDetailInstructedAmount(invoice),
		dateOfSignature:       i.getDetailDateOfSignature(now),
		name:                  i.getDetailName(customer),
		identification:        i.getDetailIdentification(customer),
		iban:                  i.getDetailCustomerBankAccount(customer),
		purposeCode:           i.configService.GetString("bdd.purposeCode"),
		remittanceInformation: i.getDetailRemittanceInformation(invoice, products),
		isBusiness:            i.getDetailIsBusiness(customer),
	}
}

func (i invoicesToBddConverter) getDetailEndToEndIdentifier(now time.Time, invoice model.Invoice) string {
	return fmt.Sprintf("%s.%s", i.getMessageIdentification(now), invoice.Id)
}

func (i invoicesToBddConverter) getDetailInstructedAmount(invoice model.Invoice) string {
	return fmt.Sprintf("%.02f", invoice.Amount())
}

func (i invoicesToBddConverter) getDetailDateOfSignature(now time.Time) string {
	return now.Format(domain.YearMonthDayLayout)
}

func (i invoicesToBddConverter) getDetailName(customer model.Customer) string {
	return customer.InvoiceHolder.Name
}

func (i invoicesToBddConverter) getDetailIdentification(customer model.Customer) string {
	country := i.configService.GetString("bdd.country")
	return i.getSepaIndentifier(customer.InvoiceHolder.TaxID.String(), country, "000")
}

func (i invoicesToBddConverter) getDetailCustomerBankAccount(customer model.Customer) string {
	return customer.InvoiceHolder.Iban.String()
}

func (i invoicesToBddConverter) getDetailRemittanceInformation(invoice model.Invoice, products map[string]model.Product) string {
	const maxLength = 140
	invoiceDescription := i.getShortNameInvoiceDescription(invoice, products)
	info, err := goutils.Abbreviate(invoiceDescription, maxLength)
	if err != nil {
		return invoiceDescription[0:maxLength]
	}
	return info
}

func (i invoicesToBddConverter) getShortNameInvoiceDescription(invoice model.Invoice, products map[string]model.Product) string {
	var lines []string
	for _, line := range invoice.Lines {
		units := strconv.FormatFloat(line.Units, 'f', -1, 64)
		desc := fmt.Sprintf("%sx%s", units, products[line.ProductId].ShortName)
		lines = append(lines, desc)
	}
	return strings.Join(lines, ", ")
}

func (i invoicesToBddConverter) getDetailIsBusiness(customer model.Customer) bool {
	return customer.InvoiceHolder.IsBusiness
}

func (i invoicesToBddConverter) getSepaIndentifier(taxID string, country string, suffix string) string {
	return fmt.Sprintf("%s%s%03s%09s",
		strings.ToUpper(country),
		common.NewMod9710(taxID, country).Checksum(),
		suffix,
		taxID,
	)
}

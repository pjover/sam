package bdd

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/goutils"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type InvoicesToBddConverter interface {
	Run(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd
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

func (i invoicesToBddConverter) Run(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) Bdd {
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
	suffix := i.calculateControlCode(bddPrefix, datetime)
	return fmt.Sprintf("%s-%s-%s", bddPrefix, datetime, suffix)
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
	return now.Format("2006-01-02")
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
	return now.Format("2006-01-02")
}

func (i invoicesToBddConverter) getDetailName(customer model.Customer) string {
	return customer.InvoiceHolder.Name
}

func (i invoicesToBddConverter) getDetailIdentification(customer model.Customer) string {
	country := i.configService.GetString("bdd.country")
	return i.getSepaIndentifier(customer.InvoiceHolder.TaxID, country, "000")
}

func (i invoicesToBddConverter) getDetailCustomerBankAccount(customer model.Customer) string {
	return customer.InvoiceHolder.BankAccount
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
		desc := fmt.Sprintf("%.1f x %s", line.Units, products[line.ProductId].ShortName)
		lines = append(lines, desc)
	}
	return strings.Join(lines, ", ")
}

func (i invoicesToBddConverter) getDetailIsBusiness(customer model.Customer) bool {
	return customer.InvoiceHolder.IsBusiness
}

func (i invoicesToBddConverter) getSepaIndentifier(taxID string, country string, suffix string) string {
	return fmt.Sprintf("%s%03s%09s",
		i.calculateControlCode(taxID, country),
		suffix,
		taxID,
	)
}

func (i invoicesToBddConverter) calculateControlCode(params ...string) string {
	preparedParams := i.prepareParams(params...)
	assignedWeightsToLetters := i.assignWeightsToLetters(preparedParams)
	appliedModel := i.apply9710Model(assignedWeightsToLetters)
	//return i.apply9710Model(i.assignWeightsToLetters(i.prepareParams(params...)))
	return appliedModel
}

func (i invoicesToBddConverter) prepareParams(params ...string) string {
	rawCode := strings.Join(params, "")
	preparedParam := i.prepareParam(rawCode)
	return preparedParam
}

func (i invoicesToBddConverter) prepareParam(rawCode string) string {
	var param string
	if rawCode != "" {
		param = strings.ReplaceAll(rawCode, " ", "")
		param = strings.ReplaceAll(param, "-", "")
	}
	result := fmt.Sprintf("%s00", param)
	return result
}

func (i invoicesToBddConverter) assignWeightsToLetters(code string) string {
	runes := []rune(code)
	var buffer bytes.Buffer
	for _, r := range runes {
		stringValue := string(r)
		log.Println(stringValue)
		intValue := int(r)
		var v int
		if r >= 'A' {
			v = intValue - 'A' + 10
		} else {
			v = intValue - '0'
		}

		buffer.WriteString(strconv.Itoa(v))
	}
	result := buffer.String()
	return result
}

func (i invoicesToBddConverter) apply9710Model(input string) string {
	//Applies the 97-10 model according to ISO-7604 (http://is.gd/9HE1zs)
	in := new(big.Int)
	in, ok := in.SetString(input, 10)
	if !ok {
		log.Fatalf("cannot convert %s to big integer", input)
	}

	mod := new(big.Int).Mod(in, big.NewInt(97))

	id3 := mod.Int64()

	id := 98 - id3
	result := fmt.Sprintf("%02d", id)
	return result
}

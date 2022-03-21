package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
	"strconv"
)

type invoiceReport struct {
	configService ports.ConfigService
	dbService     ports.DbService
}

func NewInvoiceReport(configService ports.ConfigService, dbService ports.DbService) invoiceReport {
	return invoiceReport{
		configService: configService,
		dbService:     dbService,
	}
}

func (i invoiceReport) Run(id string) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant l'informe de la factura %s ...\n", id))

	invoice, err := i.dbService.FindInvoice(id)
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut la factura %s: %s", id, err)
	}

	customer, err := i.dbService.FindCustomer(invoice.CustomerId)
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut trobar el client %d de la factura %s: %s", invoice.CustomerId, id, err)
	}

	report := Report{
		PageOrientation: consts.Portrait,
		Title:           fmt.Sprintf("Factura %s", id),
		SubReports: []SubReport{
			CardSubReport{
				Title: "Title 1",
				Align: consts.Left,
				Captions: []string{
					"Factura",
					"Data",
					"NIF/CIF",
					"Client",
					"Adreça",
					"Codi",
					"Infants",
				},
				Widths: []uint{
					4,
					18,
				},
				Data: i.buildInvoiceHeaderData(invoice, customer),
			},
			CardSubReport{
				Title: "Title 2",
				Align: consts.Left,
				Captions: []string{
					"Factura",
					"Data",
					"NIF/CIF",
					"Client",
					"Adreça",
					"Codi",
					"Infants",
				},
				Widths: []uint{
					4,
					18,
				},
				Data: i.buildInvoiceHeaderData(invoice, customer),
			},
		},
	}

	dirPath, err := i.configService.GetInvoicesDirectory()
	if err != nil {
		return "", err
	}
	filePath := path.Join(
		dirPath,
		fmt.Sprintf("%s (%d).pdf", invoice.Id, invoice.CustomerId),
	)
	err = report.SaveToFile(filePath)
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("Generat l'informe de productes a '%s'\n", filePath))
	return buffer.String(), nil

}

func (i invoiceReport) buildInvoiceHeaderData(invoice model.Invoice, customer model.Customer) [][]string {

	var data = [][]string{
		{invoice.Id},
		{invoice.DateFmt()},
		{customer.InvoiceHolder.TaxID},
		{customer.InvoiceHolder.Name},
		{customer.InvoiceHolder.Address.CompleteAddress()},
		{strconv.Itoa(customer.Id)},
		{customer.ChildrenNames(",")},
	}
	return data
}

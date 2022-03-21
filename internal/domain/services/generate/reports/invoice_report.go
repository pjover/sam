package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
	"sort"
	"strconv"
	"strings"
)

type InvoiceReport struct {
	configService ports.ConfigService
	dbService     ports.DbService
}

func NewInvoiceReport(configService ports.ConfigService, dbService ports.DbService) InvoiceReport {
	return InvoiceReport{
		configService: configService,
		dbService:     dbService,
	}
}

func (i InvoiceReport) SingleInvoice(id string) (string, error) {

	invoice, err := i.dbService.FindInvoice(id)
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut la factura %s: %s", id, err)
	}

	customer, err := i.dbService.FindCustomer(invoice.CustomerId)
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut trobar el client %d de la factura %s: %s", invoice.CustomerId, id, err)
	}

	products, err := i.loadProducts()
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut carregar els productes des de la base de dades: %s", err)
	}

	return i.run(invoice, customer, products)
}

func (i InvoiceReport) loadProducts() (map[string]model.Product, error) {
	products, err := i.dbService.FindAllProducts()
	if err != nil {
		return nil, err
	}

	var productMap = make(map[string]model.Product)
	for _, product := range products {
		productMap[product.Id] = product
	}
	return productMap, nil
}

func (i InvoiceReport) run(invoice model.Invoice, customer model.Customer, products map[string]model.Product) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant l'informe de la factura %s ...\n", invoice.Id))

	report := Report{
		PageOrientation: consts.Portrait,
		Title:           "Factura",
		Footer:          i.footer(invoice),
		SubReports: []SubReport{
			CardSubReport{
				Title: "",
				Align: consts.Left,
				Captions: []string{
					"Número",
					"Data",
					"NIF/CIF",
					"Client",
					"Adreça",
					"Infants",
				},
				Widths: []uint{
					2,
					6,
				},
				Data: i.headerData(invoice, customer),
			},
			TableSubReport{
				Title: "Detall dels consums",
				Align: consts.Left,
				Captions: []string{
					"Unitats",
					"Element",
					"Preu",
					"Import",
					"% IVA",
					"IVA",
					"Total",
				},
				Widths: []uint{
					1,
					4,
					1,
					1,
					1,
					1,
					1,
				},
				Data: i.linesData(invoice, products),
			},
			CardSubReport{
				Title: "Total",
				Align: consts.Left,
				Captions: []string{
					"Suma imports",
					"Suma IVA",
					"Total factura",
				},
				Widths: []uint{
					2,
					6,
				},
				Data: i.summaryData(invoice),
			},
			CardSubReport{
				Title: "Notes",
				Align: consts.Left,
				Captions: []string{
					"",
					"",
				},
				Widths: []uint{
					1,
					9,
				},
				Data: i.notesData(invoice, customer),
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

func (i InvoiceReport) footer(invoice model.Invoice) string {

	var ids []string
	for _, id := range invoice.ChildrenIds {
		ids = append(ids, strconv.Itoa(id))
	}
	return fmt.Sprintf("%d: %s [%s]", invoice.CustomerId, strings.Join(ids, ","), invoice.YearMonth)
}

func (i InvoiceReport) headerData(invoice model.Invoice, customer model.Customer) [][]string {

	var data = [][]string{
		{invoice.Id},
		{invoice.DateFmt()},
		{customer.InvoiceHolder.TaxID},
		{customer.InvoiceHolder.Name},
		{customer.InvoiceHolder.Address.CompleteAddress()},
		{customer.ChildrenNames(",")},
	}
	return data
}

func (i InvoiceReport) linesData(invoice model.Invoice, products map[string]model.Product) [][]string {
	var data [][]string
	for _, line := range invoice.Lines {
		price := line.Units * line.ProductPrice
		vat := price * line.TaxPercentage
		var row = []string{
			fmt.Sprintf("%.2f", line.Units),
			products[line.ProductId].Name,
			fmt.Sprintf("%.2f", line.ProductPrice),
			fmt.Sprintf("%.2f", price),
			fmt.Sprintf("%.2f", line.TaxPercentage),
			fmt.Sprintf("%.2f", vat),
			fmt.Sprintf("%.2f", price+vat),
		}
		data = append(data, row)
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i][1] < data[j][1]
	})
	return data
}

func (i InvoiceReport) summaryData(invoice model.Invoice) [][]string {
	var price, vat, sum float64
	for _, line := range invoice.Lines {
		price += line.Units * line.ProductPrice
		vat += price * line.TaxPercentage
		sum += price + vat
	}

	var data = [][]string{
		{fmt.Sprintf("%.2f", price)},
		{fmt.Sprintf("%.2f", vat)},
		{fmt.Sprintf("%.2f", sum)},
	}
	return data
}

func (i InvoiceReport) notesData(invoice model.Invoice, customer model.Customer) [][]string {
	return [][]string{
		{i.getPaymentType(invoice, customer)},
		{i.getNotes(invoice, customer)},
	}
}

func (i InvoiceReport) getNotes(invoice model.Invoice, customer model.Customer) string {
	var buffer bytes.Buffer

	if customer.Note != "" {
		buffer.WriteString(fmt.Sprintf("- %s\n", customer.Note))
	}
	if invoice.Note != "" {
		buffer.WriteString(fmt.Sprintf("- %s\n", customer.Note))
	}
	return buffer.String()
}

func (i InvoiceReport) getPaymentType(invoice model.Invoice, customer model.Customer) string {
	paymentType := invoice.PaymentType
	if paymentType == payment_type.Invalid {
		paymentType = customer.InvoiceHolder.PaymentType
	}
	return fmt.Sprintf("- Tipus de pagament: %s", paymentType.String())
}

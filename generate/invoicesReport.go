package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"log"
	"path"
	"sam/model"
	"sam/storage"
	"sam/translate/catalan"
	"sam/util"
	"sort"
	"strings"
	"time"
)

type InvoicesReportGenerator struct {
	getManager      util.HttpGetManager
	customerStorage storage.CustomerStorage
}

func NewInvoicesReportGenerator(getManager util.HttpGetManager) InvoicesReportGenerator {
	return InvoicesReportGenerator{
		getManager,
		storage.NewCustomerStorage(),
	}
}

func (i InvoicesReportGenerator) generate() (string, error) {

	invoices, err := i.getInvoices(i.getManager)
	if err != nil {
		return "", err
	}

	contents := i.buildContents(invoices)

	filePath := path.Join(
		getWorkingDirectory(),
		viper.GetString("files.invoicesReport"),
	)
	month, err := time.Parse(util.YearMonthLayout, viper.GetString("yearMonth"))
	if err != nil {
		log.Fatal(err)
	}

	reportInfo := util.ReportInfo{
		consts.Landscape,
		consts.Left,
		fmt.Sprintf("Factures %s", catalan.MonthName(month)),
		[]util.Column{
			{"Factura", 1},
			{"Data", 1},
			{"Client", 2},
			{"Infants", 2},
			{"Concepte", 4},
			{"Import", 1},
			{"Pagament", 1},
		},
		contents,
		filePath,
	}
	err = util.PdfReport(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de clients a '%s'", filePath), nil
}

type monthInvoices struct {
	Embedded struct {
		Invoices []model.Invoice `json:"invoices"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (i InvoicesReportGenerator) getInvoices(getManager util.HttpGetManager) (*monthInvoices, error) {
	ym := viper.GetString("yearMonth")
	url := fmt.Sprintf("%s/invoices/search/findByYearMonthIn?yearMonths=%s", viper.GetString("urls.hobbit"), ym)
	invoices := new(monthInvoices)
	err := getManager.Type(url, invoices)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i InvoicesReportGenerator) buildContents(invoices *monthInvoices) [][]string {
	var contents [][]string
	for _, invoice := range invoices.Embedded.Invoices {
		customer := i.customer(invoice)

		var line = []string{
			invoice.Code(),
			invoice.Date,
			customer.FirstAdultNameWithCode(),
			customer.ChildrenNames("\n"),
			i.lines(invoice),
			i.amount(invoice),
			i.payment(invoice),
		}
		contents = append(contents, line)
	}
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents
}

func (i InvoicesReportGenerator) customer(invoice model.Invoice) model.Customer {
	customer, err := i.customerStorage.GetCustomer(invoice.CustomerID)
	if err != nil {
		log.Fatal(err)
	}
	return customer
}

func (i InvoicesReportGenerator) lines(invoice model.Invoice) string {
	var lines []string
	for _, line := range invoice.Lines {
		lines = append(lines, fmt.Sprintf(
			"%.1f %s (%.2f)",
			line.Units,
			line.ProductID,
			line.Units*line.ProductPrice,
		),
		)
	}
	return strings.Join(lines, ", ")
}

func (i InvoicesReportGenerator) amount(invoice model.Invoice) string {
	var amount float32
	for _, line := range invoice.Lines {
		amount += line.Units * line.ProductPrice
	}
	return fmt.Sprintf("%.2f", amount)
}

func (i InvoicesReportGenerator) payment(invoice model.Invoice) string {
	switch invoice.PaymentType {
	case "BANK_DIRECT_DEBIT":
		return "Rebut"
	case "BANK_TRANSFER":
		return "Tranferència"
	case "CASH":
		return "Efectiu"
	case "VOUCHER":
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

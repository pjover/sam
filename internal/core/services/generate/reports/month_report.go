package reports

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/core"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/lang"
	"log"
	"path"
	"sort"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
)

type MonthReport struct {
	configService ports.ConfigService
	langService   lang.LangService
	dbService     ports.DbService
}

func NewMonthReport(configService ports.ConfigService, langService lang.LangService, dbService ports.DbService) MonthReport {
	return MonthReport{
		configService: configService,
		langService:   langService,
		dbService:     dbService,
	}
}

func (m MonthReport) Run() (string, error) {
	yearMonth, month := m.getMonth()
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant l'informe de factures del mes %s ...\n", yearMonth))

	invoices, err := m.getInvoices(yearMonth)
	if err != nil {
		return "", err
	}
	buffer.WriteString(fmt.Sprintf("Recuperades %d factures del mes %s\n", len(invoices), yearMonth))

	contents, err := m.buildContents(invoices)
	if err != nil {
		return "", err
	}

	wd, err := m.configService.GetWorkingDirectory()
	if err != nil {
		return "", err
	}
	filePath := path.Join(wd, m.configService.Get("files.invoicesReport"))

	reportInfo := ReportInfo{
		consts.Landscape,
		consts.Left,
		fmt.Sprintf("Factures %s", m.langService.MonthName(month)),
		[]Column{
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
	err = Report(reportInfo)
	if err != nil {
		return "", fmt.Errorf("error generant l'informe: %s", err)
	}

	buffer.WriteString(fmt.Sprintf("Generat l'informe de clients a '%s'", filePath))
	return buffer.String(), nil
}

func (m MonthReport) getInvoices(yearMonth string) ([]model.Invoice, error) {
	invoices, err := m.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return nil, fmt.Errorf("error recuperant les factures del mes %s: %s", yearMonth, err)
	}
	if len(invoices) == 0 {
		return nil, fmt.Errorf("no s'han trobat factures al mes %s", yearMonth)
	}
	return invoices, nil
}

func (m MonthReport) buildContents(invoices []model.Invoice) ([][]string, error) {
	var contents [][]string
	for _, invoice := range invoices {
		customer, err := m.customer(invoice)
		if err != nil {
			return nil, fmt.Errorf("error al recuperar el client %d de la factura %s: %s", invoice.CustomerID, invoice.Code, err)
		}

		var line = []string{
			invoice.Code,
			invoice.DateFmt(),
			customer.FirstAdultNameWithCode(),
			customer.ChildrenNames("\n"),
			invoice.LinesFmt(", "),
			fmt.Sprintf("%.2f", invoice.Amount()),
			invoice.PaymentFmt(),
		}
		contents = append(contents, line)
	}
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents, nil
}

func (m MonthReport) getMonth() (string, time.Time) {
	yearMonth := m.configService.Get("yearMonth")
	month, err := time.Parse(core.YearMonthLayout, yearMonth)
	if err != nil {
		log.Fatal(fmt.Errorf("format incorrecte a la variable de configuració yearMonth '%s': %s", yearMonth, err))
	}
	return yearMonth, month
}

func (m MonthReport) customer(invoice model.Invoice) (model.Customer, error) {
	customer, err := m.dbService.FindCustomer(invoice.CustomerID)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

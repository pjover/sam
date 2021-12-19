package reports

import (
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
	fmt.Println("Generant l'informe de factures del mes ...")

	invoices, err := m.getInvoices()
	if err != nil {
		return "", err
	}

	contents, err := m.buildContents(invoices)
	if err != nil {
		return "", err
	}
	filePath := path.Join(
		m.configService.GetWorkingDirectory(),
		m.configService.Get("files.invoicesReport"),
	)
	month, err := time.Parse(core.YearMonthLayout, m.configService.Get("yearMonth"))
	if err != nil {
		log.Fatal(err)
	}

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
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de clients a '%s'", filePath), nil
}

func (m MonthReport) getInvoices() ([]model.Invoice, error) {
	ym := m.configService.Get("yearMonth")
	return m.dbService.FindInvoicesByYearMonth(ym)
}

func (m MonthReport) buildContents(invoices []model.Invoice) ([][]string, error) {
	var contents [][]string
	for _, invoice := range invoices {
		customer, err := m.customer(invoice)
		if err != nil {
			return nil, err
		}

		var line = []string{
			invoice.Code(),
			invoice.Date.String(),
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

func (m MonthReport) customer(invoice model.Invoice) (model.Customer, error) {
	customer, err := m.dbService.FindCustomer(invoice.CustomerID)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

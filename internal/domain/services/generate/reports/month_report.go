package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/lang"
	"path"
	"sort"
)

type MonthReport struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
	langService   lang.LangService
}

func NewMonthReport(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService, langService lang.LangService) MonthReport {
	return MonthReport{
		configService: configService,
		langService:   langService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (m MonthReport) Run() (string, error) {
	yearMonth := m.configService.GetCurrentYearMonth()
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant l'informe de factures del mes %s ...\n", yearMonth))

	invoices, err := m.getInvoices(yearMonth)
	if err != nil {
		return "", err
	}
	buffer.WriteString(fmt.Sprintf("Recuperades %d factures del mes %s\n", len(invoices), yearMonth))

	data, err := m.buildData(invoices)
	if err != nil {
		return "", err
	}

	reportDefinition := ReportDefinition{
		PageOrientation: consts.Landscape,
		Title:           fmt.Sprintf("Factures %s", m.langService.MonthName(yearMonth.Month())),
		Footer:          m.osService.Now().Format(domain.YearMonthDayLayout),
		SubReports: []SubReport{
			TableSubReport{
				Align: consts.Left,
				Captions: []string{
					"Factura",
					"Data",
					"Client",
					"Infants",
					"Concepte",
					"Import",
					"Pagament",
				},
				Widths: []uint{
					1,
					1,
					2,
					2,
					4,
					1,
					1,
				},
				Data: data,
			},
		},
	}

	wd := m.configService.GetWorkingDirectory()
	filePath := path.Join(
		wd,
		m.configService.GetString("files.invoicesReport"),
	)

	reportService := NewReportService(m.configService)
	err = reportService.SaveToFile(reportDefinition, filePath)
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("Generat l'informe de clients a '%s'", filePath))
	return buffer.String(), nil
}

func (m MonthReport) getInvoices(yearMonth model.YearMonth) ([]model.Invoice, error) {
	invoices, err := m.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return nil, fmt.Errorf("error recuperant les factures del mes %s: %s", yearMonth, err)
	}
	if len(invoices) == 0 {
		return nil, fmt.Errorf("no s'han trobat factures al mes %s", yearMonth)
	}
	return invoices, nil
}

func (m MonthReport) buildData(invoices []model.Invoice) ([][]string, error) {
	var data [][]string
	for _, invoice := range invoices {
		customer, err := m.customer(invoice)
		if err != nil {
			return nil, fmt.Errorf("error al recuperar el client %d de la factura %s: %s", invoice.CustomerId, invoice.Id, err)
		}

		var line = []string{
			invoice.Id,
			invoice.DateFmt(),
			customer.FirstAdultNameWithId(),
			customer.ChildrenNamesWithId("\n"),
			invoice.LinesFmt(", "),
			fmt.Sprintf("%.2f", invoice.Amount()),
			invoice.PaymentType.String(),
		}
		data = append(data, line)
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	return data, nil
}

func (m MonthReport) customer(invoice model.Invoice) (model.Customer, error) {
	customer, err := m.dbService.FindCustomer(invoice.CustomerId)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

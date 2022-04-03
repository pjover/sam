package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/pjover/sam/internal/domain/services/loader"
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

	bulkLoader := loader.NewBulkLoader(m.configService, m.dbService)
	invoices, err := bulkLoader.LoadMonthInvoicesByPaymentType()
	if err != nil {
		return "", err
	}

	subReports, err := m.paymentTypeTables(invoices)
	if err != nil {
		return "", err
	}
	cardSubReport := m.summaryCard(invoices)
	subReports = append(subReports, cardSubReport)

	reportDefinition := ReportDefinition{
		PageOrientation: consts.Landscape,
		Title:           fmt.Sprintf("Factures %s", m.langService.MonthName(yearMonth.Month())),
		Footer:          m.osService.Now().Format(domain.YearMonthDayLayout),
		SubReports:      subReports,
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

func (m MonthReport) paymentTypeTables(invoices map[payment_type.PaymentType][]model.Invoice) ([]SubReport, error) {
	var subReports []SubReport
	for paymentType, paymentInvoices := range invoices {
		data, err := m.buildData(paymentInvoices)
		if err != nil {
			return nil, err
		}
		subReport := TableSubReport{
			Title: paymentType.Format(),
			Align: consts.Left,
			Captions: []string{
				"Factura",
				"Data",
				"Client",
				"Infants",
				"Concepte",
				"Import",
			},
			Widths: []uint{
				1,
				1,
				2,
				2,
				5,
				1,
			},
			Data: data,
		}
		subReports = append(subReports, subReport)
	}
	return subReports, nil
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
		}
		data = append(data, line)
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	return data, nil
}

func (m MonthReport) summaryCard(invoices map[payment_type.PaymentType][]model.Invoice) SubReport {

	var amountsByPaymentType = make(map[payment_type.PaymentType]float64)
	for paymentType, paymentInvoices := range invoices {
		var amountSum float64
		for _, invoice := range paymentInvoices {
			amountSum += invoice.Amount()
		}
		amountsByPaymentType[paymentType] = amountSum
	}

	var captions []string
	var data [][]string
	var total float64
	for paymentType, sum := range amountsByPaymentType {
		total += sum
		captions = append(captions, paymentType.Format())
		datum := []string{fmt.Sprintf("%.2f", sum)}
		data = append(data, datum)
	}

	captions = append(captions, "TOTAL")
	datum := []string{fmt.Sprintf("%.2f", total)}
	data = append(data, datum)
	cardSubReport := CardSubReport{
		Title:    "Resum",
		Align:    consts.Left,
		Captions: captions,
		Widths: []uint{
			2,
			6,
		},
		Data: data,
	}
	return cardSubReport
}

func (m MonthReport) customer(invoice model.Invoice) (model.Customer, error) {
	customer, err := m.dbService.FindCustomer(invoice.CustomerId)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

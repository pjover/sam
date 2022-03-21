package reports

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/lang"
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

	data, err := m.buildData(invoices)
	if err != nil {
		return "", err
	}

	report := Report{
		PageOrientation: consts.Landscape,
		Title:           fmt.Sprintf("Factures %s", m.langService.MonthName(month)),
		Footer:          time.Now().Format("2006-01-02"),
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

	wd, err := m.configService.GetWorkingDirectory()
	if err != nil {
		return "", err
	}
	filePath := path.Join(
		wd,
		m.configService.GetString("files.invoicesReport"),
	)
	err = report.SaveToFile(filePath)
	if err != nil {
		return "", err
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

func (m MonthReport) getMonth() (string, time.Time) {
	yearMonth := m.configService.GetString("yearMonth")
	month, err := time.Parse(domain.YearMonthLayout, yearMonth)
	if err != nil {
		log.Fatal(fmt.Errorf("format incorrecte a la variable de configuració yearMonth '%s': %s", yearMonth, err))
	}
	return yearMonth, month
}

func (m MonthReport) customer(invoice model.Invoice) (model.Customer, error) {
	customer, err := m.dbService.FindCustomer(invoice.CustomerId)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

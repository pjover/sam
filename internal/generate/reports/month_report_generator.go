package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core"
	model2 "github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/lang"
	"log"
	"path"
	"sort"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/generate"
	"github.com/spf13/viper"
)

type MonthReportGenerator struct {
	getManager    tuk.HttpGetManager
	dbService     ports.DbService
	configService ports.ConfigService
	langService   lang.LangService
}

func NewMonthReportGenerator() generate.Generator {
	cfgService := cfg.NewConfigService()
	return MonthReportGenerator{
		tuk.NewHttpGetManager(),
		mongo_db.NewDbService(),
		cfgService,
		lang.NewLangService(cfgService.Get("lang")),
	}
}

func (i MonthReportGenerator) Generate() (string, error) {
	fmt.Println("Generant l'informe de factures del mes ...")

	invoices, err := i.getInvoices(i.getManager)
	if err != nil {
		return "", err
	}

	contents, err := i.buildContents(invoices)
	if err != nil {
		return "", err
	}
	filePath := path.Join(
		i.configService.GetWorkingDirectory(),
		viper.GetString("files.invoicesReport"),
	)
	month, err := time.Parse(core.YearMonthLayout, viper.GetString("yearMonth"))
	if err != nil {
		log.Fatal(err)
	}

	reportInfo := ReportInfo{
		consts.Landscape,
		consts.Left,
		fmt.Sprintf("Factures %s", i.langService.MonthName(month)),
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
	err = PdfReport(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de clients a '%s'", filePath), nil
}

type monthInvoices struct {
	Embedded struct {
		Invoices []model2.Invoice `json:"invoices"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (i MonthReportGenerator) getInvoices(getManager tuk.HttpGetManager) (*monthInvoices, error) {
	ym := viper.GetString("yearMonth")
	url := fmt.Sprintf("%s/invoices/search/findByYearMonthIn?yearMonths=%s", viper.GetString("urls.hobbit"), ym)
	invoices := new(monthInvoices)
	err := getManager.Type(url, invoices)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i MonthReportGenerator) buildContents(invoices *monthInvoices) ([][]string, error) {
	var contents [][]string
	for _, invoice := range invoices.Embedded.Invoices {
		customer, err := i.customer(invoice)
		if err != nil {
			return nil, err
		}

		var line = []string{
			invoice.Code(),
			invoice.Date,
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

func (i MonthReportGenerator) customer(invoice model2.Invoice) (model2.Customer, error) {
	customer, err := i.dbService.GetCustomer(invoice.CustomerID)
	if err != nil {
		return model2.Customer{}, err
	}
	return customer, nil
}

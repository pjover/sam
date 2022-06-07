package generate

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/generate/bdd"
	"github.com/pjover/sam/internal/domain/services/generate/reports"
	"github.com/pjover/sam/internal/domain/services/lang"
)

type generateService struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
	langService   lang.LangService
}

func NewGenerateService(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService, langService lang.LangService) ports.GenerateService {
	return generateService{
		configService: configService,
		langService:   langService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (g generateService) CustomerReport() (string, error) {
	generator := reports.NewCustomerReport(g.configService, g.dbService, g.osService)
	return generator.Run()
}

func (g generateService) MonthReport(yearMonth model.YearMonth) (string, error) {
	generator := reports.NewMonthReport(g.configService, g.dbService, g.osService, g.langService)
	return generator.Run(yearMonth)
}

func (g generateService) ProductReport() (string, error) {
	generator := reports.NewProductsReport(g.configService, g.dbService, g.osService)
	return generator.Run()
}

func (g generateService) SingleInvoice(invoiceId string) (string, error) {
	generator := reports.NewInvoiceReport(g.configService, g.osService, g.dbService)
	return generator.SingleInvoice(invoiceId)
}

func (g generateService) MonthInvoices(yearMonth model.YearMonth) (string, error) {
	generator := reports.NewInvoiceReport(g.configService, g.osService, g.dbService)
	return generator.MonthInvoices(yearMonth)
}

func (g generateService) BddFile() (string, error) {
	generator := bdd.NewBddService(g.configService, g.dbService, g.osService)
	return generator.Run()
}

func (g generateService) CustomersCards() (string, error) {
	generator := reports.NewCustomerCardsReports(g.configService, g.dbService, g.osService)
	return generator.Run()
}

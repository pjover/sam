package generate

import (
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/generate/reports"
	"github.com/pjover/sam/internal/domain/services/lang"
)

type generateService struct {
	configService ports.ConfigService
	langService   lang.LangService
	dbService     ports.DbService
}

func NewGenerateService(configService ports.ConfigService, langService lang.LangService, dbService ports.DbService) ports.GenerateService {
	return generateService{
		configService: configService,
		langService:   langService,
		dbService:     dbService,
	}
}

func (g generateService) CustomerReport() (string, error) {
	generator := reports.NewCustomerReport(g.configService, g.dbService)
	return generator.Run()
}

func (g generateService) MonthReport() (string, error) {
	generator := reports.NewMonthReport(g.configService, g.langService, g.dbService)
	return generator.Run()
}

func (g generateService) ProductReport() (string, error) {
	generator := reports.NewProductsReport(g.configService, g.dbService)
	return generator.Run()
}

func (g generateService) SingleInvoice(invoiceId string) (string, error) {
	generator := reports.NewInvoiceReport(g.configService, g.dbService)
	return generator.SingleInvoice(invoiceId)
}

func (g generateService) MonthInvoices() (string, error) {
	generator := reports.NewInvoiceReport(g.configService, g.dbService)
	return generator.MonthInvoices()
}

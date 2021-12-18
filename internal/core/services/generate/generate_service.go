package generate

import (
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/generate/reports"
)

type generateService struct {
	dbService ports.DbService
}

func NewGenerateService(dbService ports.DbService) ports.GenerateService {
	return generateService{
		dbService: dbService,
	}
}

func (g generateService) GenerateProduct() (string, error) {
	generator := reports.NewProductsReportGenerator(g.dbService)
	return generator.Generate()
}

package e2e

import (
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/pjover/sam/internal/domain/services/generate"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/pjover/sam/internal/domain/services/list"
	"github.com/pjover/sam/internal/domain/services/loader"
	"github.com/pjover/sam/test/fakes"
)

func InjectDependencies() ports.CommandManager {

	configService := fakes.FakeConfigService()
	dbService := fakes.FakeDbService()
	osService := fakes.FakeOsService()
	langService := lang.NewCatLangService()

	billingService := billing.NewBillingService(configService, dbService, osService)
	bulkLoader := loader.NewBulkLoader(configService, dbService)
	listService := list.NewListService(configService, dbService, bulkLoader)
	generateService := generate.NewGenerateService(configService, dbService, osService, langService)

	return NewE2eCommandManager(configService, billingService, listService, generateService)
}

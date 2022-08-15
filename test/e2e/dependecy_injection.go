package e2e

import (
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/pjover/sam/internal/domain/services/list"
	"github.com/pjover/sam/internal/domain/services/loader"
	"github.com/pjover/sam/test/fakes"
)

func InjectDependencies() ports.CommandManager {

	configService := fakes.FakeConfigService()
	dbService := fakes.FakeDbService()
	osService := fakes.FakeOsService()

	billingService := billing.NewBillingService(configService, dbService, osService)
	bulkLoader := loader.NewBulkLoader(configService, dbService)
	listService := list.NewListService(configService, dbService, bulkLoader)

	return NewE2eCommandManager(billingService, listService)
}

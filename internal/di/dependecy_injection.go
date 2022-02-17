package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/os"
	consumCmd "github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/pjover/sam/internal/generate/bbd"
	"github.com/pjover/sam/internal/generate/invoices"
)

func MainDI(configService ports.ConfigService, cmdManager cli.CmdManager) {

	osService := os.NewOsService()
	langService := lang.NewCatLangService()
	dbService := mongo_db.NewDbService(configService)
	postManager := hobbit.NewHttpPostManager()

	adminServiceDI(configService, cmdManager, osService, langService)
	editServiceDI(configService, cmdManager, osService)
	displayServiceDI(dbService, cmdManager)
	generateServiceDI(configService, langService, dbService, cmdManager)
	listServiceDI(dbService, cmdManager, osService)
	searchServiceDI(dbService, cmdManager)
	billingServiceDI(dbService, osService, cmdManager, postManager)

	// TODO move to DI and remove method AddTmpCommand
	cmdManager.AddTmpCommand(consumCmd.NewRectifyConsumptionsCmd(postManager, dbService))

	cmdManager.AddTmpCommand(generate.NewGenerateBddCmd(bbd.NewBddGenerator()))
	cmdManager.AddTmpCommand(generate.NewGenerateMonthInvoicesCmd(invoices.NewMonthInvoicesGenerator()))
	cmdManager.AddTmpCommand(generate.NewGenerateSingleInvoiceCmd(invoices.NewSingleInvoiceGenerator()))
}

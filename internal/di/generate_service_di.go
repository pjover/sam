package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	generateCli "github.com/pjover/sam/internal/adapters/cli/generate"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/generate"
	"github.com/pjover/sam/internal/core/services/lang"
)

func generateServiceDI(configService ports.ConfigService, langService lang.LangService, dbService ports.DbService, cmdManager cli.CmdManager) {
	generateService := generate.NewGenerateService(configService, langService, dbService)
	cmdManager.AddCommand(generateCli.NewGenerateCustomerReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateMonthReportCmd(generateService))
	cmdManager.AddCommand(generateCli.NewGenerateProductReportCmd(generateService))
}

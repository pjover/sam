package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	generateCli "github.com/pjover/sam/internal/adapters/cli/generate"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/generate"
)

func generateServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	generateService := generate.NewGenerateService(dbService)
	cmdManager.AddCommand(generateCli.NewGenerateProductCmd(generateService))
}

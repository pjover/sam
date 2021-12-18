package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	listCli "github.com/pjover/sam/internal/adapters/cli/list"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/list"
)

func listServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	listService := list.NewListService(dbService)
	cmdManager.AddCommand(listCli.NewListProductCmd(listService))
}

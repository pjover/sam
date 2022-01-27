package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	listCli "github.com/pjover/sam/internal/adapters/cli/list"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/list"
)

func listServiceDI(dbService ports.DbService, cmdManager cli.CmdManager, osService ports.OsService) {
	listService := list.NewListService(dbService)
	cmdManager.AddCommand(listCli.NewListProductsCmd(listService))
	cmdManager.AddCommand(listCli.NewListInvoicesCmd(listService, osService))
	cmdManager.AddCommand(listCli.NewListCustomersCmd(listService))
	cmdManager.AddCommand(listCli.NewListChildrenCmd(listService))
	cmdManager.AddCommand(listCli.NewListMailsCmd(listService))
	cmdManager.AddCommand(listCli.NewListConsumptionsCmd(listService))
}

package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	searchCmd "github.com/pjover/sam/internal/adapters/cli/search"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/search"
)

func searchServiceDI(dbService ports.DbService, cmdManager cli.CmdManager) {
	searchService := search.NewSearchService(dbService)
	cmdManager.AddCommand(searchCmd.NewSearchCustomerCmd(searchService))
}

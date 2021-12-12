package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/edit"
	"github.com/pjover/sam/internal/adapters/mongo_express"
	"github.com/pjover/sam/internal/core/ports"
)

func editServiceDI(cmdManager cli.CmdManager, osService ports.OsService) {
	editService := mongo_express.NewEditService(osService)
	cmdManager.AddCommand(edit.NewEditCustomerCmd(editService))
	cmdManager.AddCommand(edit.NewEditInvoiceCmd(editService))
	cmdManager.AddCommand(edit.NewEditProductCmd(editService))
}

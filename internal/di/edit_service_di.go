package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/edit"
	"github.com/pjover/sam/internal/adapters/mongo_express"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services"
)

func editServiceDI(cfgService ports.ConfigService, cmdManager cli.CmdManager, osService ports.OsService) {
	externalEditor := mongo_express.NewExternalEditor(cfgService, osService)
	editService := services.NewEditService(externalEditor)
	cmdManager.AddCommand(edit.NewEditCustomerCmd(editService))
	cmdManager.AddCommand(edit.NewEditInvoiceCmd(editService))
	cmdManager.AddCommand(edit.NewEditProductCmd(editService))
}

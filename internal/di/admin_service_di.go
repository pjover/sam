package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/admin"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services"
	"github.com/pjover/sam/internal/core/services/lang"
)

func adminServiceDI(cfgService ports.ConfigService, cmdManager cli.CmdManager, osService ports.OsService, langService lang.LangService) {
	adminService := services.NewAdminService(cfgService, osService, langService)
	cmdManager.AddCommand(admin.NewBackupCmd(adminService))
	cmdManager.AddCommand(admin.NewDirectoryCmd(adminService))
}

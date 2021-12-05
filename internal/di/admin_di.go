package di

import (
	"github.com/pjover/sam/internal/adapters/primary/cli/admin"
	"github.com/pjover/sam/internal/core/services"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
)

func adminServiceDI(rootCmd *cobra.Command) {
	timer := shared.SamTimeManager{}
	adminService := services.NewAdminService(timer)

	backupCmd := admin.NewBackupCmd(adminService)
	rootCmd.AddCommand(backupCmd.Cmd())
}

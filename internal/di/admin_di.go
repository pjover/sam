package di

import (
	"github.com/pjover/sam/internal/adapters/primary/cli/admin"
	"github.com/pjover/sam/internal/core/env"
	"github.com/pjover/sam/internal/core/os"
	"github.com/pjover/sam/internal/core/services"
	"github.com/spf13/cobra"
)

func adminServiceDI(rootCmd *cobra.Command) {
	adminService := services.NewAdminService(os.NewTimeManager(), os.NewFileManager(), env.NewConfigManager(), os.NewOpenManager())

	backupCmd := admin.NewBackupCmd(adminService)
	rootCmd.AddCommand(backupCmd.Cmd())

	directoryCmd := admin.NewDirectoryCmd(adminService)
	rootCmd.AddCommand(directoryCmd.Cmd())
}

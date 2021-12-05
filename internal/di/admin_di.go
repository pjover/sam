package di

import (
	"github.com/pjover/sam/internal/adapters/cfg"
	admin2 "github.com/pjover/sam/internal/adapters/cli/admin"
	"github.com/pjover/sam/internal/core/os"
	"github.com/pjover/sam/internal/core/services"
	"github.com/spf13/cobra"
)

func adminServiceDI(rootCmd *cobra.Command) {
	adminService := services.NewAdminService(cfg.NewConfigService(), os.NewTimeManager(), os.NewFileManager(), os.NewOpenManager())

	backupCmd := admin2.NewBackupCmd(adminService)
	rootCmd.AddCommand(backupCmd.Cmd())

	directoryCmd := admin2.NewDirectoryCmd(adminService)
	rootCmd.AddCommand(directoryCmd.Cmd())
}

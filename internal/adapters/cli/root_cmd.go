package cli

import (
	"github.com/pjover/sam/internal/core"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/cobra"
)

type CmdManager interface {
	AddCommand(cmd Cmd)
	AddTmpCommand(cmd *cobra.Command) //TODO remove
	Execute()
}

type cmdManager struct {
	cfgService ports.ConfigService
	rootCmd    *cobra.Command
}

func NewCmdManager(cfgService ports.ConfigService) CmdManager {
	// RootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "sam",
		Short: "A Command Line Interface to Hobbit service",
		Long: `A Command Line Interface to Hobbit service in Go.
	Complete documentation is available at https://github.com/pjover/sam`,
		Version: core.Version,
	}
	cobra.OnInitialize(cfgService.Init)
	return cmdManager{cfgService, rootCmd}
}

func (c cmdManager) GetRootCmd() *cobra.Command {
	return c.rootCmd
}

func (c cmdManager) AddCommand(cmd Cmd) {
	c.rootCmd.AddCommand(cmd.Cmd())
}

func (c cmdManager) AddTmpCommand(cmd *cobra.Command) { //TODO Remove
	c.rootCmd.AddCommand(cmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It only needs to happen once to the RootCmd.
func (c cmdManager) Execute() {
	cobra.CheckErr(c.rootCmd.Execute())
}

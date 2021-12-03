package cmd

import (
	"fmt"
	"github.com/pjover/sam/internal/cmd/adm"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/edit"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
	"os"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sam",
	Short: "A Command Line Interface to Hobbit service",
	Long: `A Command Line Interface to Hobbit service in Go.
	Complete documentation is available at https://github.com/pjover/sam`,
	Version: shared.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(InitConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sam)")

	RootCmd.AddCommand(adm.NewDirectoryCmd())

	RootCmd.AddCommand(consum.NewBillConsumptionsCmd())
	RootCmd.AddCommand(consum.NewInsertConsumptionsCmd())
	RootCmd.AddCommand(consum.NewRectifyConsumptionsCmd())

	RootCmd.AddCommand(display.NewDisplayCustomerCmd())
	RootCmd.AddCommand(display.NewDisplayInvoiceCmd())
	RootCmd.AddCommand(display.NewDisplayProductCmd())

	RootCmd.AddCommand(edit.NewEditCustomerCmd())
	RootCmd.AddCommand(edit.NewEditInvoiceCmd())
	RootCmd.AddCommand(edit.NewEditProductCmd())

	RootCmd.AddCommand(generate.NewGenerateBddCmd())
	RootCmd.AddCommand(generate.NewGenerateCustomersReportCmd())
	RootCmd.AddCommand(generate.NewGenerateMonthInvoicesCmd())
	RootCmd.AddCommand(generate.NewGenerateMonthReportCmd())
	RootCmd.AddCommand(generate.NewGenerateProductsReportCmd())
	RootCmd.AddCommand(generate.NewGenerateSingleInvoiceCmd())

	RootCmd.AddCommand(list.NewListChildrenCmd())
	RootCmd.AddCommand(list.NewListConsumptionsCmd())
	RootCmd.AddCommand(list.NewListCustomersCmd())
	RootCmd.AddCommand(list.NewListInvoicesCmd())
	RootCmd.AddCommand(list.NewListMailsCmd())
	RootCmd.AddCommand(list.NewListProductsCmd())

	RootCmd.AddCommand(search.NewSearchCustomerCmd())
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".sam" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("sam")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.sam")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	shared.DefaultConfig(home)
}

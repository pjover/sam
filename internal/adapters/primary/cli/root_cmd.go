package cli

import (
	"fmt"
	"os"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(InitConfig)
}

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sam",
	Short: "A Command Line Interface to Hobbit service",
	Long: `A Command Line Interface to Hobbit service in Go.
	Complete documentation is available at https://github.com/pjover/sam`,
	Version: shared.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It only needs to happen once to the RootCmd.
func Execute() *cobra.Command {
	cobra.CheckErr(rootCmd.Execute())
	return rootCmd
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".sam" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName("sam")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.sam")

	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	shared.DefaultConfig(home)
}

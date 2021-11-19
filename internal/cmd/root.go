package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"sam/internal/util"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sam",
	Short: "A Command Line Interface to Hobbit service",
	Long: `A Command Line Interface to Hobbit service in Go.
	Complete documentation is available at https://github.com/pjover/sam`,
	Version: util.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sam)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

	util.DefaultConfig(home)
}

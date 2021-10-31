package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"

	"sam/adm"
	"sam/comm"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sam",
	Short: "A Command Line Interface to Hobbit service",
	Long: `A Command Line Interface to Hobbit service in Go.
	Complete documentation is available at https://github.com/pjover/sam`,
	Version: comm.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sam)")
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
		viper.SetConfigName(".sam")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	dirHome := path.Join(home, "Sam")
	viper.SetDefault("dirs.home", dirHome)
	yearMonth, dirName := adm.GetDirConfig(false, false)
	viper.SetDefault("dirs.current", dirName)
	viper.SetDefault("yearMonth", yearMonth)
}

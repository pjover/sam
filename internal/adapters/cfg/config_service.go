package cfg

import (
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

type configService struct {
}

func NewConfigService() ports.ConfigService {
	return configService{}
}

func (c configService) Get(key string) string {
	return viper.GetString(key)
}

func (c configService) Set(key string, value string) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}

func (c configService) Init() {
	home := c.findHomeDirectory()
	c.searchConfigInHomeDirectory(home)
	c.loadEnvironmentVariables()
	c.readConfigFile()
	c.loadDefaultConfig(home)
}

func (c configService) findHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("cannot find home directory: %s", err)
	}
	return home
}

func (c configService) searchConfigInHomeDirectory(home string) {
	viper.AddConfigPath(home)
	viper.SetConfigName("sam")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.sam")
}

func (c configService) loadEnvironmentVariables() {
	viper.AutomaticEnv()
}

func (c configService) readConfigFile() {
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}
}

func (c configService) loadDefaultConfig(home string) {
	viper.SetDefault("dirs.config", "$HOME/.sam")

	appDirectory := path.Join(home, "Sam")
	viper.SetDefault("dirs.home", appDirectory)
	viper.SetDefault("dirs.reports", "$HOME")
	viper.SetDefault("dirs.invoicesName", "invoices")
	viper.SetDefault("dirs.backup", "$HOME/.sam")

	viper.SetDefault("urls.hobbit", "http://localhost:8080")
	viper.SetDefault("urls.mongoExpress", "http://localhost:8081/db/hobbit_prod")

	viper.SetDefault("files.customersReport", "Customers.pdf")
	viper.SetDefault("files.productsReport", "Products.pdf")
	viper.SetDefault("files.invoicesReport", "Factures.pdf")
	viper.SetDefault("files.logo", "logo.png")

	viper.SetDefault("business.name", "BusinessName")
	viper.SetDefault("business.addressLine1", "AddressLine1")
	viper.SetDefault("business.addressLine2", "AddressLine2")
	viper.SetDefault("business.addressLine3", "AddressLine3")
	viper.SetDefault("business.addressLine4", "AddressLine4")
	viper.SetDefault("business.taxIdLine", "TaxIdLine")

}

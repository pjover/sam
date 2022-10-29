package cfg

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

type configService struct {
}

func NewConfigService() ports.ConfigService {
	service := configService{}
	service.Init()
	return service
}

func (c configService) GetString(key string) string {
	return viper.GetString(key)
}

func (c configService) SetString(key string, value string) error {
	return c.set(key, value)
}

func (c configService) set(key string, value interface{}) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}

func (c configService) GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func (c configService) SetTime(key string, value time.Time) error {
	return c.set(key, value)
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
	log.Printf("Using config file %s", viper.ConfigFileUsed())

	appDirectory := path.Join(home, "Sam")
	viper.SetDefault("dirs.home", appDirectory)
	viper.SetDefault("dirs.reports", "$HOME")
	viper.SetDefault("dirs.backup", "$HOME/.sam")

	viper.SetDefault("urls.hobbit", "http://localhost:8080")
	viper.SetDefault("urls.mongoExpress", "http://localhost:8081/db/hobbit_prod")

	viper.SetDefault("files.customersReport", "Llistat de clients.pdf")
	viper.SetDefault("files.productsReport", "Llistat de productes.pdf")
	viper.SetDefault("files.invoicesReport", "Llistat de factures %s.pdf")
	viper.SetDefault("files.logo", "logo.png")

	viper.SetDefault("business.name", "BusinessName")
	viper.SetDefault("business.addressLine1", "AddressLine1")
	viper.SetDefault("business.addressLine2", "AddressLine2")
	viper.SetDefault("business.addressLine3", "AddressLine3")
	viper.SetDefault("business.addressLine4", "AddressLine4")
	viper.SetDefault("business.taxIdLine", "TaxIdLine")

	viper.SetDefault("db.server", "mongodb://localhost:27017")
	viper.SetDefault("db.name", "hobbit")

	viper.SetDefault("reports.invoicesFolderName", "factures")
	viper.SetDefault("reports.customersCardsFolderName", "clients")
	viper.SetDefault("reports.lastCustomersCardsUpdated", time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local))

	viper.SetDefault("entities.newProductFileName", "new_product.json")
	viper.SetDefault("entities.newCustomerFileName", "new_customer.json")
}

func (c configService) GetCurrentYearMonth() model.YearMonth {
	yearMonth, err := model.StringToYearMonth(c.GetString("yearMonth"))
	if err != nil {
		log.Fatalf("no s'ha trobat el actual mes al valor 'yearMonth' de la configuració")
	}
	return yearMonth
}

func (c configService) SetCurrentYearMonth(yearMonth model.YearMonth) error {
	return c.SetString("yearMonth", yearMonth.String())
}

func (c configService) GetConfigDirectory() string {
	return c.GetString("dirs.config")
}

func (c configService) GetHomeDirectory() string {
	return c.GetString("dirs.home")
}
func (c configService) GetWorkingDirectory() string {
	return path.Join(c.GetString("dirs.home"), c.GetString("dirs.current"))
}

func (c configService) GetInvoicesDirectory() string {
	return path.Join(c.GetWorkingDirectory(), c.GetString("reports.invoicesFolderName"))
}

func (c configService) GetReportsDirectory() string {
	return c.GetString("dirs.reports")
}

func (c configService) GetCustomersCardsDirectory() string {
	return path.Join(c.GetReportsDirectory(), c.GetString("reports.customersCardsFolderName"))
}

func (c configService) GetBackupDirectory() string {
	return c.GetString("dirs.backup")
}

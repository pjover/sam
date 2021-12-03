package shared

import (
	"github.com/spf13/viper"
	"path"
)

func DefaultConfig(home string) {
	viper.SetDefault("dirs.config", "$HOME/.sam")

	dirHome := path.Join(home, "Sam")
	viper.SetDefault("dirs.home", dirHome)
	viper.SetDefault("dirs.reports", "$HOME")
	viper.SetDefault("dirs.invoicesName", "invoices")

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

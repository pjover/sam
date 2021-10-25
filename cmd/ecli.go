package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
	"strconv"
)

var ecliCmd = &cobra.Command{
	Use:         "ecli codiClient",
	Short:       "Edita les dades d'un client",
	Long:        "Obri un navegador per a editar les dades d'un client indicant el seu codi",
	Example:     `   ecli 152     Edita les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	Aliases:     []string{"edita-client"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateCustomerCode(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		customerCode, _ := strconv.Atoi(args[0])
		return adm.EditCustomer(customerCode)
	},
}

func init() {
	rootCmd.AddCommand(ecliCmd)
}

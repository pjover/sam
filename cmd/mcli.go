package cmd

import (
	"sam/adm"
	"sam/comm"
	"strconv"

	"github.com/spf13/cobra"
)

var mcliCmd = &cobra.Command{
	Use:         "mcli codiClient",
	Short:       "Mostra les dades d'un client",
	Long:        "Mostra les dades d'un client indicant el seu codi",
	Example:     `   mcli 152     Mostra les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	Args: func(cmd *cobra.Command, args []string) error {
		return comm.ValidateCode(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		customerCode, _ := strconv.Atoi(args[0])
		return adm.DisplayCustomer(customerCode)
	},
}

func init() {
	rootCmd.AddCommand(mcliCmd)
}

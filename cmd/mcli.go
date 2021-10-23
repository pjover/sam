package cmd

import (
	"sam/adm"

	"github.com/spf13/cobra"
)

var CustomerCode int
var ChildCode int

var mcliCmd = &cobra.Command{
	Use:   "mcli",
	Short: "Mostra les dades d'un client",
	Long:  "Mostra les dades d'un client, se pot indicar el client o l'infant",
	Example: `   mcli -i 1520    Mostra les dades del client de l'infant 1520
   mcli -c 152     Mostra les dades del client 152`,
	Annotations: map[string]string{"ADM": "Comandes d'adminitració"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMcli()
	},
}

func init() {
	rootCmd.AddCommand(mcliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mcliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mcliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runMcli() error {
	if ChildCode > 0 {
		CustomerCode = ChildCode / 10
	}
	return adm.DisplayCustomer(CustomerCode)
}

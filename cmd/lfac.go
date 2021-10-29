package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var lfacCmd = &cobra.Command{
	Use:   "lfac [AAAA-MM]",
	Short: "Llista les factura del mes indicat",
	Long:  "Llista les factura del mes indicat, si no s'especifica el mes agafa l'actual",
	Example: `   lfac     Llista les factura del mes
   lfac 2021-10    Llista les factura del mes d'Octubre de 2021`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-factures"},
	Args: func(cmd *cobra.Command, args []string) error {
		return validateNumberOfArgsGreaterThan(1, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return adm.ListYearMonthInvoices("")
		} else {
			return adm.ListYearMonthInvoices(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(lfacCmd)
}

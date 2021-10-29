package cmd

import (
	"github.com/spf13/cobra"
	"sam/adm"
)

var lproCmd = &cobra.Command{
	Use:         "lpro",
	Short:       "Llista tots els productes",
	Long:        "Llista tots els productes que hi han a la base de dades",
	Example:     `   lpro    "Llista tots els productes`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-producted"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return adm.ListProducts()
	},
}

func init() {
	rootCmd.AddCommand(lproCmd)
}

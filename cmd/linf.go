package cmd

import (
	"sam/adm"

	"github.com/spf13/cobra"
)

var linfCmd = &cobra.Command{
	Use:         "linf",
	Short:       "Llista tots els infants",
	Long:        "Llista tots els infants per grups",
	Example:     `   linf      Llista tots els infants`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-infants"},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		_, err := manager.ListChildren()
		return err
	},
}

func init() {
	rootCmd.AddCommand(linfCmd)
}

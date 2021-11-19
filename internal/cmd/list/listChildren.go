package list

import (
	"fmt"
	"sam/adm"
	"sam/internal/cmd"

	"github.com/spf13/cobra"
)

var listChildrenCmd = &cobra.Command{
	Use:         "llistaInfants",
	Short:       "Llista tots els infants",
	Long:        "Llista tots els infants per grups",
	Example:     `   llistaInfants      Llista tots els infants`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases: []string{
		"linf",
		"llistainfants", "llista-infants",
		"llistarInfants", "llistarinfants", "llistar-infants",
		"lchi",
		"listChildren", "listchildren", "list-children",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewListManager()
		msg, err := manager.ListChildren()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(listChildrenCmd)
}

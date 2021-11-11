package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/adm"
)

var generateBddCmd = &cobra.Command{
	Use:         "generaRebuts",
	Short:       "Genera el fitxer de rebuts",
	Long:        "Genera el fitxer de rebuts de les factures pendents",
	Example:     `   generaRebuts    Genera el fitxer de rebuts`,
	Annotations: map[string]string{"GEN": "Comandes de generació"},
	Aliases: []string{
		"greb",
		"generarebuts", "genera-rebuts",
		"generarRebuts", "generarrebuts", "generar-rebuts",
		"gbdd",
		"generateBdd", "generatebdd", "generate-bdd",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := adm.NewGenerateManager(nil)
		msg, err := manager.GenerateBdd()
		fmt.Println(msg)
		return err
	},
}

func init() {
	rootCmd.AddCommand(generateBddCmd)
}

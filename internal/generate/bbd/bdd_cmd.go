package bbd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/cmd"
	"sam/internal/generate"
)

func newBddCmd(generator generate.Generator) *cobra.Command {
	return &cobra.Command{
		Use:         "generaRebuts",
		Short:       "Genera el fitxer de rebuts",
		Long:        "Genera el fitxer de rebuts de les factures pendents",
		Example:     `   generaRebuts    Genera el fitxer de rebuts`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"greb",
			"generarebuts",
			"genera-rebuts",
			"generarRebuts",
			"generarrebuts",
			"generar-rebuts",
			"gbdd",
			"generateBdd",
			"generatebdd",
			"generate-bdd",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generator.Generate()
			fmt.Println(msg)
			return err
		},
	}
}

func init() {
	generator := NewBddGenerator()
	command := newBddCmd(generator)
	cmd.RootCmd.AddCommand(command)
}

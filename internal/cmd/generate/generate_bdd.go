package generate

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/internal/cmd"
	"sam/internal/generate"
	"sam/internal/generate/bbd"
)

func newGenerateBddCmd(generator generate.Generator) *cobra.Command {
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
	generator := bbd.NewBddGenerator()
	command := newGenerateBddCmd(generator)
	cmd.RootCmd.AddCommand(command)
}

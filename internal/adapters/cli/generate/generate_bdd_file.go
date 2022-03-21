package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateBddFileCmd struct {
	generateService ports.GenerateService
}

func NewGenerateBddFileCmd(generateService ports.GenerateService) cli.Cmd {
	return generateBddFileCmd{
		generateService: generateService,
	}
}

func (e generateBddFileCmd) Cmd() *cobra.Command {
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
			"gbddf",
			"generateBddFile",
			"generatebddfile",
			"generate-bdd-file",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := e.generateService.BddFile()
			fmt.Println(msg)
			return err
		},
	}
}

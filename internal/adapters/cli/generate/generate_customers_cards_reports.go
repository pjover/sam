package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type generateCustomersCardsReportsCmd struct {
	generateService ports.GenerateService
}

func NewGenerateCustomersCardsReportsCmd(generateService ports.GenerateService) cli.Cmd {
	return generateCustomersCardsReportsCmd{
		generateService: generateService,
	}
}

func (g generateCustomersCardsReportsCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "generaFitxesClients",
		Short:       "Genera les fitxes dels clients",
		Long:        "Genera els informe del les fitxes dels clients",
		Example:     `   generaFitxesClients      Genera les fitxes dels clients`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"gfic",
			"generafitxesclients",
			"genera-fitxes-clients",
			"generarFitxesClients",
			"generarfitxesclients",
			"generar-fitxes-clients",
			"gccards",
			"generateCustomersCards",
			"generatecustomerscards",
			"generate-customers-cards",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := g.generateService.CustomersCards()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

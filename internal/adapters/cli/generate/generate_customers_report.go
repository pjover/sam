package generate

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/cobra"
)

type generateCustomerReportCmd struct {
	generateService ports.GenerateService
}

func NewGenerateCustomerReportCmd(generateService ports.GenerateService) cli.Cmd {
	return generateCustomerReportCmd{
		generateService: generateService,
	}
}

func (e generateCustomerReportCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "generaInformeClients",
		Short:       "Genera l'informe dels clients",
		Long:        "Genera l'informe dels clients actius",
		Example:     `   generaInformeClients      Genera l'informe dels clients`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"ginfc",
			"generainfclients",
			"genera-inf-clients",
			"generarInfClients",
			"generarinfclients",
			"generar-inf-clients",
			"gcrep",
			"generateCustomersReport",
			"generatecustomersreport",
			"generate-customers-report",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := e.generateService.CustomerReport()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

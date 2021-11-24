package generate

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/generate"
	"github.com/pjover/sam/internal/generate/reports"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(NewGenerateCustomersReportCmd())
}

func NewGenerateCustomersReportCmd() *cobra.Command {
	return newGenerateCustomersReportCmd(reports.NewCustomersReportGenerator())
}

func newGenerateCustomersReportCmd(generator generate.Generator) *cobra.Command {
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
			msg, err := generator.Generate()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

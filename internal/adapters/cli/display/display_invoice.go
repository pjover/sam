package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
	"strings"
)

type displayInvoiceCmd struct {
	displayService ports.DisplayService
}

func NewDisplayInvoiceCmd(displayService ports.DisplayService) cli.Cmd {
	return displayInvoiceCmd{
		displayService: displayService,
	}
}

func (e displayInvoiceCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "mostraFactura codiFactura",
		Short:       "Mostra les dades d'una factura",
		Long:        "Mostra les dades d'uan factura indicant el seu codi",
		Example:     `   mostraFactura f-3945     Mostra les dades de la factura F-3945`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"mfac",
			"mostrafactura",
			"mostra-factura",
			"mostrarFactura",
			"mostrarfactura",
			"mostrar-factura",
			"dinv",
			"displayInvoice",
			"displayinvoice",
			"display-invoice",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := strings.ToUpper(args[0])
			msg, err := e.displayService.DisplayInvoice(id)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}

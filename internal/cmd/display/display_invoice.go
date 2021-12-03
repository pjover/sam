package display

import (
	"fmt"
	"strings"

	"github.com/pjover/sam/internal/display"
	"github.com/pjover/sam/internal/shared"

	"github.com/spf13/cobra"
)

func NewDisplayInvoiceCmd() *cobra.Command {
	return newDisplayInvoiceCmd(display.NewInvoiceDisplay())
}
func newDisplayInvoiceCmd(dsp display.Display) *cobra.Command {
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
		Args: shared.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			code := strings.ToUpper(args[0])
			msg, err := dsp.Display(code)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

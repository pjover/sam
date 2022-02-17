package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"strings"

	"github.com/spf13/cobra"
)

type editInvoiceCmd struct {
	editService ports.EditService
}

func NewEditInvoiceCmd(editService ports.EditService) cli.Cmd {
	return editInvoiceCmd{
		editService: editService,
	}
}

func (e editInvoiceCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "editaFactura codiFactura",
		Short:       "Edita les dades d'una factura",
		Long:        "Obri un navegador per a editar les dades d'una factura indicant el seu codi",
		Example:     `   editaFactura f-3945     Edita les dades de la factura F-3945`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"efac",
			"editafactura",
			"edita-factura",
			"editarFactura",
			"editarfactura",
			"editar-factura",
			"einv",
			"editInvoice",
			"editinvoice",
			"edit-invoice",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := strings.ToUpper(args[0])
			msg, err := e.editService.EditInvoice(id)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}

package edit

import (
	"strings"

	"github.com/pjover/sam/internal/edit"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/cobra"
)

func NewEditInvoiceCmd() *cobra.Command {
	return newEditInvoiceCmd(edit.NewInvoiceEditor())
}

func newEditInvoiceCmd(editor edit.Editor) *cobra.Command {
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
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			invoiceCode := strings.ToUpper(args[0])
			return editor.Edit(invoiceCode)
		},
	}
}

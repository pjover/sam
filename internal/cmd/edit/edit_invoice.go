package edit

import (
	"github.com/spf13/cobra"
	"sam/internal/cmd"
	"sam/internal/edit"
	"sam/internal/util"
	"strings"
)

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

func init() {
	editor := edit.NewInvoiceEditor()
	command := newEditInvoiceCmd(editor)
	cmd.RootCmd.AddCommand(command)
}

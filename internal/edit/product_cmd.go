package edit

import (
	"github.com/spf13/cobra"
	"sam/cmd"
	"sam/internal/util"
)

func newEditProductCmd(editor Editor) *cobra.Command {
	return &cobra.Command{
		Use:         "editaProducte codiProducte",
		Short:       "Edita les dades d'un producte",
		Long:        "Obri un navegador per a editar les dades d'un producte indicant el seu codi",
		Example:     `   editaProducte age     Edita les dades del producte AGE`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"epro",
			"editaproducte",
			"edita-producte",
			"editarProducte",
			"editarproducte",
			"editarproducte",
			"editProduct",
			"editproduct",
			"edit-product",
		},
		Args: util.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			productCode, err := util.ParseProductCode(args[0])
			if err != nil {
				return err
			}
			return editor.Edit(productCode)
		},
	}
}

func init() {
	editor := newProductEditor()
	command := newEditProductCmd(editor)
	cmd.RootCmd.AddCommand(command)
}

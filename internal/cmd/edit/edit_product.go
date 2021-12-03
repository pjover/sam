package edit

import (
	"github.com/pjover/sam/internal/edit"
	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/cobra"
)

func NewEditProductCmd() *cobra.Command {
	return newEditProductCmd(edit.NewProductEditor())
}
func newEditProductCmd(editor edit.Editor) *cobra.Command {
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
		Args: shared.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			productCode, err := shared.ParseProductCode(args[0])
			if err != nil {
				return err
			}
			return editor.Edit(productCode)
		},
	}
}

package adm

import (
	"github.com/spf13/cobra"
	"sam/internal/adm"
	"sam/internal/cmd"
)

var previousMonth bool
var nextMonth bool

func newDirectoryCmd(manager adm.DirectoryManager) *cobra.Command {
	return &cobra.Command{
		Use:   "directori",
		Short: "Crea el directori de treball",
		Long: `Crea el directori de treball per a les factures del mes
   - Si no s'especifica el mes, agafa l'actual
   - Si no existeix el directori, el crea
   - Actualitza la configuració amb el directori de treball
   - La configuració del directori romandrà activa fins que es torni a executar aquesta comanda de nou`,
		Example: `   dir             Crea el directori de treball per al mes actual
   directori -a    Crea el directori de treball per al mes anterior
   directori -s    Crea el directori de treball per al mes següent`,
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"directory",
			"dir",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return manager.Create(previousMonth, nextMonth)
		},
	}
}

func init() {
	manager := adm.NewDirectoryManager()
	command := newDirectoryCmd(manager)
	command.Flags().BoolVarP(&previousMonth, "anterior", "a", false, "Es treballa al mes anterior al mes actual")
	command.Flags().BoolVarP(&nextMonth, "seguent", "s", false, "Es treballa al mes següent al mes actual")
	cmd.RootCmd.AddCommand(command)
}

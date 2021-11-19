package cmd

import (
	"sam/adm"
	"sam/internal/util"

	"github.com/spf13/cobra"
)

var previousMonth bool
var nextMonth bool

var directoryCmd = &cobra.Command{
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
	Aliases:     []string{"directory", "dir"},
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := adm.Directories{Timer: util.SamTimeManager{}}
		return dir.CreateDirectory(previousMonth, nextMonth)
	},
}

func init() {
	rootCmd.AddCommand(directoryCmd)
	directoryCmd.Flags().BoolVarP(&previousMonth, "anterior", "a", false, "Es treballa al mes anterior al mes actual")
	directoryCmd.Flags().BoolVarP(&nextMonth, "seguent", "s", false, "Es treballa al mes següent al mes actual")
}

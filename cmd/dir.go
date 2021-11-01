package cmd

import (
	"sam/adm"
	"sam/util"

	"github.com/spf13/cobra"
)

var previousMonth bool
var nextMonth bool

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Crea el directori de treball",
	Long: `Crea el directori de treball per a les factures del mes
   - Si no s'especifica el mes, agafa l'actual
   - Si no existeix el directori, el crea
   - Actualitza la configuració amb el directori de treball
   - La configuració del directori romandrà activa fins que es torni a executar aquesta comanda de nou`,
	Example: `   dir       Crea el directori de treball per al mes actual
   dir -a    Crea el directori de treball per al mes anterior
   dir -s    Crea el directori de treball per al mes següent`,
	Annotations: map[string]string{"ADM": "Comandes d'administració"},
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := adm.Directories{Timer: util.SamTimeManager{}}
		return dir.CreateDirectory(previousMonth, nextMonth)
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
	dirCmd.Flags().BoolVarP(&previousMonth, "anterior", "a", false, "Es treballa al mes anterior al mes actual")
	dirCmd.Flags().BoolVarP(&nextMonth, "seguent", "s", false, "Es treballa al mes següent al mes actual")
}

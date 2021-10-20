/*
Copyright © 2021 Pere Jover

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"sam/amd"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := runDir(); err != nil {
			return err
		} else {
			return nil
		}
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)
	dirCmd.Flags().BoolVarP(&previousMonth, "anterior", "a", false, "Es treballa al mes anterior al mes actual")
	dirCmd.Flags().BoolVarP(&nextMonth, "seguent", "s", false, "Es treballa al mes següent al mes actual")
}

func runDir() error {
	return amd.Run(previousMonth, nextMonth)
}

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
	"fmt"

	"github.com/spf13/cobra"
)

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Crea el directori de treball",
	Long: `Crea el directori de treball per a les factures del mes
   - Comprova que al directori actual existeixi el fitxer de configuració general sam.yaml
     - si no el crea amb els valors per defecte i surt
   - Crea el directori
   - Si no s'especifica el mes, agafa l'actual
   - Crea el fitxer de configuració local sam-local.yaml al directori de treball
   - Canvia al nou directori`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dir called")
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//dirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	dirCmd.Flags().StringP("mes", "m", "", "Mes a facturar, si no s'especifica el mes, agafa l'actual")

}

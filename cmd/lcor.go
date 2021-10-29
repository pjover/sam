package cmd

import (
	"sam/adm"

	"github.com/spf13/cobra"
)

var ei1 bool
var ei2 bool
var ei3 bool

var lcorCmd = &cobra.Command{
	Use:   "lcor [-ei1 | ei2 | ei3]",
	Short: "Llista els correus electrònics",
	Long: `Llista els correus electrònics dels nins dels cursos indicats
   - Si no s'especifica el curs, es llisten tots
   - Es poden indicar varis cursos a l'hora`,
	Example: `   lcor       Llista els correus electrònics de tots els nins
   lcor -ei1    Llista els correus electrònics del curs EI1
   lcor -ei1 -ei2    Llista els correus electrònics del cursos EI1 i EI2`,
	Annotations: map[string]string{"ADM": "Comandes de llistats"},
	Aliases:     []string{"llista-correus"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return adm.ListEmails(ei1, ei2, ei3)
	},
}

func init() {
	rootCmd.AddCommand(lcorCmd)
	lcorCmd.Flags().BoolVarP(&ei1, "ei1", "1", false, "Educació infantil 1")
	lcorCmd.Flags().BoolVarP(&ei2, "ei2", "2", false, "Educació infantil 2")
	lcorCmd.Flags().BoolVarP(&ei3, "ei3", "3", false, "Educació infantil 3")
}

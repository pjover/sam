package list

import (
	"fmt"

	"github.com/pjover/sam/internal/cmd"
	"github.com/pjover/sam/internal/list"

	"github.com/spf13/cobra"
)

var ei1 bool
var ei2 bool
var ei3 bool

func init() {
	cmd.RootCmd.AddCommand(NewListMailsCmd())
}

func NewListMailsCmd() *cobra.Command {
	command := newListMailsCmd(list.NewListMails())
	command.Flags().BoolVarP(&ei1, "ei1", "1", false, "Educació infantil 1")
	command.Flags().BoolVarP(&ei2, "ei2", "2", false, "Educació infantil 2")
	command.Flags().BoolVarP(&ei3, "ei3", "3", false, "Educació infantil 3")
	return command
}

func newListMailsCmd(manager list.ListMails) *cobra.Command {
	return &cobra.Command{
		Use:   "llistaCorreus",
		Short: "Llista els correus electrònics",
		Long: `Llista els correus electrònics dels nins dels cursos indicats
   - Si no s'especifica el curs, es llisten tots
   - Es poden indicar varis cursos a l'hora`,
		Example: `   llistaCorreus              Llista els correus electrònics de tots els nins
   llistaCorreus -ei1         Llista els correus electrònics del curs EI1
   llistaCorreus -ei1 -ei2    Llista els correus electrònics del cursos EI1 i EI2`,
		Annotations: map[string]string{"ADM": "Comandes de llistats"},
		Aliases: []string{
			"lcor",
			"llistacorreus",
			"llista-correus",
			"llistarCorreus",
			"llistarcorreus",
			"llistar-correus",
			"lmai",
			"listMails",
			"listmails",
			"list-mails",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := manager.List(ei1, ei2, ei3)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

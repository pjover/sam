package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/cobra"
)

type listMailsCmd struct {
	listService ports.ListService
	ei1         bool
	ei2         bool
	ei3         bool
	i           bool
}

func NewListMailsCmd(listService ports.ListService) cli.Cmd {
	return listMailsCmd{
		listService: listService,
	}
}

func (l listMailsCmd) Cmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "llistaCorreus [--ei1 | --ei12 | --ei3] [--idioma]",
		Short: "Llista els correus electrònics",
		Long: `Llista els correus electrònics dels clients, d'un curs en concret o separats per idioma
   - Si no s'especifica res, es llisten tots
   - Si s'especifica un curs, únicament es llisten els correus electrònics dels clients amb infants a aquest curs
   - Si s'especifica un curs, únicament es llisten els correus electrònics dels clients amb infants a aquest curs
   - Es poden indicar varis cursos a l'hora`,
		Example: `   llistaCorreus              Llista els correus electrònics de tots els clients
   llistaCorreus --ei1         Llista els correus electrònics del curs EI1
   llistaCorreus --idioma      Llista els correus electrònics separats per idioma`,
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
			msg, err := l.callListMails()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
	command.Flags().BoolVarP(&l.ei1, "ei1", "1", false, "Educació infantil 1")
	command.Flags().BoolVarP(&l.ei2, "ei2", "2", false, "Educació infantil 2")
	command.Flags().BoolVarP(&l.ei3, "ei3", "3", false, "Educació infantil 3")
	command.Flags().BoolVarP(&l.i, "idioma", "i", false, "Tots els correus electrònics separats per idioma")
	return command
}

func (l listMailsCmd) callListMails() (string, error) {
	if l.i {
		return l.listService.ListMails("", true)
	} else {
		if l.ei1 {
			return l.listService.ListMails("EI_1", false)
		} else if l.ei2 {
			return l.listService.ListMails("EI_2", false)
		} else if l.ei3 {
			return l.listService.ListMails("EI_3", false)
		} else {
			return l.listService.ListMails("ALL", false)
		}
	}
}

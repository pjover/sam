package admin

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/cobra"
)

type directoryCmd struct {
	adminService ports.AdminService
}

var previousMonth bool
var nextMonth bool

func NewDirectoryCmd(adminService ports.AdminService) cli.Cmd {
	return directoryCmd{
		adminService: adminService,
	}
}

func (d directoryCmd) Cmd() *cobra.Command {
	cmd := d.newCmd()
	cmd.Flags().BoolVarP(&previousMonth, "anterior", "a", false, "Es treballa al mes anterior al mes actual")
	cmd.Flags().BoolVarP(&nextMonth, "seguent", "s", false, "Es treballa al mes següent al mes actual")
	return cmd
}

func (d directoryCmd) newCmd() *cobra.Command {
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
			msg, err := d.adminService.CreateDirectory(previousMonth, nextMonth)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

package admin

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type backupCmd struct {
	adminService ports.AdminService
}

func NewBackupCmd(adminService ports.AdminService) cli.Cmd {
	return backupCmd{
		adminService: adminService,
	}
}

func (b backupCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "copiaDeSeguretat",
		Short:       "Fa una còpia de seguretat de la base de dades",
		Long:        "Fa una còpia de seguretat de la base de dades a un fitxer a un directori especificat a la configuració",
		Example:     "   copiaDeSeguretat             Fa una còpia de seguretat de la base de dades",
		Annotations: map[string]string{"ADM": "Comandes d'administració"},
		Aliases: []string{
			"copiadeseguretat",
			"backup",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := b.adminService.Backup()
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		},
	}
}

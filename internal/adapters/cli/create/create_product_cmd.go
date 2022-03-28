package create

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
	"path"
)

type createProductCmd struct {
	createService ports.CreateService
	configService ports.ConfigService
	osService     ports.OsService
}

func NewCreateProductCmd(createService ports.CreateService, configService ports.ConfigService, osService ports.OsService) cli.Cmd {
	return createProductCmd{
		createService: createService,
		configService: configService,
		osService:     osService,
	}
}

func (e createProductCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "createProducte fitxerJsonDelProducte",
		Short:       "Crea un producte nou",
		Long:        "Crea un producte nou carregant la definició des d'un fitxer JSON situat el directori de treball",
		Example:     `   creaaProducte fitxerJsonDelProducte     Crea un producte nou`,
		Annotations: map[string]string{"CRE": "Comandes de creació d'entitats"},
		Aliases: []string{
			"cpro",
			"creaproducte",
			"crea-producte",
			"crearProducte",
			"crearproducte",
			"createProduct",
			"createproduct",
			"create-product",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			workingDir := e.configService.GetWorkingDirectory()
			filePath := path.Join(workingDir, args[0])

			content, err := e.osService.ReadFile(filePath)
			if err != nil {
				return err
			}
			msg, err := e.createService.CreateProduct(content)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}

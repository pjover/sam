package create

import (
	"encoding/json"
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/domain/model"
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

func (c createProductCmd) Cmd() *cobra.Command {
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
			product, err := c.loadProduct(args[0])
			if err != nil {
				return err
			}

			msg, err := c.createService.CreateProduct(product)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}

func (c createProductCmd) loadProduct(filename string) (product model.Product, err error) {
	workingDir := c.configService.GetWorkingDirectory()
	filePath := path.Join(workingDir, filename)

	content, err := c.osService.ReadFile(filePath)
	if err != nil {
		return model.Product{}, err
	}

	err = json.Unmarshal(content, &product)
	if err != nil {
		return model.Product{}, fmt.Errorf("error llegint el JSON del nou producte: %s", err)
	}
	return product, nil
}

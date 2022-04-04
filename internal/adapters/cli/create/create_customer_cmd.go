package create

import (
	"encoding/json"
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli"
	"github.com/pjover/sam/internal/adapters/cli/create/dto"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
	"path"
)

type createCustomerCmd struct {
	createService ports.CreateService
	configService ports.ConfigService
	osService     ports.OsService
}

func NewCreateCustomerCmd(createService ports.CreateService, configService ports.ConfigService, osService ports.OsService) cli.Cmd {
	return createCustomerCmd{
		createService: createService,
		configService: configService,
		osService:     osService,
	}
}

func (c createCustomerCmd) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:         "createClient fitxerJsonDelClient",
		Short:       "Crea un client nou",
		Long:        "Crea un client nou carregant la definició des d'un fitxer JSON situat el directori de treball",
		Example:     `   creaClient fitxerJsonDelClient     Crea un client nou`,
		Annotations: map[string]string{"CRE": "Comandes de creació d'entitats"},
		Aliases: []string{
			"ccli",
			"creaclient",
			"crea-client",
			"crearClient",
			"crearclient",
			"createCustomer",
			"createcustomer",
			"create-customer",
		},
		Args: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			customer, err := c.loadCustomer(args[0])
			if err != nil {
				return err
			}

			msg, err := c.createService.CreateCustomer(customer)
			if msg != "" {
				fmt.Println(msg)
			}
			return err
		},
	}
}

func (c createCustomerCmd) loadCustomer(filename string) (customer model.TransientCustomer, err error) {
	workingDir := c.configService.GetWorkingDirectory()
	filePath := path.Join(workingDir, filename)

	content, err := c.osService.ReadFile(filePath)
	if err != nil {
		return model.TransientCustomer{}, err
	}

	var customerDto dto.TransientCustomer
	err = json.Unmarshal(content, &customerDto)
	if err != nil {
		return model.TransientCustomer{}, fmt.Errorf("error llegint el JSON del nou client: %s", err)
	}

	customer = dto.TransientCustomerToModel(customerDto)
	return customer, nil
}

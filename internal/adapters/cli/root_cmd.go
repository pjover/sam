package cli

import (
	"github.com/pjover/sam/internal"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type CmdManager interface {
	AddCommand(cmd Cmd)
	AddTmpCommand(cmd *cobra.Command) //TODO remove
	Execute()
}

type cmdManager struct {
	configService ports.ConfigService
	rootCmd       *cobra.Command
}

func NewCmdManager(configService ports.ConfigService) CmdManager {
	// RootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "sam",
		Short: "Gestor de facturació de Hobbiton",
		Long: `Gestor de facturació de Hobbiton (+ info: https://github.com/pjover/sam)

El cicle normal es:
  1. insertaConsums: insertar consums
  2. llistaConsums: resum de consums per comprovar els totals
  3. facturaConsums: converteix els consums a factures
  4. generaFactures: generar els PDFs de les factures
  5. generaInformeMes: per crear el PDF amb el resum de les factures
  6. generaRebuts: per generar el fitxer de rebuts que després pujarem a CaixaBank
  7. pujar el fitxer de rebuts a CaixaBanc i signar l'operació amb CaixaSign
`,
		Version: internal.Version,
	}
	cobra.OnInitialize(configService.Init)
	return cmdManager{configService, rootCmd}
}

func (c cmdManager) GetRootCmd() *cobra.Command {
	return c.rootCmd
}

func (c cmdManager) AddCommand(cmd Cmd) {
	c.rootCmd.AddCommand(cmd.Cmd())
}

func (c cmdManager) AddTmpCommand(cmd *cobra.Command) { //TODO Remove
	c.rootCmd.AddCommand(cmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It only needs to happen once to the RootCmd.
func (c cmdManager) Execute() {
	cobra.CheckErr(c.rootCmd.Execute())
}

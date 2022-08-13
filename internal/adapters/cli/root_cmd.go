package cli

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/cobra"
)

type commandManager struct {
	configService ports.ConfigService
	rootCmd       *cobra.Command
}

func NewCommandManager(configService ports.ConfigService) ports.CommandManager {
	// RootCmd represents the base command when called without any subcommands
	title := fmt.Sprintf("sam v%s, Gestor de facturació de Hobbiton", domain.Version)
	rootCmd := &cobra.Command{
		Use:   "sam",
		Short: title,
		Long: title + ` (+ info: https://github.com/pjover/sam)

El cicle normal es:
  1. insertaConsums: insertar consums
  2. llistaConsums: resum de consums per comprovar els totals
  3. facturaConsums: converteix els consums a factures
  4. generaFactures: generar els PDFs de les factures
  5. generaInformeMes: per crear el PDF amb el resum de les factures
  6. generaRebuts: per generar el fitxer de rebuts que després pujarem a CaixaBank
  7. pujar el fitxer de rebuts a CaixaBanc i signar l'operació amb CaixaSign
`,
		Version: domain.Version,
	}
	cobra.OnInitialize(configService.Init)
	return commandManager{configService, rootCmd}
}

func (c commandManager) GetRootCmd() *cobra.Command {
	return c.rootCmd
}

func (c commandManager) AddCommand(cmd interface{}) {
	command := cmd.(Cmd)
	c.rootCmd.AddCommand(command.Cmd())
}

func (c commandManager) AddTmpCommand(cmd *cobra.Command) { //TODO Remove
	c.rootCmd.AddCommand(cmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// It only needs to happen once to the RootCmd.
func (c commandManager) Execute() {
	cobra.CheckErr(c.rootCmd.Execute())
}

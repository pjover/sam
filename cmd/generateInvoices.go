package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sam/adm"
)

func newGenerateInvoicesCmd(generateManager adm.GenerateManager) *cobra.Command {
	return &cobra.Command{
		Use:         "generaFactures",
		Short:       "Genera els PDFs de les factures del mes",
		Long:        "Genera els PDFs de les factures del mes al directori 'factures' dedins del directori de treball",
		Example:     `   generaFactures      Genera els PDFs de les factures del mes`,
		Annotations: map[string]string{"GEN": "Comandes de generació"},
		Aliases: []string{
			"gfacs",
			"generafactures", "genera-factures",
			"generarFactures", "generarfactures", "generar-factures",
			"ginvs",
			"generateInvoices", "generateinvoices", "generate-invoices",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generateManager.GenerateInvoices()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			return err
		},
	}
}

func init() {
	cmd := newGenerateInvoicesCmd(adm.NewGenerateManager())
	rootCmd.AddCommand(cmd)
}

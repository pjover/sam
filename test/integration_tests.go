package main

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/adapters/primary/cli"
	"github.com/pjover/sam/internal/cmd/adm"
	"github.com/pjover/sam/internal/cmd/consum"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/generate"
	"github.com/pjover/sam/internal/cmd/list"
	"github.com/pjover/sam/internal/cmd/search"
	"strings"

	"github.com/spf13/cobra"
)

type integrationTest struct {
	cmd  *cobra.Command
	args []string
}

var tests = []integrationTest{
	{adm.NewDirectoryCmd(), []string{}},
	{display.NewDisplayCustomerCmd(), []string{"181"}},
	{display.NewDisplayInvoiceCmd(), []string{"f-3945"}},
	{display.NewDisplayProductCmd(), []string{"age"}},
	//{edit.NewEditCustomerCmd(), []string{"246"}},
	//{edit.NewEditInvoiceCmd(), []string{"f-3945"}},
	//{edit.NewEditProductCmd(), []string{"age"}},
	{list.NewListChildrenCmd(), []string{}},
	{list.NewListConsumptionsCmd(), []string{}},
	{list.NewListConsumptionsCmd(), []string{"246"}},
	{list.NewListCustomersCmd(), []string{}},
	{list.NewListInvoicesCmd(), []string{}},
	{list.NewListMailsCmd(), []string{}},
	{list.NewListProductsCmd(), []string{}},
	{search.NewSearchCustomerCmd(), []string{"maria"}},
	{consum.NewInsertConsumptionsCmd(), []string{"2630", "1", "QME", "2", "MME", "1", "AGE"}},
	{consum.NewInsertConsumptionsCmd(), []string{"2640", "1", "QME", "1", "MME"}},
	{consum.NewInsertConsumptionsCmd(), []string{"2620", "1", "QME", "1", "MME"}},
	{consum.NewRectifyConsumptionsCmd(), []string{"2620", "1", "MME"}},
	{consum.NewBillConsumptionsCmd(), []string{}},
	{generate.NewGenerateSingleInvoiceCmd(), []string{"f-3945"}},
	{generate.NewGenerateBddCmd(), []string{}},
	{generate.NewGenerateCustomersReportCmd(), []string{}},
	//{generate.NewGenerateMonthInvoicesCmd(), []string{}},
	{generate.NewGenerateMonthReportCmd(), []string{}},
	{generate.NewGenerateProductsReportCmd(), []string{}},
}

func main() {
	cli.InitConfig()
	var errCount int
	var sb strings.Builder
	for _, test := range tests {
		isError, msg := run(test)
		if isError {
			errCount += 1
		}
		sb.WriteString(msg)
	}
	fmt.Print(sb.String())

	if errCount == 0 {
		fmt.Printf("All %d tests passed", len(tests))
	} else {
		fmt.Printf("%d of %d tests failed", errCount, len(tests))
	}
}

func run(test integrationTest) (bool, string) {
	_, err := executeCommand(test.cmd, test.args...)
	if err != nil {
		msg := fmt.Sprintf("🔴 sam %s %s >>> %s\n", test.cmd.Name(), strings.Join(test.args, " "), err)
		return true, msg
	} else {
		msg := fmt.Sprintf("🟢 sam %s %s\n", test.cmd.Name(), strings.Join(test.args, " "))
		return false, msg
	}
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	_, err = root.ExecuteC()
	return buf.String(), err
}

package main

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/cmd/adm"
	"github.com/pjover/sam/internal/cmd/display"
	"github.com/pjover/sam/internal/cmd/list"
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
}

func main() {
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
		msg := fmt.Sprintf("🔴 %s %s > %s\n", test.cmd.Name(), strings.Join(test.args, " "), err)
		return true, msg
	} else {
		msg := fmt.Sprintf("🟢 %s %s\n", test.cmd.Name(), strings.Join(test.args, " "))
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

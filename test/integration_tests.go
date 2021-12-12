package main

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"
	"log"
	goos "os"
	"os/exec"
	"strings"
)

var light_tests = [][]string{
	{"directory"},
	{"displayCustomer", "181"},
	{"displayInvoice", "f-3945"},
	{"displayProduct", "age"},
	{"listChildren"},
	{"listConsumptions"},
	{"listConsumptions", "246"},
	{"listCustomers"},
	{"listInvoices"},
	{"listMails"},
	{"listProducts"},
	{"searchCustomer", "maria"},
	{"insertConsumptions", "2630", "1", "QME", "2", "MME", "1", "AGE"},
	{"insertConsumptions", "2640", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2460", "1", "QME", "1", "MME"},
	{"rectifyConsumptions", "2460", "1", "MME"},
	{"billConsumptions"},
}

var heavy_tests = [][]string{
	{"backup"},
	{"editCustomer", "246"},
	{"editInvoice", "f-3945"},
	{"editProduct", "age"},
	{"generateSingleInvoice", "f-3945"},
	{"generateBdd"},
	{"generateCustomersReport"},
	{"generateMonthInvoices"},
	{"generateMonthReport"},
	{"generateProductsReport"},
}

var execManager = os.NewExecManager()

func main() {
	args := goos.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Required test type as arg: 'light', 'heavy' or 'all'")
	}
	switch args[0] {
	case "light":
		test(light_tests)
	case "heavy":
		test(heavy_tests)
	case "all":
		test(append(light_tests, heavy_tests...))
	}
}

func test(tests [][]string) {
	var errCount int
	var sb strings.Builder
	for _, args := range tests {
		isError, msg := run(args...)
		if isError {
			errCount += 1
		}
		sb.WriteString(msg)
	}
	fmt.Print(sb.String())

	if errCount == 0 {
		fmt.Printf("All %d tests passed\n", len(tests))
	} else {
		fmt.Printf("%d of %d tests failed\n", errCount, len(tests))
	}
}

func run(args ...string) (isError bool, msg string) {
	err := execManager.Run("sam", args...)
	if err != nil {
		exitError := err.(*exec.ExitError)
		return true, fmt.Sprintf("🔴 sam %s : %s\n", strings.Join(args, " "), exitError.Error())
	} else {
		return false, fmt.Sprintf("🟢 sam %s\n", strings.Join(args, " "))
	}
}

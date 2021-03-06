package main

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/os"
	"log"
	goos "os"
	"os/exec"
	"strings"
)

var lightTests = [][]string{
	{"displayCustomer", "263"},
	{"displayInvoice", "f-3945"},
	{"displayProduct", "age"},
	{"listChildren"},
	{"listConsumptions"},
	{"listConsumptions", "246"},
	{"listCustomers"},
	{"listInvoices"},
	{"listMails"},
	{"listProducts"},
	{"insertConsumptions", "2630", "1", "QME", "2", "MME", "1", "AGE"},
	{"insertConsumptions", "2640", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2460", "1", "QME", "1", "MME"},
	{"rectifyConsumptions", "2460", "1", "MME"},
	{"billConsumptions"},
	{"createCustomer", "nou_client.json"},
	{"generateProductsReport"},
	{"generateCustomersReport"},
	{"generateMonthReport"},
	{"generateBddFile"},
	{"generateSingleInvoice", "f-3945"},
}

var heavyTests = [][]string{
	{"createProduct", "nou_producte.json"},
	{"backup"},
	{"generateCustomerCards"},
	{"generateMonthInvoices"},
	{"editCustomer", "246"},
	{"editInvoice", "f-3945"},
	{"editProduct", "age"},
}

var cycleTest = [][]string{
	{"insertConsumptions", "2630", "1", "QME", "2", "MME", "1", "AGE"},
	{"insertConsumptions", "2640", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2460", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2470", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2480", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2460", "1", "QME", "1", "MME"},
	{"rectifyConsumptions", "2460", "1", "MME"},
	{"insertConsumptions", "2051", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2320", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2520", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2550", "1", "QME", "1", "MME"},
	{"listConsumptions"},
}

var osService = os.NewOsService()

func main() {
	args := goos.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Required test type as arg: 'light', 'heavy' or 'all'")
	}
	switch args[0] {
	case "light":
		test(lightTests)
	case "heavy":
		test(heavyTests)
	case "cycle":
		test(cycleTest)
	case "all":
		tst := append(lightTests, cycleTest...)
		tst = append(tst, heavyTests...)
		test(tst)
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
	err := osService.RunCommand("sam", args...)
	if err != nil {
		exitError := err.(*exec.ExitError)
		return true, fmt.Sprintf("???? sam %s : %s\n", strings.Join(args, " "), exitError.Error())
	} else {
		return false, fmt.Sprintf("???? sam %s\n", strings.Join(args, " "))
	}
}

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var tests_old = [][]string{
	{"directory"},
	{"displayCustomer", "246"},
	{"displayInvoice", "f-3945"},
	{"displayProduct", "age"},
	{"editCustomer", "246"},
	{"editInvoice", "f-3945"},
	{"editProduct", "age"},
	{"listChildren"},
	{"listConsumptions"},
	{"listCustomers"},
	{"listInvoices"},
	{"listMails"},
	{"listProducts"},
	{"searchCustomer", "maria"},
	{"generateSingleInvoice", "f-3945"},
	{"insertConsumptions", "2630", "1", "QME", "2", "MME", "1", "AGE"},
	{"insertConsumptions", "2640", "1", "QME", "1", "MME"},
	{"insertConsumptions", "2460", "1", "QME", "1", "MME"},
	{"rectifyConsumptions", "2460", "1", "MME"},
	{"billConsumptions"},
	{"generateBdd"},
	{"generateMonthInvoices"},
	{"generateCustomersReport"},
	{"generateProductsReport"},
	{"generateMonthReport"},
}

func main() {
	var errCount int
	for _, args := range tests_old {
		errCount += run_old("sam", args...)
	}
	if errCount == 0 {
		fmt.Printf("All %d tests passed", len(tests_old))
	} else {
		fmt.Printf("%d of %d tests failed", errCount, len(tests_old))
	}
}

func run_old(name string, args ...string) int {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		exitError := err.(*exec.ExitError)
		fmt.Println("🔴", name, strings.Join(args, " "), ":", exitError.Error())
		return 1
	} else {
		fmt.Println("🟢", name, strings.Join(args, " "))
		return 0
	}
}

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var tests = [][]string{
	{"billConsumptions"},
	{"directory"},
	{"displayCustomer", "246"},
	{"displayInvoice", "f-3945"},
	{"displayProduct", "age"},
	{"editCustomer", "246"},
	{"editInvoice", "f-3945"},
	{"editProduct", "age"},
	{"generateBdd"},
	{"generateInvoice", "f-3945"},
	{"insertConsumptions", "2460", "1", "QME"},
	{"listChildren"},
	{"listConsumptions"},
	{"listCustomers"},
	{"listInvoices"},
	{"listMails"},
	{"listProducts"},
	{"rectifyConsumptions", "2460", "1", "QME"},
	{"searchCustomer", "maria"},
}

func main() {
	var errCount int
	for _, args := range tests {
		errCount += run("sam", args...)
	}
	if errCount == 0 {
		fmt.Printf("All %d tests passed", len(tests))
	} else {
		fmt.Printf("%d of %d tests failed", errCount, len(tests))
	}
}

func run(name string, args ...string) int {
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

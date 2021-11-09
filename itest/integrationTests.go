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
	for _, args := range tests {
		run("sam", args...)
	}
}

func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Run()
	if err != nil {
		exitError := err.(*exec.ExitError)
		fmt.Println("🔴", name, strings.Join(args, " "), ":", exitError.Error())
	} else {
		fmt.Println("🟢", name, strings.Join(args, " "))
	}
}

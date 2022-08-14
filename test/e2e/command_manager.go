package e2e

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/domain/ports"
	"log"
	"strings"
)

type CommandType uint

const (
	InsertConsumptions CommandType = iota
	ListConsumptions
	GenerateMonthInvoices
	GenerateBddFile
	GenerateMonthReport
	GenerateCustomersReport
	GenerateCustomerCardsReport
)

var stringValues = []string{
	"insertConsumptions",
	"listConsumptions",
	"generateMonthInvoices",
	"generateBddFile",
	"generateMonthReport",
	"generateCustomersReport",
	"generateCustomerCardsReport",
}

func (c CommandType) String() string {
	return stringValues[c]
}

type Arguments []string

type Expected string

type Command struct {
	commandType CommandType
	arguments   Arguments
}

type e2eCommandManager struct {
	commands       []Command
	billingService ports.BillingService
}

func NewE2eCommandManager(billingService ports.BillingService) ports.CommandManager {
	return &e2eCommandManager{
		billingService: billingService,
	}
}

func (f *e2eCommandManager) AddCommand(cmd interface{}) {
	command := cmd.(Command)
	f.commands = append(f.commands, command)
}

func (f e2eCommandManager) Execute() []string {
	if len(f.commands) == 0 {
		return []string{"NO COMMANDS TO RUN"}
	}

	var result []string
	for i, command := range f.commands {
		fmt.Printf("%d. Running command '%s %s'\n", i+1, command.commandType.String(), strings.Join(command.arguments, " "))
		msg := f.run(command)
		fmt.Println(msg)
		result = append(result, msg)
	}
	return result
}

func (f e2eCommandManager) run(command Command) string {
	switch command.commandType {
	case InsertConsumptions:
		return f.runInsertConsumptions(command)
	}
	return "NO COMMAND RAN"
}

func (f e2eCommandManager) runInsertConsumptions(command Command) string {
	id, consumptions, _, err := billing.ParseConsumptionsArgs(command.arguments, "")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := f.billingService.InsertConsumptions(id, consumptions, "")
	if err != nil {
		return err.Error()
	}
	return msg
}
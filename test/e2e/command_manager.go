package e2e

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
	"strings"
)

type CommandType uint

const (
	Invalid CommandType = iota
	InsertConsumptions
	ListConsumptions
	GenerateMonthInvoices
	GenerateBddFile
	GenerateMonthReport
	GenerateCustomersReport
	GenerateCustomerCardsReport
)

var stringValues = []string{
	"",
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

type Command struct {
	commandType CommandType
	arguments   Arguments
}

func NewCommand(commandType CommandType, arguments []string) Command {
	return Command{
		commandType: commandType,
		arguments:   arguments,
	}
}

type FakeCommandManager struct {
	commands []Command
}

func NewFakeCommandManager() ports.CommandManager {
	return &FakeCommandManager{}
}

func (f *FakeCommandManager) AddCommand(cmd interface{}) {
	command := cmd.(Command)
	f.commands = append(f.commands, command)
	fmt.Printf("Added '%s' command (%d commands so far)\n", command.commandType.String(), len(f.commands))
}

func (f FakeCommandManager) Execute() {
	if len(f.commands) == 0 {
		fmt.Println("NO COMMANDS TO RUN")
		return
	}
	for i, command := range f.commands {
		fmt.Printf("%d. Running command '%s %s'\n", i+1, command.commandType.String(), strings.Join(command.arguments, " "))
		fmt.Println(f.run(command))
	}
}

func (f FakeCommandManager) run(command Command) string {
	switch command.commandType {
	case InsertConsumptions:
		return f.insertConsumptions(command.arguments)
	}
	return "NO COMMAND RAN"
}

func (f FakeCommandManager) insertConsumptions(arguments Arguments) string {
	return "  insertConsumptions FAKE OUTPUT"
}

package e2e

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"log"
	"strings"
)

type CommandType uint

const (
	InsertConsumptions CommandType = iota
	ListConsumptions
	BillConsumptions
	GenerateBddFile
)

var stringValues = []string{
	"insertConsumptions",
	"listConsumptions",
	"billConsumptions",
	"generateBddFile",
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
	commands        []Command
	configService   ports.ConfigService
	billingService  ports.BillingService
	listService     ports.ListService
	generateService ports.GenerateService
}

func NewE2eCommandManager(configService ports.ConfigService, billingService ports.BillingService, listService ports.ListService, generateService ports.GenerateService) ports.CommandManager {
	return &e2eCommandManager{
		configService:   configService,
		billingService:  billingService,
		listService:     listService,
		generateService: generateService,
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
	case ListConsumptions:
		return f.runListConsumptions()
	case BillConsumptions:
		return f.runBillConsumptions()
	case GenerateBddFile:
		return f.runGenerateBddFile()
	}
	return "NO COMMAND RAN"
}

func (f e2eCommandManager) runInsertConsumptions(command Command) string {
	id, consumptions, _, _, err := billing.ParseConsumptionsArgs(command.arguments, "2022-08", "")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := f.billingService.InsertConsumptions(id, consumptions, model.NewYearMonth(2022, 8), "")
	if err != nil {
		return err.Error()
	}
	return msg
}

func (f e2eCommandManager) runListConsumptions() string {
	msg, err := f.listService.ListConsumptions()
	if err != nil {
		return err.Error()
	}
	return msg
}

func (f e2eCommandManager) runBillConsumptions() string {
	msg, err := f.billingService.BillConsumptions()
	if err != nil {
		return err.Error()
	}
	return msg
}

func (f *e2eCommandManager) runGenerateBddFile() string {
	msg, err := f.generateService.BddFile()
	if err != nil {
		return err.Error()
	}
	return msg
}

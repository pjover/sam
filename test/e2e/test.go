package main

import (
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/test/fakes"
)

func main() {
	commandManager := endToEndTestDI()
	commandManager.Execute()
}

func endToEndTestDI() ports.CommandManager {
	commandManager := commandManager()

	return commandManager
}

func commandManager() ports.CommandManager {
	commandManager := fakes.NewFakeCommandManager()
	commandManager.AddCommand(fakes.NewCommand(fakes.InsertConsumptions, fakes.Arguments{"2630", "1", "QME", "2", "MME", "1", "AGE"}))
	commandManager.AddCommand(fakes.NewCommand(fakes.InsertConsumptions, fakes.Arguments{"2631", "1", "QME", "1", "MME", "1", "AGE"}))
	return commandManager
}

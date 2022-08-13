package e2e

import "github.com/pjover/sam/internal/domain/ports"

func InjectDependencies() ports.CommandManager {
	commandManager := commandManager()

	return commandManager
}

func commandManager() ports.CommandManager {
	commandManager := NewFakeCommandManager()
	commandManager.AddCommand(NewCommand(InsertConsumptions, Arguments{"2630", "1", "QME", "2", "MME", "1", "AGE"}))
	commandManager.AddCommand(NewCommand(InsertConsumptions, Arguments{"2631", "1", "QME", "1", "MME", "1", "AGE"}))
	return commandManager
}

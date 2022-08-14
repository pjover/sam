package e2e

func e2eMain(commands []Command) []string {
	commandManager := InjectDependencies()
	for _, command := range commands {
		commandManager.AddCommand(command)
	}
	return commandManager.Execute()
}

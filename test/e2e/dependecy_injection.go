package e2e

import (
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/domain/ports"
)

func InjectDependencies() ports.CommandManager {
	cmdManager := NewFakeCommandManager()

	configService := cfg.NewConfigService()           // TODO Fake it!
	dbService := mongo_db.NewDbService(configService) // TODO Fake it!
	osService := os.NewOsService()                    // TODO Fake it!

	addBillingCommands(cmdManager, configService, dbService, osService)

	return cmdManager
}

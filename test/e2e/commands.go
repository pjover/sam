package e2e

import (
	"github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/domain/ports"
	"log"
)

func (f FakeCommandManager) insertConsumptions(command Command) string {
	id, consumptions, _, err := billing.ParseConsumptionsArgs(command.arguments, "")
	if err != nil {
		log.Fatal(err)
	}

	service := command.service.(ports.BillingService)
	msg, err := service.InsertConsumptions(id, consumptions, "")
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

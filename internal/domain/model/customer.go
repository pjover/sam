package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/language"
	"strings"
	"time"
)

type Customer struct {
	Id            int
	Active        bool
	Children      []Child
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          string
	Language      language.Language
	ChangedOn     time.Time
}

func (c Customer) String() string {
	return fmt.Sprintf("%d  %-25s  %-55s  %s", c.Id, c.FirstAdultName(), c.ChildrenNamesWithId(", "), c.InvoiceHolder.PaymentInfoFmt())
}

func (c Customer) FirstAdult() Adult {
	for _, adult := range c.Adults {
		if adult.Role == adult_role.Mother {
			return adult
		}
	}
	return c.Adults[0]
}

func (c Customer) FirstAdultName() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s", adult.Name, adult.Surname)
}

func (c Customer) FirstAdultNameWithId() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s (%d)", adult.Name, adult.Surname, c.Id)
}

func (c Customer) ChildrenNamesWithId(joinWith string) string {
	var names []string
	for _, child := range c.Children {
		names = append(names, child.NameWithId())
	}
	return strings.Join(names, joinWith)
}

func (c Customer) ChildrenNames(joinWith string) string {
	var names []string
	for _, child := range c.Children {
		names = append(names, child.Name)
	}
	return strings.Join(names, joinWith)
}

func (c Customer) ChildrenNamesWithSurname(joinWith string) string {
	return fmt.Sprintf("%s %s", c.ChildrenNames(joinWith), c.Children[0].Surname)
}

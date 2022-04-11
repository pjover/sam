package model

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/language"
	"strings"
	"time"
)

type Customer struct {
	id            int
	active        bool
	children      []Child
	adults        []Adult
	invoiceHolder InvoiceHolder
	note          string
	language      language.Language
	changedOn     time.Time
}

func NewCustomer(
	id int,
	active bool,
	children []Child,
	adults []Adult,
	invoiceHolder InvoiceHolder,
	note string,
	language language.Language,
	changedOn time.Time,
) Customer {
	return Customer{
		id:            id,
		active:        active,
		children:      children,
		adults:        adults,
		invoiceHolder: invoiceHolder,
		note:          note,
		language:      language,
		changedOn:     changedOn,
	}
}

func (c Customer) Id() int {
	return c.id
}

func (c Customer) Active() bool {
	return c.active
}

func (c Customer) Children() []Child {
	return c.children
}

func (c Customer) Adults() []Adult {
	return c.adults
}

func (c Customer) InvoiceHolder() InvoiceHolder {
	return c.invoiceHolder
}

func (c Customer) Note() string {
	return c.note
}

func (c Customer) Language() language.Language {
	return c.language
}

func (c Customer) ChangedOn() time.Time {
	return c.changedOn
}

func (c Customer) String() string {
	return fmt.Sprintf("%d  %-25s  %-55s  %s", c.id, c.FirstAdultName(), c.ChildrenNamesWithId(", "), c.invoiceHolder.PaymentInfoFmt())
}

func (c Customer) FirstAdult() Adult {
	for _, adult := range c.Adults() {
		if adult.Role() == adult_role.Mother {
			return adult
		}
	}
	return c.Adults()[0]
}

func (c Customer) FirstAdultName() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s", adult.Name(), adult.Surname())
}

func (c Customer) FirstAdultNameWithId() string {
	adult := c.FirstAdult()
	return fmt.Sprintf("%s %s (%d)", adult.Name(), adult.Surname(), c.Id())
}

func (c Customer) ChildrenNamesWithId(joinWith string) string {
	var names []string
	for _, child := range c.Children() {
		names = append(names, child.NameWithId())
	}
	return strings.Join(names, joinWith)
}

func (c Customer) ChildrenNames(joinWith string) string {
	var names []string
	for _, child := range c.Children() {
		names = append(names, child.Name())
	}
	return strings.Join(names, joinWith)
}

func (c Customer) ChildrenNamesWithSurname(joinWith string) string {
	return fmt.Sprintf("%s %s", c.ChildrenNames(joinWith), c.Children()[0].Surname())
}

type TransientCustomer struct {
	Children      []TransientChild
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          string
	Language      language.Language
}

func (t TransientCustomer) Validate() error {

	for _, child := range t.Children {
		err := child.Validate()
		if err != nil {
			return err
		}
	}

	for _, adult := range t.Adults {
		err := adult.validate()
		if err != nil {
			return err
		}
	}

	err := t.InvoiceHolder.Validate()
	if err != nil {
		return err
	}

	if t.Language == language.Undefined {
		return errors.New("el llenguatge del client (Language) és incorrecte, ha d'esser CA, EN o ES")
	}

	return nil
}

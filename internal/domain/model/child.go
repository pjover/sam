package model

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"time"
)

type Child struct {
	id            int
	name          string
	surname       string
	secondSurname string
	taxID         TaxId
	birthDate     time.Time
	group         group_type.GroupType
	note          string
	active        bool
}

func NewChild(
	id int,
	name string,
	surname string,
	secondSurname string,
	taxID TaxId,
	birthDate time.Time,
	group group_type.GroupType,
	note string,
	active bool,
) Child {
	return Child{
		id:            id,
		name:          name,
		surname:       surname,
		secondSurname: secondSurname,
		taxID:         taxID,
		birthDate:     birthDate,
		group:         group,
		note:          note,
		active:        active,
	}
}

func (c Child) Id() int {
	return c.id
}

func (c Child) Name() string {
	return c.name
}

func (c Child) Surname() string {
	return c.surname
}

func (c Child) SecondSurname() string {
	return c.secondSurname
}

func (c Child) TaxID() TaxId {
	return c.taxID
}

func (c Child) BirthDate() time.Time {
	return c.birthDate
}

func (c Child) Group() group_type.GroupType {
	return c.group
}

func (c Child) Note() string {
	return c.note
}

func (c Child) Active() bool {
	return c.active
}

func (c Child) validate() error {
	if c.name == "" {
		return errors.New("el nom de l'infant (Name) no pot estar buit")
	}

	if c.surname == "" {
		return errors.New("el primer llinatge de l'infant (Surname) no pot estar buit")
	}

	var emptyBirthDate = time.Time{}
	if c.birthDate == emptyBirthDate {
		return errors.New("la data de naixement de l'infant (BirthDate) no pot estar buida")
	}

	if c.group == group_type.Undefined {
		return errors.New("el grup de l'infant (Group) és incorrecte, ha d'esser EI_1, EI_2 o EI_3")
	}

	return nil
}

func (c Child) String() string {
	return fmt.Sprintf("%d  %-30s  %s  %s", c.id, c.NameAndSurname(), c.group, c.birthDate.Format(domain.YearMonthDayLayout))
}

func (c Child) NameAndSurname() string {
	return fmt.Sprintf("%s %s", c.name, c.surname)
}

func (c Child) NameWithId() string {
	return fmt.Sprintf("%s %s (%d)", c.name, c.surname, c.id)
}

type TransientChild struct {
	Name          string
	Surname       string
	SecondSurname string
	TaxID         TaxId
	BirthDate     time.Time
	Group         group_type.GroupType
	Note          string
}

package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"time"
)

type Child struct {
	Id            int
	Name          string
	Surname       string
	SecondSurname string
	TaxID         string
	BirthDate     time.Time
	Group         group_type.GroupType
	Note          string
	Active        bool
}

func (c Child) String() string {
	return fmt.Sprintf("%d  %-30s  %s  %s", c.Id, c.NameAndSurname(), c.Group, c.BirthDate.Format(domain.YearMonthDayLayout))
}

func (c Child) NameAndSurname() string {
	return fmt.Sprintf("%s %s", c.Name, c.Surname)
}

func (c Child) NameWithId() string {
	return fmt.Sprintf("%s %s (%d)", c.Name, c.Surname, c.Id)
}

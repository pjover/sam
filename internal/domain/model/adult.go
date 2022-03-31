package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"time"
)

type Adult struct {
	Name             string
	Surname          string
	SecondSurname    string
	TaxID            string
	Role             adult_role.AdultRole
	Address          Address
	Email            string
	MobilePhone      string
	HomePhone        string
	GrandMotherPhone string
	GrandParentPhone string
	WorkPhone        string
	BirthDate        time.Time
	Nationality      string
}

func (a Adult) MobilePhoneFmt() string {
	if len(a.MobilePhone) != 9 {
		return a.MobilePhone
	}
	return fmt.Sprintf(
		"%s %s %s",
		a.MobilePhone[0:3],
		a.MobilePhone[3:6],
		a.MobilePhone[6:9],
	)
}

func (a Adult) NameAndSurname() string {
	return fmt.Sprintf("%s %s", a.Name, a.Surname)
}

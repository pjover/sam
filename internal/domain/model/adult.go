package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"time"
)

type Adult struct {
	name             string
	surname          string
	secondSurname    string
	taxID            TaxId
	role             adult_role.AdultRole
	address          Address
	email            string
	mobilePhone      string
	homePhone        string
	grandMotherPhone string
	grandParentPhone string
	workPhone        string
	birthDate        time.Time
	nationality      Nationality
}

func NewAdult(
	name string,
	surname string,
	secondSurname string,
	taxID TaxId,
	role adult_role.AdultRole,
	address Address,
	email string,
	mobilePhone string,
	homePhone string,
	grandMotherPhone string,
	grandParentPhone string,
	workPhone string,
	birthDate time.Time,
	nationality Nationality,
) Adult {
	return Adult{
		name:             name,
		surname:          surname,
		secondSurname:    secondSurname,
		taxID:            taxID,
		role:             role,
		address:          address,
		email:            email,
		mobilePhone:      mobilePhone,
		homePhone:        homePhone,
		grandMotherPhone: grandMotherPhone,
		grandParentPhone: grandParentPhone,
		workPhone:        workPhone,
		birthDate:        birthDate,
		nationality:      nationality,
	}
}

func (a Adult) Name() string {
	return a.name
}

func (a Adult) Surname() string {
	return a.surname
}

func (a Adult) SecondSurname() string {
	return a.secondSurname
}

func (a Adult) TaxID() TaxId {
	return a.taxID
}

func (a Adult) Role() adult_role.AdultRole {
	return a.role
}

func (a Adult) Address() Address {
	return a.address
}

func (a Adult) Email() string {
	return a.email
}

func (a Adult) MobilePhone() string {
	return a.mobilePhone
}

func (a Adult) HomePhone() string {
	return a.homePhone
}

func (a Adult) GrandMotherPhone() string {
	return a.grandMotherPhone
}

func (a Adult) GrandParentPhone() string {
	return a.grandParentPhone
}

func (a Adult) WorkPhone() string {
	return a.workPhone
}

func (a Adult) BirthDate() time.Time {
	return a.birthDate
}

func (a Adult) Nationality() Nationality {
	return a.nationality
}

func (a Adult) MobilePhoneFmt() string {
	if len(a.mobilePhone) != 9 {
		return a.mobilePhone
	}
	return fmt.Sprintf(
		"%s %s %s",
		a.mobilePhone[0:3],
		a.mobilePhone[3:6],
		a.mobilePhone[6:9],
	)
}

func (a Adult) NameAndSurname() string {
	return fmt.Sprintf("%s %s", a.name, a.surname)
}

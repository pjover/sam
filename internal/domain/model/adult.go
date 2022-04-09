package model

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"net/mail"
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

func (a Adult) validate() error {
	if a.name == "" {
		return errors.New("el nom de l'adult (Name) no pot estar buit")
	}

	if a.surname == "" {
		return errors.New("el primer llinatge de l'adult (Surname) no pot estar buit")
	}

	if a.taxID == emptyTaxId {
		return errors.New("el DNI/NIE de l'adult (TaxId) no pot estar buit")
	}

	if a.role == adult_role.Invalid {
		return fmt.Errorf("el rol de l'adult (Role) no és vàlid, ha d'esser MOTHER, FATHER o TUTOR")
	}

	_, err := mail.ParseAddress(a.email)
	if err != nil {
		return fmt.Errorf("el correu electrònic de l'adult (Email) no és vàlid")
	}

	err = a.address.validate()
	if err != nil {
		return err
	}

	var emptyBirthDate = time.Time{}
	if a.birthDate == emptyBirthDate {
		return errors.New("la data de naixement de l'adult (BirthDate) no pot estar buida")
	}

	if a.nationality == emptyNationality {
		return errors.New("la nacionalitat de l'adult (Nationality) no pot estar buida")
	}

	return nil
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

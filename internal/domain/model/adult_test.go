package model

import (
	"errors"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testDate = time.Date(2019, 5, 25, 0, 0, 0, 0, time.UTC)

func TestAdult_validate(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			"Valid",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			nil,
		},
		{
			"Empty name",
			fields{
				name:    "",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el nom de l'adult (Name) no pot estar buit"),
		},
		{
			"Empty surname",
			fields{
				name:    "Some name",
				surname: "",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el primer llinatge de l'adult (Surname) no pot estar buit"),
		},
		{
			"Empty TaxId",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty(""),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el DNI/NIE de l'adult (TaxId) no pot estar buit"),
		},
		{
			"Invalid role",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el rol de l'adult (Role) no és vàlid, ha d'esser MOTHER, FATHER o TUTOR"),
		},
		{
			"Invalid email",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some_email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el correu electrònic de l'adult (Email) no és vàlid"),
		},
		{
			"Invalid address name",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"7777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("el codi postal (ZipCode) ha de tenir 5 números"),
		},
		{
			"Empty BirthDate",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   time.Time{},
				nationality: NewNationalityOrEmpty("US"),
			},
			errors.New("la data de naixement de l'adult (BirthDate) no pot estar buida"),
		},
		{
			"Empty Nationality",
			fields{
				name:    "Some name",
				surname: "Some surname",
				taxID:   NewTaxIdOrEmpty("60235657Z"),
				role:    adult_role.Mother,
				email:   "some@email.com",
				address: NewAddress(
					"Some street",
					"77777",
					"Some city",
					"Some state",
				),
				birthDate:   testDate,
				nationality: NewNationalityOrEmpty(""),
			},
			errors.New("la nacionalitat de l'adult (Nationality) no pot estar buida"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Adult{
				name:             tt.fields.name,
				surname:          tt.fields.surname,
				secondSurname:    tt.fields.secondSurname,
				taxID:            tt.fields.taxID,
				role:             tt.fields.role,
				address:          tt.fields.address,
				email:            tt.fields.email,
				mobilePhone:      tt.fields.mobilePhone,
				homePhone:        tt.fields.homePhone,
				grandMotherPhone: tt.fields.grandMotherPhone,
				grandParentPhone: tt.fields.grandParentPhone,
				workPhone:        tt.fields.workPhone,
				birthDate:        tt.fields.birthDate,
				nationality:      tt.fields.nationality,
			}
			got := sut.validate()
			assert.Equal(t, tt.want, got)
		})
	}
}

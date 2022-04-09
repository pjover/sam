package model

import (
	"errors"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChild_validate(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"Valid",
			fields{
				name:      "Some name",
				surname:   "Some surname",
				birthDate: testDate,
				group:     group_type.Ei1,
			},
			errors.New("el nom de l'infant (Name) no pot estar buit"),
		},
		{
			"Empty name",
			fields{
				name:      "Some name",
				surname:   "Some surname",
				birthDate: testDate,
				group:     group_type.Ei1,
			},
			errors.New("el primer llinatge de l'infant (Surname) no pot estar buit"),
		},
		{
			"Empty BirthDate",
			fields{
				name:      "Some name",
				surname:   "Some surname",
				birthDate: testDate,
				group:     group_type.Ei1,
			},
			errors.New("la data de naixement de l'infant (BirthDate) no pot estar buida"),
		},
		{
			"Empty Group",
			fields{
				name:      "Some name",
				surname:   "Some surname",
				birthDate: testDate,
				group:     group_type.Ei1,
			},
			errors.New("el grup de l'infant (Group) és incorrecte, ha d'esser EI_1, EI_2 o EI_3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Child{
				id:            tt.fields.id,
				name:          tt.fields.name,
				surname:       tt.fields.surname,
				secondSurname: tt.fields.secondSurname,
				taxID:         tt.fields.taxID,
				birthDate:     tt.fields.birthDate,
				group:         tt.fields.group,
				note:          tt.fields.note,
				active:        tt.fields.active,
			}
			got := sut.validate()
			assert.Equal(t, tt.wantErr, got)
		})
	}
}

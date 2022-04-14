package model

import (
	"errors"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransientChild_validate(t *testing.T) {
	type fields struct {
		Name          string
		Surname       string
		SecondSurname string
		TaxID         TaxId
		BirthDate     time.Time
		Group         group_type.GroupType
		Note          string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"Valid",
			fields{
				Name:      "Some name",
				Surname:   "Some surname",
				BirthDate: testDate,
				Group:     group_type.Ei1,
			},
			nil,
		},
		{
			"Empty Name",
			fields{
				Name:      "",
				Surname:   "Some surname",
				BirthDate: testDate,
				Group:     group_type.Ei1,
			},
			errors.New("el nom de l'infant (Name) no pot estar buit"),
		},
		{
			"Empty Surname",
			fields{
				Name:      "Some name",
				Surname:   "",
				BirthDate: testDate,
				Group:     group_type.Ei1,
			},
			errors.New("el primer llinatge de l'infant (Surname) no pot estar buit"),
		},
		{
			"Empty BirthDate",
			fields{
				Name:      "Some name",
				Surname:   "Some surname",
				BirthDate: time.Time{},
				Group:     group_type.Ei1,
			},
			errors.New("la data de naixement de l'infant (BirthDate) no pot estar buida"),
		},
		{
			"Empty Group",
			fields{
				Name:      "Some name",
				Surname:   "Some surname",
				BirthDate: testDate,
				Group:     group_type.Undefined,
			},
			errors.New("el grup de l'infant (Group) és incorrecte, ha d'esser EI_1, EI_2 o EI_3"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			sut := TransientChild{
				Name:          tt.fields.Name,
				Surname:       tt.fields.Surname,
				SecondSurname: tt.fields.SecondSurname,
				TaxID:         tt.fields.TaxID,
				BirthDate:     tt.fields.BirthDate,
				Group:         tt.fields.Group,
				Note:          tt.fields.Note,
			}
			got := sut.Validate()
			assert.Equal(t, tt.wantErr, got)
		})
	}
}

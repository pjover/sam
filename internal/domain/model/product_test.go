package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_validate(t *testing.T) {
	type fields struct {
		id            string
		name          string
		shortName     string
		price         float64
		taxPercentage float64
		isSubsidy     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			"Valid",
			fields{
				"TST",
				"Some name",
				"ShortName",
				10,
				0.16,
				false,
			},
			nil,
		},
		{
			"Empty ID",
			fields{
				"",
				"Some name",
				"ShortName",
				10,
				0.16,
				false,
			},
			errors.New("el id del producte no pot estar buit"),
		},
		{
			"ID with lower case",
			fields{
				"abc",
				"Some name",
				"ShortName",
				10,
				0.16,
				false,
			},
			errors.New("el id del producte ha d'estar en majúscules"),
		},
		{
			"ID with less than 3 characters",
			fields{
				"AB",
				"Some name",
				"ShortName",
				10,
				0.16,
				false,
			},
			errors.New("el id del producte ha de tenir 3 caràcters"),
		},
		{
			"ID with more than 3 characters",
			fields{
				"ABCD",
				"Some name",
				"ShortName",
				10,
				0.16,
				false,
			},
			errors.New("el id del producte ha de tenir 3 caràcters"),
		},
		{
			"Empty shortName",
			fields{
				"ABC",
				"Some name",
				"",
				10,
				0.16,
				false,
			},
			errors.New("el nom curt del producte (ShortName) no pot estar buit"),
		},
		{
			"Empty name",
			fields{
				"ABC",
				"",
				"ShortName",
				10,
				0.16,
				false,
			},
			errors.New("el nom del producte (Name) no pot estar buit"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Product{
				id:            tt.fields.id,
				name:          tt.fields.name,
				shortName:     tt.fields.shortName,
				price:         tt.fields.price,
				taxPercentage: tt.fields.taxPercentage,
				isSubsidy:     tt.fields.isSubsidy,
			}
			got := sut.Validate()
			assert.Equal(t, tt.want, got)
		})
	}
}

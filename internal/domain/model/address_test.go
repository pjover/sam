package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress_validate(t *testing.T) {
	type fields struct {
		street  string
		zipCode string
		city    string
		state   string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			"All empty",
			fields{
				"",
				"",
				"",
				"",
			},
			nil,
		},
		{
			"Street empty, with ZIP",
			fields{
				"",
				"Not empty",
				"",
				"",
			},
			errors.New("si el carrer (Street) és buit, la resta de camps han d'esser buits"),
		},
		{
			"Street empty, with city",
			fields{
				"",
				"",
				"Not empty",
				"",
			},
			errors.New("si el carrer (Street) és buit, la resta de camps han d'esser buits"),
		},
		{
			"Street empty, with state",
			fields{
				"",
				"",
				"",
				"Not empty",
			},
			errors.New("si el carrer (Street) és buit, la resta de camps han d'esser buits"),
		},
		{
			"Zip code empty",
			fields{
				"Some street",
				"",
				"Some city",
				"Some state",
			},
			errors.New("el codi postal (ZipCode) no pot estar buit"),
		},
		{
			"Zip code with more numbers",
			fields{
				"Some street",
				"777777",
				"Some city",
				"Some state",
			},
			errors.New("el codi postal (ZipCode) ha de tenir 5 números"),
		},
		{
			"Zip code with less numbers",
			fields{
				"Some street",
				"7777",
				"Some city",
				"Some state",
			},
			errors.New("el codi postal (ZipCode) ha de tenir 5 números"),
		},
		{
			"Zip code with letters",
			fields{
				"Some street",
				"7777A",
				"Some city",
				"Some state",
			},
			errors.New("el codi postal (ZipCode) només pot tenir números"),
		},
		{
			"City empty",
			fields{
				"Some street",
				"77777",
				"",
				"Some state",
			},
			errors.New("la ciutat (City) no pot estar buida"),
		},
		{
			"State empty",
			fields{
				"Some street",
				"77777",
				"Some city",
				"",
			},
			errors.New("el estat o regió (State) no pot estar buit"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := Address{
				street:  tt.fields.street,
				zipCode: tt.fields.zipCode,
				city:    tt.fields.city,
				state:   tt.fields.state,
			}
			got := sut.validate()
			assert.Equal(t, tt.want, got)
		})
	}
}

package model

import (
	"errors"
	"github.com/biter777/countries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNationality(t *testing.T) {
	tests := []struct {
		name       string
		alpha2Code string
		want       Nationality
		wantErr    error
	}{
		{
			name:       "ok",
			alpha2Code: "UK",
			want:       Nationality{countries.GBR},
			wantErr:    nil,
		},
		{
			name:       "empty",
			alpha2Code: "",
			want:       Nationality{},
			wantErr:    errors.New("la nacionalitat '' no és vàlida"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNationality(tt.alpha2Code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNationality_String(t *testing.T) {
	tests := []struct {
		name        string
		nationality Nationality
		want        string
	}{
		{
			name:        "ok",
			nationality: Nationality{countries.GBR},
			want:        "GB",
		},
		{
			name:        "empty",
			nationality: Nationality{countries.Unknown},
			want:        "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.nationality.String())
		})
	}
}

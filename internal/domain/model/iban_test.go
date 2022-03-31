package model

import (
	"fmt"
	"github.com/biter777/countries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_extractCountryCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    countries.CountryCode
		wantErr error
	}{
		{
			name: "All strings",
			code: "ES2830668859978258529057",
			want: countries.ES,
		},
		{
			name: "Upper case",
			code: "es2830668859978258529057",
			want: countries.ES,
		},
		{
			name:    "Valid ISO 3166-1 alpha-2 country",
			code:    "xy2830668859978258529057",
			want:    countries.Unknown,
			wantErr: fmt.Errorf("'xy' is an invalid ISO 3166-1 alpha-2 country"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractCountryCode(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_isNumber(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "All numbers",
			text: "2830668859978258529057",
			want: true,
		},
		{
			name: "With letter numbers",
			text: "283066885997825852d9057",
			want: false,
		},
		{
			name: "With letters numbers",
			text: "2830668859978s5852d9057",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumber(tt.text)
			assert.Equal(t, tt.want, got)
		})
	}
}

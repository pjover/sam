package model

import (
	"errors"
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
			wantErr: errors.New("'xy' is an invalid ISO 3166-1 alpha-2 country"),
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

func Test_extractCheckDigits(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		code    string
		want    string
		wantErr error
	}{
		{
			name: "Two digits",
			code: "ES2830668859978258529057",
			want: "28",
		},
		{
			name:    "Not two digits",
			code:    "ESC830668859978258529057",
			want:    "",
			wantErr: errors.New("'C8' is an invalid two numbers IBAN check digits"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractCheckDigits(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_extractBban(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		code    string
		want    string
		wantErr error
	}{
		{
			name: "All digits",
			code: "ES2830668859978258529057",
			want: "30668859978258529057",
		},
		{
			name: "Valid code",
			code: "GB98MIDL07009312345678",
			want: "MIDL07009312345678",
		},

		{
			name:    "Invalid code",
			code:    "GB98MIDÇ07009312345678",
			want:    "",
			wantErr: errors.New("'MIDÇ07009312345678' is an invalid IBAN Basic Bank Account Number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractBban(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

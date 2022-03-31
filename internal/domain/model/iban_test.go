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

func Test_prepareCode(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "Without separators",
			code: "GB98MIDL07009312345678",
			want: "GB98MIDL07009312345678",
		},
		{
			name: "With space separators",
			code: "GB98 MIDL 0700 9312 3456 78",
			want: "GB98MIDL07009312345678",
		},
		{
			name: "Without hypens separators",
			code: "GB98-MIDL-0700-9312-3456-78",
			want: "GB98MIDL07009312345678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prepareCode(tt.code)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewBankAccount(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    IBAN
		wantErr error
	}{
		{
			name: "All digits",
			code: "ES2830668859978258529057",
			want: IBAN{
				countryCode: countries.ESP,
				checkDigits: "28",
				bban:        "30668859978258529057",
			},
			wantErr: nil,
		},
		{
			name: "With letters",
			code: "GB98MIDL07009312345678",
			want: IBAN{
				countryCode: countries.GBR,
				checkDigits: "98",
				bban:        "MIDL07009312345678",
			},
			wantErr: nil,
		},
		{
			name:    "With wrong countryCode",
			code:    "xy98MIDL07009312345678",
			want:    IBAN{},
			wantErr: errors.New("'xy' is an invalid ISO 3166-1 alpha-2 country"),
		},
		{
			name:    "With wrong checkDigits",
			code:    "GBx9MIDL07009312345678",
			want:    IBAN{},
			wantErr: errors.New("'x9' is an invalid two numbers IBAN check digits"),
		},
		{
			name:    "With wrong BBAN",
			code:    "GB98MIDÖL07009312345678",
			want:    IBAN{},
			wantErr: errors.New("'MIDÖL07009312345678' is an invalid IBAN Basic Bank Account Number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBankAccount(tt.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

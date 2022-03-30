package sequence_type

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSequenceType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  SequenceType
	}{
		{
			name:  "Invalid",
			value: "",
			want:  Invalid,
		},
		{
			name:  "StandardInvoice",
			value: "STANDARD_INVOICE",
			want:  StandardInvoice,
		},
		{
			name:  "SpecialInvoice",
			value: "SPECIAL_INVOICE",
			want:  SpecialInvoice,
		},
		{
			name:  "RectificationInvoice",
			value: "RECTIFICATION_INVOICE",
			want:  RectificationInvoice,
		},
		{
			name:  "Customer",
			value: "CUSTOMER",
			want:  Customer,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSequenceType(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSequenceType_Format(t *testing.T) {
	tests := []struct {
		name string
		s    SequenceType
		want string
	}{
		{
			name: "Invalid",
			s:    Invalid,
			want: "Indefinit",
		},
		{
			name: "StandardInvoice",
			s:    StandardInvoice,
			want: "Factura (rebut)",
		},
		{
			name: "SpecialInvoice",
			s:    SpecialInvoice,
			want: "Factura (no rebut)",
		},
		{
			name: "RectificationInvoice",
			s:    RectificationInvoice,
			want: "Rectificació",
		},
		{
			name: "Customer",
			s:    Customer,
			want: "Client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Format()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSequenceType_Prefix(t *testing.T) {
	tests := []struct {
		name string
		s    SequenceType
		want string
	}{

		{
			name: "Invalid",
			s:    Invalid,
			want: "",
		},
		{
			name: "StandardInvoice",
			s:    StandardInvoice,
			want: "F",
		},
		{
			name: "SpecialInvoice",
			s:    SpecialInvoice,
			want: "X",
		},
		{
			name: "RectificationInvoice",
			s:    RectificationInvoice,
			want: "R",
		},
		{
			name: "Customer",
			s:    Customer,
			want: "C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Prefix()
			assert.Equal(t, tt.want, got)
		})
	}
}

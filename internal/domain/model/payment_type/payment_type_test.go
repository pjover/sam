package payment_type

import (
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPaymentType(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  PaymentType
	}{
		{
			name:  "Invalid",
			value: "Anything",
			want:  Invalid,
		},
		{
			name:  "BankDirectDebit",
			value: "BANK_DIRECT_DEBIT",
			want:  BankDirectDebit,
		},
		{
			name:  "BankTransfer",
			value: "BANK_TRANSFER",
			want:  BankTransfer,
		},
		{
			name:  "Voucher",
			value: "VOUCHER",
			want:  Voucher,
		},
		{
			name:  "Cash",
			value: "CASH",
			want:  Cash,
		},
		{
			name:  "Rectification",
			value: "RECTIFICATION",
			want:  Rectification,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPaymentType(tt.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPaymentType_Format(t *testing.T) {
	tests := []struct {
		name string
		p    PaymentType
		want string
	}{
		{
			name: "Invalid",
			p:    Invalid,
			want: "Indefinit",
		},
		{
			name: "BankDirectDebit",
			p:    BankDirectDebit,
			want: "Rebut",
		},
		{
			name: "BankTransfer",
			p:    BankTransfer,
			want: "Tranferència",
		},
		{
			name: "Voucher",
			p:    Voucher,
			want: "Xec escoleta",
		},
		{
			name: "Cash",
			p:    Cash,
			want: "Efectiu",
		},
		{
			name: "Rectification",
			p:    Rectification,
			want: "Rectificació",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Format()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPaymentType_SequenceType(t *testing.T) {
	tests := []struct {
		name string
		p    PaymentType
		want sequence_type.SequenceType
	}{
		{
			name: "Invalid",
			p:    Invalid,
			want: sequence_type.Invalid,
		},
		{
			name: "BankDirectDebit",
			p:    BankDirectDebit,
			want: sequence_type.StandardInvoice,
		},
		{
			name: "BankTransfer",
			p:    BankTransfer,
			want: sequence_type.SpecialInvoice,
		},
		{
			name: "Voucher",
			p:    Voucher,
			want: sequence_type.SpecialInvoice,
		},
		{
			name: "Cash",
			p:    Cash,
			want: sequence_type.SpecialInvoice,
		},
		{
			name: "Rectification",
			p:    Rectification,
			want: sequence_type.RectificationInvoice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.SequenceType()
			assert.Equal(t, tt.want, got)
		})
	}
}

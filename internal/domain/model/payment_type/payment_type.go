package payment_type

import (
	"github.com/pjover/sam/internal/domain/model/sequence_type"
	"strings"
)

type PaymentType uint

const (
	Invalid PaymentType = iota
	BankDirectDebit
	BankTransfer
	Voucher
	Cash
	Rectification
)

var stringValues = []string{
	"",
	"BANK_DIRECT_DEBIT",
	"BANK_TRANSFER",
	"VOUCHER",
	"CASH",
	"RECTIFICATION",
}

func (p PaymentType) String() string {
	return formatValues[p]
}

func NewPaymentType(value string) PaymentType {
	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return PaymentType(i)
		}
	}
	return Invalid
}

var formatValues = []string{
	"Indefinit",
	"Rebut",
	"Tranferència",
	"Xec escoleta",
	"Efectiu",
	"Rectificació",
}

func (p PaymentType) Format() string {
	return formatValues[p]
}

func (p PaymentType) SequenceType() sequence_type.SequenceType {
	switch p {
	case BankDirectDebit:
		return sequence_type.StandardInvoice
	case BankTransfer:
		return sequence_type.SpecialInvoice
	case Voucher:
		return sequence_type.SpecialInvoice
	case Cash:
		return sequence_type.SpecialInvoice
	case Rectification:
		return sequence_type.RectificationInvoice
	}
	return sequence_type.Invalid
}

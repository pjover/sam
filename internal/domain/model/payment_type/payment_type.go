package payment_type

import "strings"

type PaymentType uint

const (
	Invalid PaymentType = iota
	BankDirectDebit
	BankTransfer
	Voucher
	Cash
	Rectification
)

var values = []string{
	"Indefinit",
	"Rebut",
	"Tranferència",
	"Xec escoleta",
	"Efectiu",
	"Rectificació",
}

func (s PaymentType) String() string {
	return values[s]
}

func New(value string) PaymentType {
	var _values = []string{
		"",
		"BANK_DIRECT_DEBIT",
		"BANK_TRANSFER",
		"VOUCHER",
		"CASH",
		"RECTIFICATION",
	}
	value = strings.ToLower(value)
	for i, val := range _values {
		if strings.ToLower(val) == value {
			return PaymentType(i)
		}
	}
	return Invalid
}

package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

type InvoiceHolder struct {
	Name        string
	TaxID       TaxId
	Address     Address
	Email       string
	SendEmail   bool
	PaymentType payment_type.PaymentType
	Iban        IBAN
	IsBusiness  bool
}

func (i InvoiceHolder) PaymentInfoFmt() string {
	switch i.PaymentType {
	case payment_type.BankDirectDebit:
		return fmt.Sprintf("Rebut %s", i.Iban.Format())
	case payment_type.BankTransfer:
		return fmt.Sprintf("Trans. %s", i.Iban.Format())
	case payment_type.Cash:
		return "Efectiu"
	case payment_type.Voucher:
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func (i InvoiceHolder) Mail() string {
	return fmt.Sprintf("%s <%s>", i.Name, i.Email)
}

package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

type InvoiceHolder struct {
	Name        string
	TaxID       string
	Address     Address
	Email       string
	SendEmail   bool
	PaymentType payment_type.PaymentType
	BankAccount string
	IsBusiness  bool
}

func (i InvoiceHolder) PaymentInfoFmt() string {
	switch i.PaymentType {
	case payment_type.BankDirectDebit:
		return fmt.Sprintf("Rebut %s", i.BankAccountFmt())
	case payment_type.BankTransfer:
		return fmt.Sprintf("Trans. %s", i.BankAccountFmt())
	case payment_type.Cash:
		return "Efectiu"
	case payment_type.Voucher:
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func (i InvoiceHolder) BankAccountFmt() string {
	if len(i.BankAccount) != 24 {
		return i.BankAccount
	}
	return fmt.Sprintf(
		"%s %s %s %s %s %s",
		i.BankAccount[0:4],
		i.BankAccount[4:8],
		i.BankAccount[8:12],
		i.BankAccount[12:16],
		i.BankAccount[16:20],
		i.BankAccount[20:24],
	)
}

func (i InvoiceHolder) Mail() string {
	return fmt.Sprintf("%s <%s>", i.Name, i.Email)
}

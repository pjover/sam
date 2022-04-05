package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

type InvoiceHolder struct {
	name        string
	taxID       TaxId
	address     Address
	email       string
	sendEmail   bool
	paymentType payment_type.PaymentType
	iban        IBAN
	isBusiness  bool
}

func NewInvoiceHolder(
	name string,
	taxID TaxId,
	address Address,
	email string,
	sendEmail bool,
	paymentType payment_type.PaymentType,
	iban IBAN,
	isBusiness bool,
) InvoiceHolder {
	return InvoiceHolder{
		name:        name,
		taxID:       taxID,
		address:     address,
		email:       email,
		sendEmail:   sendEmail,
		paymentType: paymentType,
		iban:        iban,
		isBusiness:  isBusiness,
	}
}

func (i InvoiceHolder) Name() string {
	return i.name
}

func (i InvoiceHolder) TaxID() TaxId {
	return i.taxID
}

func (i InvoiceHolder) Address() Address {
	return i.address
}

func (i InvoiceHolder) Email() string {
	return i.email
}

func (i InvoiceHolder) SendEmail() bool {
	return i.sendEmail
}

func (i InvoiceHolder) PaymentType() payment_type.PaymentType {
	return i.paymentType
}

func (i InvoiceHolder) Iban() IBAN {
	return i.iban
}

func (i InvoiceHolder) IsBusiness() bool {
	return i.isBusiness
}

func (i InvoiceHolder) PaymentInfoFmt() string {
	switch i.paymentType {
	case payment_type.BankDirectDebit:
		return fmt.Sprintf("Rebut %s", i.iban.Format())
	case payment_type.BankTransfer:
		return fmt.Sprintf("Trans. %s", i.iban.Format())
	case payment_type.Cash:
		return "Efectiu"
	case payment_type.Voucher:
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func (i InvoiceHolder) Mail() string {
	return fmt.Sprintf("%s <%s>", i.name, i.email)
}

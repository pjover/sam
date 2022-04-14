package model

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"net/mail"
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

func (i InvoiceHolder) Validate() error {
	if i.name == "" {
		return errors.New("el nom del titular (InvoiceHolder.Name) no pot estar buit")
	}

	if i.taxID == emptyTaxId {
		return errors.New("el DNI/NIE/NIF del titular (InvoiceHolder.TaxId) no pot estar buit")
	}

	err := i.address.validate()
	if err != nil {
		return err
	}

	_, err = mail.ParseAddress(i.email)
	if err != nil {
		return fmt.Errorf("el email del titular (InvoiceHolder.Email) no és vàlid")
	}

	if i.paymentType == payment_type.Invalid {
		return errors.New("el tipus de pagament del titular (InvoiceHolder.Address) és incorrecte, ha d'esser BANK_DIRECT_DEBIT, BANK_TRANSFER, VOUCHER o CASH")
	}

	if i.iban.IsEmpty() {
		return errors.New("el IBAN (InvoiceHolder.Iban) ha d'esser vàlid qual el tipus de pagament del titular és BANK_DIRECT_DEBIT")
	}

	return nil
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

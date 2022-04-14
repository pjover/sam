package model

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceHolder_Validate(t *testing.T) {
	type fields struct {
		name        string
		taxID       TaxId
		address     Address
		email       string
		sendEmail   bool
		paymentType payment_type.PaymentType
		iban        IBAN
		isBusiness  bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			"Valid",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			nil,
		},
		{
			"Name",
			fields{
				"",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			errors.New("el nom del titular (InvoiceHolder.Name) no pot estar buit"),
		},
		{
			"Tax ID",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882J"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			errors.New("el DNI/NIE/NIF del titular (InvoiceHolder.TaxId) no pot estar buit"),
		},
		{
			"Address",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"0707",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			errors.New("el codi postal (ZipCode) ha de tenir 5 números"),
		},
		{
			"Email",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara_sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			fmt.Errorf("el email del titular (InvoiceHolder.Email) no és vàlid"),
		},
		{
			"Payment type",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.Invalid,
				NewIbanOrEmpty("ES2830668859978258529057"),
				false,
			},
			errors.New("el tipus de pagament del titular (InvoiceHolder.Address) és incorrecte," +
				" ha d'esser BANK_DIRECT_DEBIT, BANK_TRANSFER, VOUCHER o CASH"),
		},
		{
			"IBAN can't be empty with BANK_DIRECT_DEBIT payment type",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty(""),
				false,
			},
			errors.New("el IBAN (InvoiceHolder.Iban) ha d'esser vàlid qual el tipus de " +
				"pagament del titular és BANK_DIRECT_DEBIT"),
		},
		{
			"IBAN have to be valid with BANK_DIRECT_DEBIT payment type",
			fields{
				"Cara Santamaria Novella",
				NewTaxIdOrEmpty("36361882D"),
				NewAddress(
					"Carrer Ucraïna 2022, 1st",
					"07007",
					"Palma",
					"Illes Balears",
				),
				"cara@sgu.org",
				false,
				payment_type.BankDirectDebit,
				NewIbanOrEmpty("ES3830668859978258529057"),
				false,
			},
			errors.New("el IBAN (InvoiceHolder.Iban) ha d'esser vàlid qual el tipus de " +
				"pagament del titular és BANK_DIRECT_DEBIT"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := InvoiceHolder{
				name:        tt.fields.name,
				taxID:       tt.fields.taxID,
				address:     tt.fields.address,
				email:       tt.fields.email,
				sendEmail:   tt.fields.sendEmail,
				paymentType: tt.fields.paymentType,
				iban:        tt.fields.iban,
				isBusiness:  tt.fields.isBusiness,
			}
			got := sut.Validate()
			assert.Equal(t, tt.wantErr, got)
		})
	}
}

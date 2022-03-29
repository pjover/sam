package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerToModel(t *testing.T) {
	tests := []struct {
		name     string
		customer Customer
		want     model.Customer
	}{
		{
			name: "customer DTO to model",
			customer: Customer{
				Id:     148,
				Active: true,
				Children: []Child{
					{
						Id:            1850,
						Name:          "Laura",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						BirthDate:     test_data.Date,
						Group:         "EI_1",
						Active:        true,
					},
					{
						Id:            1851,
						Name:          "Aina",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						TaxID:         "60235657Z",
						BirthDate:     test_data.Date,
						Group:         "EI_1",
						Active:        true,
					},
				},
				Adults: []Adult{
					{
						Name:          "Nom de la mare",
						Surname:       "1er llinatge_mare",
						SecondSurname: "2on llinatge_mare",
						TaxID:         "36361882D",
						Role:          "MOTHER",
					},
					{
						Name:          "Nom de la pare",
						Surname:       "1er llinatge_pare",
						SecondSurname: "2on llinatge_pare",
						TaxID:         "71032204Q",
						Role:          "FATHER",
					},
				},
				InvoiceHolder: InvoiceHolder{
					Name:  "Nom de la mare 1er llinatge_mare 2on llinatge_mare",
					TaxID: "36361882D",
					Address: Address{
						Street:  "Address first line",
						ZipCode: "07007",
						City:    "Palma",
						State:   "Illes Balears",
					},
					Email:       "email@gmail.com",
					PaymentType: "BANK_DIRECT_DEBIT",
					BankAccount: "ES2830668859978258529057",
				},
				Note:     "Nota del client",
				Language: "CA",
			},
			want: test_data.Customer148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerToModel(tt.customer)
			assert.Equal(t, tt.want, got)
		})
	}
}

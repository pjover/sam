package dto

import (
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransientCustomerToModel(t *testing.T) {
	tests := []struct {
		name     string
		customer TransientCustomer
		want     model.TransientCustomer
	}{
		{
			name: "customer DTO to model",
			customer: TransientCustomer{
				Children: []TransientChild{
					{
						Name:          "Laura",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						BirthDate:     "2019-05-25",
						Group:         "EI_1",
					},
					{
						Name:          "Aina",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						TaxID:         "60235657Z",
						BirthDate:     "2019-05-25",
						Group:         "EI_1",
					},
				},
				Adults: []Adult{
					{
						Name:          "Cara",
						Surname:       "Santamaria",
						SecondSurname: "Novella",
						TaxID:         "36361882D",
						Role:          "MOTHER",
						Address: Address{
							Street:  "Carrer Ucraïna 2022, 1st",
							ZipCode: "07007",
							City:    "Palma",
							State:   "Illes Balears",
						},
						Email:            "cara@sgu.org",
						MobilePhone:      "654321098",
						HomePhone:        "987654321",
						GrandMotherPhone: "685698789",
						GrandParentPhone: "658785478",
						WorkPhone:        "987525444",
						BirthDate:        test_data.TestDate.Format(domain.YearMonthDayLayout),
						Nationality:      "US",
					},
					{
						Name:          "Bob",
						Surname:       "Novella",
						SecondSurname: "Sagan",
						TaxID:         "71032204Q",
						Role:          "FATHER",
						Address: Address{
							Street:  "Carrer Ucraïna 2022, 1st",
							ZipCode: "07007",
							City:    "Palma",
							State:   "Illes Balears",
						},
						Email:            "bob@sgu.org",
						MobilePhone:      "654321097",
						HomePhone:        "987654322",
						GrandMotherPhone: "685698788",
						GrandParentPhone: "658785477",
						WorkPhone:        "987525446",
						BirthDate:        test_data.TestDate.Format(domain.YearMonthDayLayout),
						Nationality:      "UK",
					},
				},
				InvoiceHolder: InvoiceHolder{
					Name:  "Cara Santamaria Novella",
					TaxID: "36361882D",
					Address: Address{
						Street:  "Carrer Ucraïna 2022, 1st",
						ZipCode: "07007",
						City:    "Palma",
						State:   "Illes Balears",
					},
					Email:       "cara@sgu.org",
					PaymentType: "BANK_DIRECT_DEBIT",
					Iban:        "ES2830668859978258529057",
				},
				Note:     "Nota del client 148",
				Language: "CA",
			},
			want: model.TransientCustomer{
				Children: []model.TransientChild{
					{
						Name:          "Laura",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						TaxID:         model.NewTaxIdOrEmpty(""),
						BirthDate:     test_data.TestDate,
						Group:         group_type.EI_1,
						Note:          "",
					},
					{
						Name:          "Aina",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						TaxID:         model.NewTaxIdOrEmpty("60235657Z"),
						BirthDate:     test_data.TestDate,
						Group:         group_type.EI_1,
						Note:          "",
					},
				},
				Adults: []model.Adult{
					test_data.AdultMother148,
					test_data.AdultFather148,
				},
				InvoiceHolder: test_data.InvoiceHolder148,
				Note:          "Nota del client 148",
				Language:      language.Catalan,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TransientCustomerToModel(tt.customer)
			assert.Equal(t, tt.want, got)
		})
	}
}

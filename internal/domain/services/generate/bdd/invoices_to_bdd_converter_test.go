package bdd

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports/mocks"
	"github.com/pjover/sam/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_invoicesToBddConverter_calculateControlCode(t *testing.T) {
	tests := []struct {
		name   string
		params []string
		want   string
	}{
		{
			name:   "case 1",
			params: []string{"3066 8859 9782 5852 9057ES"},
			want:   "28",
		},
		{
			name:   "case 2",
			params: []string{"3001 2859 0880 2660 6142ES"},
			want:   "02",
		},
		{
			name:   "case 3",
			params: []string{"31795040719243310258ES"},
			want:   "87",
		},
		{
			name:   "case 4",
			params: []string{"3118-2176-0723-9984-7410ES"},
			want:   "60",
		},
		{
			name:   "case 5",
			params: []string{"HOBB", "20180707204308000"},
			want:   "24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := invoicesToBddConverter{}
			got := sut.calculateControlCode(tt.params...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_invoicesToBddConverter_getSepaIndentifier(t *testing.T) {
	type args struct {
		taxID   string
		country string
		suffix  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				taxID:   "36361882D",
				country: "ES",
				suffix:  "000",
			},
			want: "ES4200036361882D",
		},
		{
			name: "Case 2",
			args: args{
				taxID:   "37866397W",
				country: "ES",
				suffix:  "000",
			},
			want: "ES5500037866397W",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := invoicesToBddConverter{}
			got := sut.getSepaIndentifier(tt.args.taxID, tt.args.country, tt.args.suffix)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_invoicesToBddConverter_Run(t *testing.T) {
	type args struct {
		invoices  []model.Invoice
		customers map[int]model.Customer
		products  map[string]model.Product
	}
	tests := []struct {
		name string
		args args
		want Bdd
	}{
		{
			name: "returns the complete Bdd object",
			args: args{
				invoices: []model.Invoice{
					test_data.InvoiceF100,
					test_data.InvoiceF101,
					test_data.InvoiceF102,
					test_data.InvoiceF103,
				},
				customers: map[int]model.Customer{
					148: test_data.Customer148,
					149: test_data.Customer149,
				},
				products: map[string]model.Product{
					"TST": test_data.ProductTST,
					"XXX": test_data.ProductXXX,
					"YYY": test_data.ProductYYY,
				},
			},
			want: Bdd{
				messageIdentification:   "HOBB-20180707204308000-24",
				creationDateTime:        "2018-07-07T20:43:08",
				numberOfTransactions:    4,
				controlSum:              "146.60",
				name:                    "Centre d'Educació Infantil Hobbiton, S.L.",
				identification:          "ES92000B57398000",
				requestedCollectionDate: "2018-07-07",
				country:                 "ES",
				addressLine1:            "Carrer de Bisbe Rafael Josep Verger, 4",
				addressLine2:            "07010 Palma, Illes Balears",
				iban:                    "ES8004872157762000009714",
				bic:                     "GBMNESMMXXX",
				details: []BddDetail{
					{
						endToEndIdentifier:    "HOBB-20180707204308000-24.F-100",
						instructedAmount:      "36.65",
						dateOfSignature:       "2018-07-07",
						name:                  "Nom de la mare 1er llinatge_mare 2on llinatge_mare",
						identification:        "ES4200036361882D",
						iban:                  "ES2830668859978258529057",
						purposeCode:           "OTHR",
						remittanceInformation: "1xTstProduct, 3xXxxProduct, 1.5xYyyProduct",
						isBusiness:            false,
					},
					{
						endToEndIdentifier:    "HOBB-20180707204308000-24.F-101",
						instructedAmount:      "36.65",
						dateOfSignature:       "2018-07-07",
						name:                  "Nom de la mare 1er llinatge_mare 2on llinatge_mare",
						identification:        "ES4200036361882D",
						iban:                  "ES2830668859978258529057",
						purposeCode:           "OTHR",
						remittanceInformation: "1xTstProduct, 3xXxxProduct, 1.5xYyyProduct",
						isBusiness:            false,
					},
					{
						endToEndIdentifier:    "HOBB-20180707204308000-24.F-102",
						instructedAmount:      "36.65",
						dateOfSignature:       "2018-07-07",
						name:                  "Nom empresa",
						identification:        "ES5500037866397W",
						iban:                  "ES2830668859978258529057",
						purposeCode:           "OTHR",
						remittanceInformation: "1xTstProduct, 3xXxxProduct, 1.5xYyyProduct",
						isBusiness:            true,
					},
					{
						endToEndIdentifier:    "HOBB-20180707204308000-24.F-103",
						instructedAmount:      "36.65",
						dateOfSignature:       "2018-07-07",
						name:                  "Nom empresa",
						identification:        "ES5500037866397W",
						iban:                  "ES2830668859978258529057",
						purposeCode:           "OTHR",
						remittanceInformation: "1xTstProduct, 3xXxxProduct, 1.5xYyyProduct",
						isBusiness:            true,
					},
				},
			},
		},
	}

	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("GetString", "business.name").Return("Centre d'Educació Infantil Hobbiton, S.L.")
	mockedConfigService.On("GetString", "bdd.id").Return("ES92000B57398000")
	mockedConfigService.On("GetString", "bdd.country").Return("ES")
	mockedConfigService.On("GetString", "business.addressLine1").Return("Carrer de Bisbe Rafael Josep Verger, 4")
	mockedConfigService.On("GetString", "business.addressLine2").Return("07010 Palma, Illes Balears")
	mockedConfigService.On("GetString", "bdd.iban").Return("ES8004872157762000009714")
	mockedConfigService.On("GetString", "bdd.bankBic").Return("GBMNESMMXXX")
	mockedConfigService.On("GetString", "bdd.prefix").Return("HOBB")
	mockedConfigService.On("GetString", "bdd.purposeCode").Return("OTHR")

	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(time.Date(2018, time.July, 7, 20, 43, 8, 0, time.UTC))

	sut := NewInvoicesToBddConverter(mockedConfigService, mockedOsService)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sut.Run(tt.args.invoices, tt.args.customers, tt.args.products)
			assert.Equal(t, tt.want, got)
		})
	}
}

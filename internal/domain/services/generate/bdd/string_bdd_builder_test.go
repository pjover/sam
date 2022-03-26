package bdd

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func readTextFileIntoString(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func Test_stringBddBuilder_Build(t *testing.T) {
	tests := []struct {
		name string
		bdd  Bdd
		want string
	}{
		{
			name: "given a BDD should build the XML",
			bdd: Bdd{
				messageIdentification:   "HOBB-20180707204338029-50",
				creationDateTime:        "2018-07-07T20:43:38",
				numberOfTransactions:    4,
				controlSum:              "146.60",
				name:                    "Centre d'Educació Infantil Hobbiton, S.L.",
				identification:          "ES92000B57398000",
				requestedCollectionDate: "2018-07-08",
				country:                 "ES",
				addressLine1:            "Carrer de Bisbe Rafael Josep Verger, 4",
				addressLine2:            "07010 Palma, Illes Balears",
				iban:                    "ES8004872157762000009714",
				bic:                     "GBMNESMMXXX",
				details: []BddDetail{
					{
						endToEndIdentifier:    "HOBB-20180707204338029-50.F-100",
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
						endToEndIdentifier:    "HOBB-20180707204338029-50.F-101",
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
						endToEndIdentifier:    "HOBB-20180707204338029-50.F-102",
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
						endToEndIdentifier:    "HOBB-20180707204338029-50.F-103",
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
			want: readTextFileIntoString("string_bdd_builder.q1x"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringBddBuilder{}
			gotContent := s.Build(tt.bdd)
			assert.Equal(t, tt.want, gotContent)
		})
	}
}

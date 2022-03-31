package model

import (
	"fmt"
	"github.com/biter777/countries"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type IBAN struct {
	// countryCode using ISO 3166-1 alpha-2 – two letters
	countryCode countries.CountryCode
	//checkDigits for validation, two digits
	checkDigits string
	//BBAN Basic Bank Account Number, up to 30 alphanumeric characters that are country-specific
	bban string
}

//func NewBankAccount(code string) (IBAN, error) {
//	var preparedCode string
//	if code != "" {
//		preparedCode = strings.ReplaceAll(code, " ", "")
//		preparedCode = strings.ReplaceAll(preparedCode, "-", "")
//	}
//	return BankAccount(preparedCode)
//}

//func prepareCode(code string) string {
//
//}

func extractCountryCode(code string) (countries.CountryCode, error) {
	cc := code[0:2]
	countryCode := countries.ByName(cc)
	if !countryCode.IsValid() {
		return countries.Unknown, fmt.Errorf("'%s' is an invalid ISO 3166-1 alpha-2 country", cc)
	}
	return countryCode, nil
}

func extractCheckDigits(code string) (string, error) {
	checkDigits := code[2:4]

	_, err := strconv.Atoi(checkDigits)
	if err != nil {
		return "", fmt.Errorf("'%s' is an invalid two numbers IBAN check digits", checkDigits)
	}
	return checkDigits, nil
}

func isNumber(text string) bool {
	reader := strings.NewReader(text)
	text = ""

	var r rune
	var err error
	for {
		r, _, err = reader.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func extractBban(code string) (string, error) {
	bban := code[4:]
	if isNumber(bban) {
		return bban, nil
	}
	return "", fmt.Errorf("'%s' is an invalid IBAN Basic Bank Account Number", bban)
}

//func isValidCode(text string) bool {
//	reader := strings.NewReader(text)
//	text = ""
//
//	var r rune
//	var err error
//	for {
//		r, _, err = reader.ReadRune()
//		if err == io.EOF {
//			break
//		}
//		if unicode.IsLetter(r) {
//			return false
//		}
//	}
//	return true
//}

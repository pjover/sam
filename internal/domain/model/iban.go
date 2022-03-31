package model

import (
	"fmt"
	"github.com/biter777/countries"
	"strconv"
	"strings"
)

type IBAN struct {
	// countryCode using ISO 3166-1 alpha-2 – two letters
	countryCode countries.CountryCode
	//checkDigits for validation, two digits
	checkDigits string
	//BBAN Basic Bank Account Number, up to 30 alphanumeric characters that are country-specific
	bban string
}

func NewBankAccount(code string) (IBAN, error) {
	preparedCode := prepareCode(code)

	countryCode, err := extractCountryCode(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	checkDigits, err := extractCheckDigits(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	bban, err := extractBban(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	return IBAN{
		countryCode: countryCode,
		checkDigits: checkDigits,
		bban:        bban,
	}, nil
}

func prepareCode(code string) string {
	preparedCode := strings.ToUpper(code)
	if code != "" {
		preparedCode = strings.ReplaceAll(code, " ", "")
		preparedCode = strings.ReplaceAll(preparedCode, "-", "")
	}
	return preparedCode
}

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

func extractBban(code string) (string, error) {
	bban := code[4:]
	if isValidBban(bban) {
		return bban, nil
	}
	return "", fmt.Errorf("'%s' is an invalid IBAN Basic Bank Account Number", bban)
}

func isValidBban(text string) bool {
	for _, r := range []rune(text) {
		if r >= '0' && r <= '9' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		return false
	}
	return true
}

package model

import (
	"fmt"
	"github.com/biter777/countries"
	"github.com/pjover/sam/internal/domain/services/common"
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

var emptyIban = IBAN{}

func (i IBAN) String() string {
	if i == emptyIban {
		return ""
	}
	return fmt.Sprintf("%s%s%s", i.countryCode.Alpha2(), i.checkDigits, i.bban)
}

func (i IBAN) Format() string {
	if i == emptyIban {
		return ""
	}
	str := i.String()
	if len(str) != 24 {
		return str
	}
	return fmt.Sprintf(
		"%s %s %s %s %s %s",
		str[0:4],
		str[4:8],
		str[8:12],
		str[12:16],
		str[16:20],
		str[20:24],
	)
}

func (i IBAN) IsEmpty() bool {
	return i == emptyIban
}

func NewIban(iban string) (IBAN, error) {
	if iban == "" {
		return emptyIban, nil
	}
	preparedCode := prepareIbanCode(iban)

	countryCode, err := extractIbanCountryCode(preparedCode)
	if err != nil {
		return invalidIban(iban, err)
	}

	checkDigits, err := extractIbanCheckDigits(preparedCode)
	if err != nil {
		return invalidIban(iban, err)
	}

	bban, err := extractBban(preparedCode)
	if err != nil {
		return invalidIban(iban, err)
	}

	err = validateCheckDigits(countryCode, checkDigits, bban)
	if err != nil {
		return invalidIban(iban, err)
	}

	return IBAN{
		countryCode: countryCode,
		checkDigits: checkDigits,
		bban:        bban,
	}, nil
}

func invalidIban(iban string, err error) (IBAN, error) {
	return emptyIban, fmt.Errorf("invalid IBAN '%s': %s", iban, err)
}

func validateCheckDigits(countryCode countries.CountryCode, checkDigits string, bban string) error {
	mod9710 := common.NewMod9710(bban, countryCode.Alpha2())
	checksum := mod9710.Checksum()
	if checkDigits != checksum {
		return fmt.Errorf("'%s' is an invalid two numbers IBAN check digits, does not match with '%s' checksum", checkDigits, checksum)
	}
	return nil
}

func prepareIbanCode(code string) string {
	preparedCode := strings.ToUpper(code)
	if code != "" {
		preparedCode = strings.ReplaceAll(code, " ", "")
		preparedCode = strings.ReplaceAll(preparedCode, "-", "")
	}
	return preparedCode
}

func extractIbanCountryCode(code string) (countries.CountryCode, error) {
	cc := code[0:2]
	countryCode := countries.ByName(cc)
	if !countryCode.IsValid() {
		return countries.Unknown, fmt.Errorf("'%s' is an invalid ISO 3166-1 alpha-2 country", cc)
	}
	return countryCode, nil
}

func extractIbanCheckDigits(code string) (string, error) {
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

func NewIbanOrEmpty(iban string) IBAN {
	newIban, err := NewIban(iban)
	if err != nil {
		return emptyIban
	}
	return newIban
}

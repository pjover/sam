package model

import (
	"fmt"
	"github.com/biter777/countries"
	"github.com/pjover/sam/internal/domain/services/common"
	"log"
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

func (i IBAN) String() string {
	return fmt.Sprintf("%s%s%s", i.countryCode.Alpha2(), i.checkDigits, i.bban)
}

func (i IBAN) Format() string {
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

func NewIban(code string) (IBAN, error) {
	preparedCode := prepareIbanCode(code)

	countryCode, err := extractIbanCountryCode(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	checkDigits, err := extractIbanCheckDigits(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	bban, err := extractBban(preparedCode)
	if err != nil {
		return IBAN{}, err
	}

	err = validateCheckDigits(countryCode, checkDigits, bban)
	if err != nil {
		return IBAN{}, err
	}

	return IBAN{
		countryCode: countryCode,
		checkDigits: checkDigits,
		bban:        bban,
	}, nil
}

func NewIbanOrFatal(code string) IBAN {
	iban, err := NewIban(code)
	if err != nil {
		log.Fatal(err)
	}
	return iban
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

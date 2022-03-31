package model

type IBAN struct {
	// countryCode using ISO 3166-1 alpha-2 – two letters
	countryCode string
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

func extractCountryCode(code string) (string, error) {
	cc := code[0:2]
	return cc, nil
}

//func (b BankAccount) IsValid() bool {
//	if b == "" {
//		return false
//	}
//	var controlCode
//}

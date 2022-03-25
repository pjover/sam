package bdd

type Bdd struct {
	messageIdentification   string
	creationDateTime        string
	numberOfTransactions    int
	controlSum              string
	name                    string
	identification          string
	requestedCollectionDate string
	country                 string
	addressLine1            string
	addressLine2            string
	iban                    string
	bic                     string
	details                 []BddDetail
}

type BddDetail struct {
	endToEndIdentifier    string
	instructedAmount      string
	dateOfSignature       string
	name                  string
	identification        string
	iban                  string
	purposeCode           string
	remittanceInformation string
	isBusiness            bool
}

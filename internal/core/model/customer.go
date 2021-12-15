package model

type Child struct {
	Code          int         `json:"code"`
	Name          string      `json:"name"`
	Surname       string      `json:"surname"`
	SecondSurname string      `json:"secondSurname"`
	TaxID         string      `json:"taxId"`
	BirthDate     string      `json:"birthDate"`
	Group         string      `json:"group"`
	Note          interface{} `json:"note"`
	Active        bool        `json:"active"`
	Score         float32     `json:"score"`
}

type Adult struct {
	Name             string      `json:"name"`
	Surname          string      `json:"surname"`
	SecondSurname    interface{} `json:"secondSurname"`
	TaxID            string      `json:"taxId"`
	Role             string      `json:"role"`
	Address          interface{} `json:"address"`
	Email            string      `json:"email"`
	MobilePhone      string      `json:"mobilePhone"`
	HomePhone        interface{} `json:"homePhone"`
	GrandMotherPhone interface{} `json:"grandMotherPhone"`
	GrandParentPhone interface{} `json:"grandParentPhone"`
	WorkPhone        interface{} `json:"workPhone"`
	BirthDate        interface{} `json:"birthDate"`
	Nationality      interface{} `json:"nationality"`
	Score            float32     `json:"score"`
}

type Address struct {
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type InvoiceHolder struct {
	Name        string  `json:"name"`
	TaxID       string  `json:"taxId"`
	Address     Address `json:"address"`
	Email       string  `json:"email"`
	SendEmail   bool    `json:"sendEmail"`
	PaymentType string  `json:"paymentType"`
	BankAccount string  `json:"bankAccount"`
	IsBusiness  bool    `json:"isBusiness"`
}

type Customer struct {
	//TODO Id            int           `json:"_id"`
	Active        bool          `json:"active"`
	Children      []Child       `json:"children"`
	Adults        []Adult       `json:"adults"`
	InvoiceHolder InvoiceHolder `json:"invoiceHolder"`
	Note          interface{}   `json:"note"`
	Language      string        `json:"language"`
}

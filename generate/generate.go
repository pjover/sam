package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sam/adm"
	"sam/util"
)

type GenerateManager interface {
	GenerateBdd() (string, error)
	GenerateInvoice(invoiceCode string) (string, error)
	GenerateInvoices(onlyNew bool) (string, error)
	GenerateCustomerReport() (string, error)
}

type GenerateManagerImpl struct {
	getManager  util.HttpGetManager
	postManager util.HttpPostManager
}

func NewGenerateManager() GenerateManager {
	return GenerateManagerImpl{
		util.NewHttpGetManager(),
		util.NewHttpPostManager(),
	}
}

func (g GenerateManagerImpl) GenerateBdd() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")
	url := fmt.Sprintf("%s/generate/bdd?yearMonth=%s",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
	)

	dir := getDirectory()
	currentFilenames := listFiles(dir, ".qx1")
	filename := getNextBddFilename(currentFilenames)
	return g.postManager.File(url, dir, filename)
}

func listFiles(dir string, ext string) []string {
	var filenames []string
	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == ext {
			filenames = append(filenames, info.Name())
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return filenames
}

func getNextBddFilename(currentFilenames []string) string {
	sequence := len(currentFilenames) + 1
	filename := buildBddFilename(sequence)
	for util.StringInList(filename, currentFilenames) {
		sequence += 1
		filename = buildBddFilename(sequence)
	}
	return filename
}

func buildBddFilename(sequence int) string {
	return fmt.Sprintf("bdd-%d.qx1", sequence)
}

func (g GenerateManagerImpl) GenerateInvoice(invoiceCode string) (string, error) {
	fmt.Println("Generant la factura", invoiceCode)

	url := fmt.Sprintf("%s/generate/pdf/%s", viper.GetString("urls.hobbit"), invoiceCode)

	return g.postManager.FileDefaultName(url, getDirectory())
}

func getDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}

func (g GenerateManagerImpl) GenerateInvoices(onlyNew bool) (string, error) {
	fmt.Println("Generant les factures del mes")

	url := fmt.Sprintf(
		"%s/generate/pdf?yearMonth=%s&notYetPrinted=%t",
		viper.GetString("urls.hobbit"),
		viper.GetString("yearMonth"),
		onlyNew,
	)

	dirPath := path.Join(getDirectory(), viper.GetString("dirs.invoices"))
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	return g.postManager.Zip(url, dirPath)
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

type InvoiceHolder struct {
	Name    string `json:"name"`
	TaxID   string `json:"taxId"`
	Address struct {
		Street  string `json:"street"`
		ZipCode string `json:"zipCode"`
		City    string `json:"city"`
		State   string `json:"state"`
	} `json:"address"`
	Email       string `json:"email"`
	SendEmail   bool   `json:"sendEmail"`
	PaymentType string `json:"paymentType"`
	BankAccount string `json:"bankAccount"`
	IsBusiness  bool   `json:"isBusiness"`
}

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

type ActiveCustomers struct {
	Embedded struct {
		Customers []struct {
			Children      []Child       `json:"children"`
			Adults        []Adult       `json:"adults"`
			InvoiceHolder InvoiceHolder `json:"invoiceHolder"`
			Note          interface{}   `json:"note"`
			Language      string        `json:"language"`
			Active        bool          `json:"active"`
			Links         struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Customer struct {
					Href string `json:"href"`
				} `json:"customer"`
			} `json:"_links"`
		} `json:"customers"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (g GenerateManagerImpl) GenerateCustomerReport() (string, error) {
	fmt.Println("Generant l'informe de clients")

	url := fmt.Sprintf(
		"%s/customers/search/findAllByActiveTrue",
		viper.GetString("urls.hobbit"),
	)
	customers := new(ActiveCustomers)
	err := g.getManager.Type(url, customers)
	if err != nil {
		return "", err
	}

	var contents [][]string
	for _, c := range customers.Embedded.Customers {
		adult := getFirstAdult(c.Adults)
		for _, child := range c.Children {
			if !c.Active {
				continue
			}
			var line = []string{
				formatChildName(child),
				child.Group,
				child.BirthDate,
				formatAdultName(adult),
				formatPhone(adult.MobilePhone),
				adult.Email,
				formatPaymentInfo(c.InvoiceHolder),
			}
			contents = append(contents, line)
		}

	}

	filePath := path.Join(getDirectory(), viper.GetString("files.customerReport"))
	reportInfo := adm.ReportInfo{
		filePath,
		consts.Landscape,
		"Llistat de clients",
		[]string{
			"Infant",
			"Grup",
			"Neixament",
			"Mare",
			"Mòbil",
			"Correu",
			"Pagament",
		},
		contents,
		[]uint{2, 1, 1, 2, 1, 2, 3},
		consts.Left,
	}
	err = adm.CustomerReportPdf(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de clients a '%s'", filePath), nil
}

func getFirstAdult(adults []Adult) Adult {

	for _, adult := range adults {
		if adult.Role == "MOTHER" {
			return adult
		}
	}
	return adults[0]
}

func formatChildName(child Child) string {
	return fmt.Sprintf("%d   %s %s", child.Code, child.Name, child.Surname)
}

func formatAdultName(adult Adult) string {
	return fmt.Sprintf("%s %s", adult.Name, adult.Surname)
}

func formatPaymentInfo(invoiceHolder InvoiceHolder) string {
	switch invoiceHolder.PaymentType {
	case "BANK_DIRECT_DEBIT":
		return fmt.Sprintf("Rebut %s", formatIban(invoiceHolder.BankAccount))
	case "BANK_TRANSFER":
		return fmt.Sprintf("Trans. %s", formatIban(invoiceHolder.BankAccount))
	case "CASH":
		return "Efectiu"
	case "VOUCHER":
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func formatPhone(phone string) string {
	if len(phone) != 9 {
		return phone
	}
	return fmt.Sprintf(
		"%s %s %s",
		phone[0:3],
		phone[3:6],
		phone[6:9],
	)
}

func formatIban(iban string) string {
	if len(iban) != 24 {
		return iban
	}
	return fmt.Sprintf(
		"%s %s %s %s %s %s",
		iban[0:4],
		iban[4:8],
		iban[8:12],
		iban[12:16],
		iban[16:20],
		iban[20:24],
	)
}

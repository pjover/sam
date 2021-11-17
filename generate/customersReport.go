package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"path"
	"sam/model"
	"sam/util"
	"sort"
)

type CustomersReportGenerator struct {
	getManager util.HttpGetManager
}

func NewCustomersReportGenerator(getManager util.HttpGetManager) CustomersReportGenerator {
	return CustomersReportGenerator{
		getManager,
	}
}

func (c CustomersReportGenerator) generate() (string, error) {

	customers, err := c.getCustomers(c.getManager)
	if err != nil {
		return "", err
	}

	contents := c.buildContents(customers)

	filePath := path.Join(
		viper.GetString("dirs.reports"),
		viper.GetString("files.customersReport"),
	)
	reportInfo := util.ReportInfo{
		consts.Landscape,
		consts.Left,
		"Llistat de clients",
		[]util.Column{
			{"Infant", 2},
			{"Grup", 1},
			{"Neixament", 1},
			{"Mare", 2},
			{"Mòbil", 1},
			{"Correu", 2},
			{"Pagament", 3},
		},
		contents,
		filePath,
	}
	err = util.PdfReport(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de clients a '%s'", filePath), nil
}

type activeCustomers struct {
	Embedded struct {
		Customers []model.Customer `json:"customers"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func (c CustomersReportGenerator) getCustomers(getManager util.HttpGetManager) (*activeCustomers, error) {
	url := fmt.Sprintf(
		"%s/customers/search/findAllByActiveTrue",
		viper.GetString("urls.hobbit"),
	)
	customers := new(activeCustomers)
	err := getManager.Type(url, customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c CustomersReportGenerator) buildContents(customers *activeCustomers) [][]string {
	var contents [][]string
	for _, customer := range customers.Embedded.Customers {
		adult := customer.FirstAdult()
		for _, child := range customer.Children {
			if !child.Active {
				continue
			}
			var line = []string{
				child.NameWithCode(),
				child.Group,
				child.BirthDate,
				c.formatAdultName(adult),
				c.formatPhone(adult.MobilePhone),
				adult.Email,
				c.formatPaymentInfo(customer.InvoiceHolder),
			}
			contents = append(contents, line)
		}
	}
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents
}

func (c CustomersReportGenerator) formatAdultName(adult model.Adult) string {
	return fmt.Sprintf("%s %s", adult.Name, adult.Surname)
}

func (c CustomersReportGenerator) formatPhone(phone string) string {
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

func (c CustomersReportGenerator) formatPaymentInfo(invoiceHolder model.InvoiceHolder) string {
	switch invoiceHolder.PaymentType {
	case "BANK_DIRECT_DEBIT":
		return fmt.Sprintf("Rebut %s", c.formatIban(invoiceHolder.BankAccount))
	case "BANK_TRANSFER":
		return fmt.Sprintf("Trans. %s", c.formatIban(invoiceHolder.BankAccount))
	case "CASH":
		return "Efectiu"
	case "VOUCHER":
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}

func (c CustomersReportGenerator) formatIban(iban string) string {
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

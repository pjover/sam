package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"path"
	"sam/adm"
	"sam/model"
	"sam/util"
	"sort"
)

func generateCustomersReport(getManager util.HttpGetManager) (string, error) {

	customers, err := getCustomers(getManager)
	if err != nil {
		return "", err
	}

	contents := buildContents(customers)

	filePath := path.Join(getDirectory(), viper.GetString("files.customerReport"))
	reportInfo := adm.ReportInfo{
		consts.Landscape,
		consts.Left,
		"Llistat de clients",
		[]adm.Column{
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
	err = adm.PdfReport(reportInfo)
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

func getCustomers(getManager util.HttpGetManager) (*activeCustomers, error) {
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

func buildContents(customers *activeCustomers) [][]string {
	var contents [][]string
	for _, c := range customers.Embedded.Customers {
		adult := getFirstAdult(c.Adults)
		for _, child := range c.Children {
			if !child.Active {
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
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents
}

func getFirstAdult(adults []model.Adult) model.Adult {

	for _, adult := range adults {
		if adult.Role == "MOTHER" {
			return adult
		}
	}
	return adults[0]
}

func formatChildName(child model.Child) string {
	return fmt.Sprintf("%d   %s %s", child.Code, child.Name, child.Surname)
}

func formatAdultName(adult model.Adult) string {
	return fmt.Sprintf("%s %s", adult.Name, adult.Surname)
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

func formatPaymentInfo(invoiceHolder model.InvoiceHolder) string {
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

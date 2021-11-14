package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"path"
	"sam/adm"
	"sam/model"
	"sam/util"
)

type ActiveCustomers struct {
	Embedded struct {
		Customers []model.Customer `json:"customers"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

func generateCustomersReport(getManager util.HttpGetManager) (string, error) {
	fmt.Println("Generant l'informe de clients")

	url := fmt.Sprintf(
		"%s/customers/search/findAllByActiveTrue",
		viper.GetString("urls.hobbit"),
	)
	customers := new(ActiveCustomers)
	err := getManager.Type(url, customers)
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

package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/services/generate/reports"
	"path"
	"sort"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/generate"
	"github.com/spf13/viper"
)

type CustomersReportGenerator struct {
	getManager tuk.HttpGetManager
}

func NewCustomersReportGenerator() generate.Generator {
	return CustomersReportGenerator{
		tuk.NewHttpGetManager(),
	}
}

func (c CustomersReportGenerator) Generate() (string, error) {
	fmt.Println("Generant l'informe de clients ...")

	customers, err := c.getCustomers(c.getManager)
	if err != nil {
		return "", err
	}

	contents := c.buildContents(customers)

	filePath := path.Join(
		viper.GetString("dirs.reports"),
		viper.GetString("files.customersReport"),
	)
	reportInfo := reports.ReportInfo{
		consts.Landscape,
		consts.Left,
		"Llistat de clients",
		[]reports.Column{
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
	err = reports.Report(reportInfo)
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

func (c CustomersReportGenerator) getCustomers(getManager tuk.HttpGetManager) (*activeCustomers, error) {
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
				child.BirthDate.String(),
				adult.NameSurnameFmt(),
				adult.MobilePhoneFmt(),
				adult.Email,
				customer.InvoiceHolder.PaymentInfoFmt(),
			}
			contents = append(contents, line)
		}
	}
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents
}

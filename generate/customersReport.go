package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"path"
	"sam/internal/util"
	"sam/model"
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

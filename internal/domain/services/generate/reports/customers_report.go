package reports

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
	"sort"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
)

type CustomerReport struct {
	dbService ports.DbService
}

func NewCustomerReport(dbService ports.DbService) CustomerReport {
	return CustomerReport{
		dbService: dbService,
	}
}

func (c CustomerReport) Run() (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("Generant l'informe de clients ...\n")

	customers, err := c.getCustomers()
	if err != nil {
		return "", err
	}

	contents := c.buildContents(customers)

	filePath := path.Join(
		viper.GetString("dirs.reports"),
		viper.GetString("files.customersReport"),
	)
	reportInfo := ReportInfo{
		consts.Landscape,
		consts.Left,
		"Llistat de clients",
		[]Column{
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
	err = Report(reportInfo)
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("Generat l'informe de clients a '%s'\n", filePath))
	return buffer.String(), nil
}

func (c CustomerReport) getCustomers() ([]model.Customer, error) {
	customers, err := c.dbService.FindActiveCustomers()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c CustomerReport) buildContents(customers []model.Customer) [][]string {
	var contents [][]string
	for _, customer := range customers {
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

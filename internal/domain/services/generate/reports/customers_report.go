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

	report := Report{
		PageOrientation: consts.Landscape,
		Title:           "Llistat de clients",
		SubReports: []SubReport{
			{
				Style: Table,
				Align: consts.Left,
				Captions: []string{
					"Infant",
					"Grup",
					"Neixament",
					"Mare",
					"Mòbil",
					"Correu",
					"Pagament",
				},
				Widths: []uint{
					2,
					1,
					1,
					2,
					1,
					2,
					3,
				},
				Data: c.buildData(customers),
			},
		},
	}

	filePath := path.Join(
		viper.GetString("dirs.reports"),
		viper.GetString("files.customersReport"),
	)
	err = report.SaveToFile(filePath)
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
				child.NameWithId(),
				child.Group,
				child.BirthDate.Format("2006-02-01"),
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

func (c CustomerReport) buildData(customers []model.Customer) [][]string {
	var data [][]string
	for _, customer := range customers {
		adult := customer.FirstAdult()
		for _, child := range customer.Children {
			if !child.Active {
				continue
			}
			var dataLine = []string{
				child.NameWithId(),
				child.Group,
				child.BirthDate.Format("2006-02-01"),
				adult.NameSurnameFmt(),
				adult.MobilePhoneFmt(),
				adult.Email,
				customer.InvoiceHolder.PaymentInfoFmt(),
			}
			data = append(data, dataLine)
		}
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	return data
}

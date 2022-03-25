package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
	"strconv"
	"time"
)

type CustomerCardsReports struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewCustomerCardsReports(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) CustomerCardsReports {
	return CustomerCardsReports{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (c CustomerCardsReports) Run() (string, error) {

	changedSince := c.configService.GetTime("reports.lastCustomersCardsUpdated")
	err := c.configService.SetTime("reports.lastCustomersCardsUpdated", c.osService.Now())
	if err != nil {
		return "", fmt.Errorf("no s'ha pogut actualitzar la configuració: %s", err)
	}

	customers, err := c.dbService.FindChangedCustomers(changedSince)
	if err != nil {
		_ = c.configService.SetTime("reports.lastCustomersCardsUpdated", changedSince)
		return c.revertLastCustomersCardsUpdated(changedSince, fmt.Errorf("no s'ha pogut carregar els consumidors des de la base de dades: %s", err))
	}

	reportsDir, err := c.configService.GetCustomersCardsDirectory()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant les fitxes del clients que han canviat des de %s ...\n", changedSince.Format(domain.YearMonthDayLayout)))
	for _, customer := range customers {
		msg, err := c.run(reportsDir, customer)
		if err != nil {
			return c.revertLastCustomersCardsUpdated(changedSince, err)
		}
		buffer.WriteString(msg)
	}
	buffer.WriteString(fmt.Sprintf("Generades %d fitxes de clients", len(customers)))
	return buffer.String(), nil
}

func (c CustomerCardsReports) revertLastCustomersCardsUpdated(changedSince time.Time, err error) (string, error) {
	_ = c.configService.SetTime("reports.lastCustomersCardsUpdated", changedSince)
	return "", err
}

func (c CustomerCardsReports) run(dirPath string, customer model.Customer) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Generant la fitxa del client %d, %s ...\n", customer.Id, customer.InvoiceHolder.Name))

	childrenNames := customer.ChildrenNamesWithSurname(", ")
	reportDefinition := ReportDefinition{
		PageOrientation: consts.Portrait,
		Title:           fmt.Sprintf("Fitxa del client %d: %s", customer.Id, childrenNames),
		Footer:          c.osService.Now().Format(domain.YearMonthDayLayout),
		SubReports:      c.subReports(customer),
	}

	filePath := path.Join(
		dirPath,
		fmt.Sprintf("%d-%s.pdf", customer.Id, childrenNames),
	)

	reportService := NewReportService(c.configService)
	err := reportService.SaveToFile(reportDefinition, filePath)
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("Generat la fitxa del client a '%s'\n", filePath))
	return buffer.String(), nil
}

func (c CustomerCardsReports) subReports(customer model.Customer) []SubReport {
	var subReports []SubReport
	subReports = append(subReports, c.headerSubReport(customer))
	subReports = append(subReports, c.childrenSubReports(customer.Children)...)
	subReports = append(subReports, c.adultsSubReports(customer.Adults)...)
	subReports = append(subReports, c.holderSubReport(customer.InvoiceHolder))
	return subReports
}

func (c CustomerCardsReports) headerSubReport(customer model.Customer) SubReport {
	return CardSubReport{
		Title: "",
		Align: consts.Left,
		Captions: []string{
			"Idioma",
			"Nota",
			"Actiu",
		},
		Widths: []uint{
			2,
			15,
		},
		Data: [][]string{
			{customer.Language},
			{customer.Note},
			{c.boolToYesNo(customer.Active)},
		},
	}
}

func (c CustomerCardsReports) boolToYesNo(active bool) string {
	if active {
		return "Si"
	} else {
		return "No"
	}
}

func (c CustomerCardsReports) childrenSubReports(children []model.Child) []SubReport {
	var subReports []SubReport
	for _, child := range children {
		subReports = append(subReports, c.childSubReport(child))
	}
	return subReports
}

func (c CustomerCardsReports) childSubReport(child model.Child) SubReport {
	return CardSubReport{
		Title: child.NameWithId(),
		Align: consts.Left,
		Captions: []string{
			"Codi",
			"Nom",
			"1er llinatge",
			"2on llinatge",
			"NIF/CIF",
			"Naixement",
			"Grup",
			"Nota",
			"Actiu",
		},
		Widths: []uint{
			2,
			15,
		},
		Data: [][]string{
			{strconv.Itoa(child.Id)},
			{child.Name},
			{child.Surname},
			{child.SecondSurname},
			{child.TaxID},
			{child.BirthDate.Format(domain.YearMonthDayLayout)},
			{child.Group},
			{child.Note},
			{c.boolToYesNo(child.Active)},
		},
	}
}

func (c CustomerCardsReports) adultsSubReports(adults []model.Adult) []SubReport {
	var subReports []SubReport
	for _, adult := range adults {
		subReports = append(subReports, c.adultSubReport(adult))
	}
	return subReports
}

func (c CustomerCardsReports) adultSubReport(adult model.Adult) SubReport {
	return CardSubReport{
		Title: adult.NameAndSurname(),
		Align: consts.Left,
		Captions: []string{
			"Nom",
			"1er llinatge",
			"2on llinatge",
			"NIF/CIF",
			"Naixement",
			"Nacionalitat",
			"Rol",
			"Correu",
			"Adreça",
			"Tel. mòbil",
			"Tel. casa",
			"Tel. treball",
			"Tel. padrí",
			"Tel. padrina",
		},
		Widths: []uint{
			2,
			15,
		},
		Data: [][]string{
			{adult.Name},
			{adult.Surname},
			{adult.SecondSurname},
			{adult.TaxID},
			{adult.BirthDate.Format(domain.YearMonthDayLayout)},
			{adult.Nationality},
			{adult.Role},
			{adult.Email},
			{adult.Address.CompleteAddress()},
			{adult.MobilePhone},
			{adult.HomePhone},
			{adult.WorkPhone},
			{adult.GrandParentPhone},
			{adult.GrandMotherPhone},
		},
	}
}

func (c CustomerCardsReports) holderSubReport(holder model.InvoiceHolder) SubReport {
	return CardSubReport{
		Title: "Dades de facturació",
		Align: consts.Left,
		Captions: []string{
			"Titular",
			"NIF/CIF",
			"Correu",
			"Adreça",
			"Pagament",
			"Compte corrent",
			"Enviar correu",
			"És una empresa",
		},
		Widths: []uint{
			2,
			15,
		},
		Data: [][]string{
			{holder.Name},
			{holder.TaxID},
			{holder.Email},
			{holder.Address.CompleteAddress()},
			{holder.PaymentType.String()},
			{holder.BankAccount},
			{c.boolToYesNo(holder.SendEmail)},
			{c.boolToYesNo(holder.IsBusiness)},
		},
	}
}

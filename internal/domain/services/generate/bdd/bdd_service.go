package bdd

import (
	"errors"
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/loader"
)

type BddContent struct {
	xmlText              string
	numberOfTransactions int
	controlSum           string
}

type BddService interface {
	Run() (string, error)
}

type bddService struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewBddService(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) BddService {
	return bddService{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (b bddService) Run() (string, error) {
	fmt.Println("Generant el fitxer de rebuts ...")

	bulkLoader := loader.NewBulkLoader(b.configService, b.dbService)
	customers, products, err := bulkLoader.LoadCustomersAndProducts()
	if err != nil {
		return "", err
	}

	invoices, err := b.loadInvoices()
	if err != nil {
		return "", err
	}
	if len(invoices) == 0 {
		return "No hi han rebuts pendents de generar", nil
	}

	content := b.generateContent(invoices, customers, products)
	if err != nil {
		return "", err
	}

	dirPath, filename, err := b.getFilePath()
	if err != nil {
		return "", err
	}

	filePath, err := b.saveToFile(dirPath, filename, content.xmlText)
	if err != nil {
		return "", err
	}

	err = b.updateInvoices(invoices)

	return fmt.Sprintf("S'ha generat el fitxer '%s' amb %d rebuts i import %s", filePath, content.numberOfTransactions, content.controlSum), nil
}

func (b bddService) loadInvoices() (invoices []model.Invoice, err error) {
	invoices, err = b.dbService.FindInvoicesByPaymentTypeAndSentToBank(payment_type.BankDirectDebit, false)
	if err != nil {
		return nil, errors.New("no s'han pogut recuperar les factures de rebuts pendents d'enviar al banc")
	}
	fmt.Printf("recuperades %d factures sense enviar al banc\n", len(invoices))
	return invoices, nil
}

func (b bddService) generateContent(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) BddContent {
	invoicesToBddConverter := NewInvoicesToBddConverter(b.configService, b.osService)
	bdd := invoicesToBddConverter.Convert(invoices, customers, products)

	bddBuilder := NewStringBddBuilder()

	return BddContent{
		xmlText:              bddBuilder.Build(bdd),
		numberOfTransactions: bdd.numberOfTransactions,
		controlSum:           bdd.controlSum,
	}
}

func (b bddService) getFilePath() (dirPath string, filename string, err error) {
	dirPath = b.configService.GetWorkingDirectory()
	currentFilenames, err := b.osService.ListFiles(dirPath, ".qx1")
	if err != nil {
		return "", "", err
	}

	filename = b.getNextBddFilename(currentFilenames)
	return dirPath, filename, nil
}

func (b bddService) getNextBddFilename(currentFilenames []string) string {
	sequence := len(currentFilenames) + 1
	filename := b.buildBddFilename(sequence)
	for b.stringInList(filename, currentFilenames) {
		sequence += 1
		filename = b.buildBddFilename(sequence)
	}
	return filename
}

func (b bddService) stringInList(str string, list []string) bool {
	for _, b := range list {
		if b == str {
			return true
		}
	}
	return false
}

func (b bddService) buildBddFilename(sequence int) string {
	return fmt.Sprintf("bdd-%d.qx1", sequence)
}

func (b bddService) saveToFile(dirPath string, filename string, content string) (filePath string, err error) {
	return b.osService.WriteFile(dirPath, filename, []byte(content))
}

func (b bddService) updateInvoices(invoices []model.Invoice) error {
	var updatedInvoices []model.Invoice
	for _, invoice := range invoices {
		updatedInvoices = append(updatedInvoices, invoice.SendToBank())
	}
	return b.dbService.UpdateInvoices(updatedInvoices)
}

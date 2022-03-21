package bdd

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
)

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

	bulkLoader := common.NewBulkLoader(b.configService, b.dbService)
	customers, products, err := bulkLoader.LoadCustomersAndProducts()
	if err != nil {
		return "", err
	}

	invoices, err := b.loadInvoices()
	if err != nil {
		return "", err
	}

	content, err := b.generateContent(invoices, customers, products)
	if err != nil {
		return "", err
	}

	filePath, err := b.getFilePath()
	if err != nil {
		return "", err
	}

	err = b.saveToFile(content, filePath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("S'ha generat el fitxer '%s' amb %d rebuts", filePath, len(invoices)), nil
}

func (b bddService) loadInvoices() (invoices []model.Invoice, err error) {
	yearMonth := b.configService.GetString("yearMonth")
	invoices, err = b.dbService.FindInvoicesByYearMonthAndPaymentTypeAndSentToBank(yearMonth, payment_type.BankDirectDebit, false)
	if err != nil {
		return nil, fmt.Errorf("no s'han pogut recuperar les factures de rebuts del mes %s pendents d'enviar al banc", yearMonth)
	}
	return invoices, nil
}

func (b bddService) generateContent(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product) (content string, err error) {
	return "", nil
}

func (b bddService) getFilePath() (filePath string, err error) {
	dir, err := b.configService.GetWorkingDirectory()
	if err != nil {
		return "", err
	}

	currentFilenames, err := b.osService.ListFiles(dir, ".qx1")
	if err != nil {
		return "", err
	}

	filename := b.getNextBddFilename(currentFilenames)
	return filename, nil
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

func (b bddService) saveToFile(content string, filePath string) error {
	//TODO implement me
	panic("implement me")
}

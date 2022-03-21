package bdd

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
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
	//TODO implement me
	panic("implement me")
}

func (b bddService) saveToFile(invoices []model.Invoice, customers map[int]model.Customer, products map[string]model.Product, filePath string) error {
	//TODO implement me
	panic("implement me")
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

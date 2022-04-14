package admin

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/ports/mocks"
	"github.com/pjover/sam/internal/domain/services/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func commonMocks() (*mocks.ConfigService, *mocks.OsService) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("GetConfigDirectory").Return("/home/user/.sam")
	mockedConfigService.On("GetHomeDirectory").Return("/fake/dir")
	mockedConfigService.On("GetInvoicesDirectory").Return("/fake/dir/211000-Factures del mes d'Octubre/factures")
	mockedConfigService.On("GetReportsDirectory").Return("/fake/dir/reports")
	mockedConfigService.On("GetBackupDirectory").Return("/fake/dir/reports")
	mockedConfigService.On("GetCustomersCardsDirectory").Return("/fake/dir/reports")
	mockedConfigService.On("SetCurrentYearMonth", mock.Anything).Return(nil)
	mockedConfigService.On("SetString", mock.Anything, mock.Anything).Return(nil)

	mockedOsService := new(mocks.OsService)
	mockedOsService.On("CopyFile", mock.Anything, mock.Anything).Return(nil)
	mockedOsService.On("CreateDirectory", "/fake/dir/211000-Factures del mes d'Octubre").Return(nil)
	mockedOsService.On("CreateDirectory", "/fake/dir/211000-Factures del mes d'Octubre/factures").Return(nil)
	mockedOsService.On("CreateDirectory", "/fake/dir/reports").Return(nil)
	mockedOsService.On("ItemExists", mock.Anything).Return(true, nil)
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("OpenUrlInBrowser", "/fake/dir/211000-Factures del mes d'Octubre").Return(nil)

	return mockedConfigService, mockedOsService
}

func Test_CreateDirectory_exists(t *testing.T) {
	mockedConfigService, mockedOsService := commonMocks()

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should not create the directory if exists", func(t *testing.T) {
		msg, err := sut.CreateDirectories()
		assert.Equal(t, fmt.Sprintf("%sSam v%s    /fake/dir/211000-Factures del mes d'Octubre%s", domain.ColorRed, domain.Version, domain.ColorReset), msg)
		assert.Equal(t, nil, err)
	})
}

func Test_CreateDirectory_does_not_exists(t *testing.T) {
	mockedConfigService, mockedOsService := commonMocks()

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should create the directory if does not exists", func(t *testing.T) {
		msg, err := sut.CreateDirectories()
		assert.Equal(t, fmt.Sprintf("\x1b[31mSam v%s    /fake/dir/211000-Factures del mes d'Octubre\x1b[0m", domain.Version), msg)
		assert.Equal(t, nil, err)
	})
}

func Test_Backup_ok(t *testing.T) {
	mockedConfigService, mockedOsService := commonMocks()
	mockedConfigService.On("GetString", "dirs.backup").Return("/fake/dir")
	mockedOsService.On("CreateZipFile", "/fake/dir/211031-Backup.zip", mock.Anything).Return(nil)
	mockedOsService.On("GetTempDirectory").Return("/tmp/sam", nil)
	mockedOsService.On("RunCommand", "mongoexport", "--db=hobbit", "--collection=consumption", "--out=/tmp/sam/consumption.json").Return(nil)
	mockedOsService.On("RunCommand", "mongoexport", "--db=hobbit", "--collection=customer", "--out=/tmp/sam/customer.json").Return(nil)
	mockedOsService.On("RunCommand", "mongoexport", "--db=hobbit", "--collection=invoice", "--out=/tmp/sam/invoice.json").Return(nil)
	mockedOsService.On("RunCommand", "mongoexport", "--db=hobbit", "--collection=product", "--out=/tmp/sam/product.json").Return(nil)
	mockedOsService.On("RunCommand", "mongoexport", "--db=hobbit", "--collection=sequence", "--out=/tmp/sam/sequence.json").Return(nil)

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should return message with right filename", func(t *testing.T) {
		msg, err := sut.Backup()
		assert.Equal(t, "Completada la còpia de seguretat de la base de dades a /fake/dir/211031-Backup.zip ...", msg)
		assert.Equal(t, nil, err)
	})
}

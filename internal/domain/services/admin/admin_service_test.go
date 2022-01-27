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

func Test_CreateDirectory_exists(t *testing.T) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("ItemExists", mock.Anything).Return(true, nil)

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should not create the directory if exists", func(t *testing.T) {
		msg, err := sut.CreateWorkingDirectory()
		assert.Equal(t, fmt.Sprintf("%sSam v%s    [/fake/dir/211000-Factures del mes d'Octubre]%s", domain.ColorGreen, domain.Version, domain.ColorReset), msg)
		assert.Equal(t, nil, err)
	})
}

func Test_CreateDirectory_does_not_exists(t *testing.T) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("ItemExists", mock.Anything).Return(false, nil)
	mockedOsService.On("CreateDirectory", mock.Anything).Return(nil)
	mockedOsService.On("OpenUrlInBrowser", mock.Anything).Return(nil)

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should create the directory if does not exists", func(t *testing.T) {
		msg, err := sut.CreateWorkingDirectory()
		assert.Equal(t, "Creat el directori /fake/dir/211000-Factures del mes d'Octubre", msg)
		assert.Equal(t, nil, err)
	})
}

func Test_Backup_ok(t *testing.T) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("ItemExists", mock.Anything).Return(true, nil)
	mockedConfigService.On("Get", "dirs.backup").Return("/fake/dir")
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("GetTempDirectory").Return("/tmp/sam", nil)
	mockedOsService.On("ItemExists", "/fake/dir").Return(true, nil)
	mockedOsService.On("CreateZipFile", "/fake/dir/211031-Backup.zip", mock.Anything).Return(nil)
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

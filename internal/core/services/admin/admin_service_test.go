package admin

import (
	"github.com/pjover/sam/internal/core/ports/mocks"
	"github.com/pjover/sam/internal/core/services/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_CreateDirectory(t *testing.T) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedOsService := new(mocks.OsService)
	mockedOsService.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedOsService.On("ItemExists", mock.Anything).Return(false, nil)
	mockedOsService.On("CreateDirectory", mock.Anything).Return(nil)
	mockedOsService.On("OpenUrlInBrowser", mock.Anything).Return(nil)

	sut := NewAdminService(mockedConfigService, mockedOsService, lang.NewCatLangService())

	t.Run("Should return current month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, false)
		assert.Equal(t, "Creat el directori /fake/dir/211000-Factures del mes d'Octubre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return previous month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(true, false)
		assert.Equal(t, "Creat el directori /fake/dir/210900-Factures del mes de Setembre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return next month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, true)
		assert.Equal(t, "Creat el directori /fake/dir/211100-Factures del mes de Novembre", msg)
		assert.Equal(t, nil, err)
	})
}

func Test_Backup_ok(t *testing.T) {
	mockedConfigService := new(mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.backup").Return("/fake/dir")
	mockedOsService := new(mocks.OsService)
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

package services

import (
	os_mocks "github.com/pjover/sam/internal/core/os/mocks"
	ports_mocks "github.com/pjover/sam/internal/core/ports/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateDirectory(t *testing.T) {
	mockedTimeManager := new(os_mocks.TimeManager)
	mockedTimeManager.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedFileManager := new(os_mocks.FileManager)
	mockedFileManager.On("Exists", mock.Anything).Return(false, nil)
	mockedFileManager.On("CreateDirectory", mock.Anything).Return(nil)
	mockedConfigService := new(ports_mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedOpenManager := new(os_mocks.OpenManager)
	mockedOpenManager.On("OnDefaultApp", mock.Anything).Return(nil)

	sut := NewAdminService(mockedConfigService, mockedTimeManager, mockedFileManager, mockedOpenManager)

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

func Test_Backup(t *testing.T) {
	mockedTimeManager := new(os_mocks.TimeManager)
	mockedTimeManager.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedFileManager := new(os_mocks.FileManager)
	mockedConfigService := new(ports_mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.backup").Return("/fake/dir")
	mockedOpenManager := new(os_mocks.OpenManager)

	sut := NewAdminService(mockedConfigService, mockedTimeManager, mockedFileManager, mockedOpenManager)

	t.Run("Should return message with right filename", func(t *testing.T) {
		msg, err := sut.Backup()
		assert.Equal(t, "Completada la còpia de seguretat de la base de dades a/fake/dir/211031-ºBackup.zip ...", msg)
		assert.Equal(t, nil, err)
	})
}

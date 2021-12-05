package services

import (
	env_mocks "github.com/pjover/sam/internal/core/env/mocks"
	os_mocks "github.com/pjover/sam/internal/core/os/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateDirectory(t *testing.T) {
	mockedTimeManager := new(os_mocks.TimeManager)
	mockedTimeManager.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))

	mockedFileManager := new(os_mocks.FileManager)
	mockedFileManager.On("CreateDirectory", mock.Anything).Return("OK", nil)

	mockedConfigManager := new(env_mocks.ConfigManager)
	mockedConfigManager.On("Get", "dirs.home").Return("fake/dir")
	mockedConfigManager.On("Set", mock.Anything, mock.Anything).Return(nil)

	sut := NewAdminService(mockedTimeManager, mockedFileManager, mockedConfigManager)

	t.Run("Should return current month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, false)
		assert.Equal(t, "211000-Factures del mes d'Octubre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return previous month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(true, false)
		assert.Equal(t, "210900-Factures del mes de Setembre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return next month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, true)
		assert.Equal(t, "211100-Factures del mes de Novembre", msg)
		assert.Equal(t, nil, err)
	})
}

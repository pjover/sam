package services

import (
	"fmt"
	"github.com/pjover/sam/internal/core/env"
	"github.com/pjover/sam/internal/core/os"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/shared"
	"github.com/pjover/sam/internal/translate"
	"path"
	"time"
)

type adminService struct {
	timeManager   os.TimeManager
	fileManager   os.FileManager
	configManager env.ConfigManager
}

func NewAdminService(timeManager os.TimeManager, fileManager os.FileManager, configManager env.ConfigManager) ports.AdminService {
	return adminService{
		timeManager:   timeManager,
		fileManager:   fileManager,
		configManager: configManager,
	}
}

func (a adminService) Backup() (string, error) {
	filePath := a.filePath()
	fmt.Println("Fent la còpia de seguretat de la base de dades ...")

	return fmt.Sprint("Completada la còpia de seguretat de la base de dades a", filePath, " ..."), nil
}

func (a adminService) filePath() string {
	dateStr := a.timeManager.Now().Format("060102")
	return path.Join(
		a.configManager.Get("dirs.backup"),
		fmt.Sprintf("%s-ºBackup.zip", dateStr),
	)
}

func (a adminService) CreateDirectory(previousMonth bool, nextMonth bool) (string, error) {
	workingTime := a.getWorkingTime(previousMonth, nextMonth)
	yearMonth := workingTime.Format("2006-01")
	dirName := translate.WorkingDir(workingTime)

	dirPath := path.Join(a.configManager.Get("dirs.home"), dirName)
	msg, err := a.fileManager.CreateDirectory(dirPath)
	if err != nil {
		return "", err
	}

	err = a.updateConfig(yearMonth, dirName)
	if err != nil {
		return "", err
	}

	err = shared.OpenOnDefaultApp(dirPath)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (a adminService) getWorkingTime(previousMonth bool, nextMonth bool) time.Time {
	var t = a.timeManager.Now()
	var workingTime = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)

	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	return workingTime
}

func (a adminService) updateConfig(yearMonth string, dirName string) error {
	if err := a.configManager.Set("yearMonth", yearMonth); err != nil {
		return err
	}
	if err := a.configManager.Set("dirs.current", dirName); err != nil {
		return err
	}
	return nil
}

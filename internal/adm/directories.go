package adm

import (
	"github.com/spf13/viper"
	"sam/internal/translate"
	"sam/internal/util"
	"time"
)

type DirectoryManager interface {
	Create(previousMonth bool, nextMonth bool) error
}

func NewDirectoryManager() DirectoryManager {
	return DirectoryManagerImpl{
		util.SamTimeManager{},
	}
}

type DirectoryManagerImpl struct {
	Timer util.TimeManager
}

func (d DirectoryManagerImpl) Create(previousMonth bool, nextMonth bool) error {
	yearMonth, dirName := d.GetDirConfig(previousMonth, nextMonth)

	err := util.CreateDir(dirName)
	if err != nil {
		return err
	}

	err = d.updateConfig(yearMonth, dirName)
	if err != nil {
		return err
	}
	return nil
}

func (d DirectoryManagerImpl) GetDirConfig(previousMonth bool, nextMonth bool) (string, string) {
	var workingTime = d.getCurrentMonth()
	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	yearMonth := workingTime.Format("2006-01")
	dirName := translate.WorkingDir(workingTime)
	return yearMonth, dirName
}

func (d DirectoryManagerImpl) getCurrentMonth() time.Time {
	var t = d.Timer.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func (d DirectoryManagerImpl) updateConfig(yearMonth string, dirName string) error {

	viper.Set("yearMonth", yearMonth)
	viper.Set("dirs.current", dirName)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

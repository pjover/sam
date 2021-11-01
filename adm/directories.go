package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sam/translate/catalan"
	"sam/util"
	"time"
)

type Directories struct {
	Timer util.TimeManager
}

func (d Directories) CreateDirectory(previousMonth bool, nextMonth bool) error {
	yearMonth, dirName := d.GetDirConfig(previousMonth, nextMonth)

	err := createDir(dirName)
	if err != nil {
		return err
	}

	err = updateConfig(yearMonth, dirName)
	if err != nil {
		return err
	}
	return nil
}

func (d Directories) GetDirConfig(previousMonth bool, nextMonth bool) (string, string) {
	var workingTime = d.getCurrentMonth()
	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	yearMonth := workingTime.Format("2006-01")
	dirName := catalan.WorkingDir(workingTime)
	return yearMonth, dirName
}

func (d Directories) getCurrentMonth() time.Time {
	var t = d.Timer.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func createDir(dirName string) error {
	parentDir := viper.GetString("dirs.home")
	dirPath := path.Join(parentDir, dirName)
	exists, err := fileExists(dirPath)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("El directori de treball", dirPath, "ja existeix")
		return nil
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	fmt.Println("Creat el directori de treball", dirPath, "...")
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func updateConfig(yearMonth string, dirName string) error {

	viper.Set("yearMonth", yearMonth)
	viper.Set("dirs.current", dirName)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

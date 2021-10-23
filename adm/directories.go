package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sam/translate/catalan"
	"time"
)

func CreateDirectory(previousMonth bool, nextMonth bool) error {
	dirName := GetCurrentDirName(previousMonth, nextMonth)
	err := createDir(dirName)
	if err != nil {
		return err
	}

	err = updateConfig(dirName)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentDirName(previousMonth bool, nextMonth bool) string {
	workingTime := time.Now()
	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	return catalan.WorkingDir(workingTime)
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

func updateConfig(dirName string) error {
	viper.Set("dirs.current", dirName)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

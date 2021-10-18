package amd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sam/translate/catalan"
	"time"
)

func Run(previousMonth bool, nextMonth bool) error {
	workingTime := getWorkingTime(previousMonth, nextMonth)
	dir, err := createDir(workingTime)
	if err != nil {
		return err
	}

	err = updateConfigFile(dir)
	if err != nil {
		return err
	}
	return nil
}

func getWorkingTime(previousMonth bool, nextMonth bool) time.Time {
	workingTime := time.Now()
	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	return workingTime
}

func createDir(workingTime time.Time) (string, error) {
	dirName := catalan.WorkingDir(workingTime)

	parentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dirPath := path.Join(parentDir, dirName)
	exists, err := fileExists(dirPath)
	if err != nil {
		return "", err
	}

	if exists {
		fmt.Println("El directori de treball", dirPath, "ja existeix")
		return dirPath, nil
	}

	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		return "", err
	}

	fmt.Println("Creat el directori de treball", dirPath, "...")
	return dirPath, nil
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

func updateConfigFile(dir string) error {
	viper.Set("dir", dir)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

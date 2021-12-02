package util

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
)

func GetWorkingDirectory() string {
	return path.Join(viper.GetString("dirs.home"), viper.GetString("dirs.current"))
}

func CreateDir(dirName string) error {
	parentDir := viper.GetString("dirs.home")
	dirPath := path.Join(parentDir, dirName)
	exists, err := FileExists(dirPath)
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

	fmt.Println("Creat el directori de treball", dirPath)
	return nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

package os

import (
	"fmt"
	"os"
)

type FileManager interface {
	CreateDirectory(dirName string) (msg string, err error)
	FileExists(path string) (bool, error)
}

type fileManager struct {
}

func NewFileManager() FileManager {
	return fileManager{}
}

func (f fileManager) CreateDirectory(dirPath string) (msg string, err error) {
	exists, err := f.FileExists(dirPath)
	if err != nil {
		return fmt.Sprint("Error d'accés:", err), err
	}

	if exists {
		return fmt.Sprint("El directori de treball", dirPath, "ja existeix"), nil
	}

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Sprint("Error al crear directori:", err), err
	}

	return fmt.Sprint("Creat el directori de treball", dirPath), nil
}

func (f fileManager) FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

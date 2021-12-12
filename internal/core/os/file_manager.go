package os

import (
	"archive/zip"
	"io"
	"os"
	"path"
)

type FileManager interface {
	CreateDirectory(dirPath string) error
	Exists(itemPath string) (bool, error)
	GetTempDirectory() (string, error)
	Zip(zipFilePath string, files []string) error
}

type fileManager struct {
}

func NewFileManager() FileManager {
	return fileManager{}
}

func (f fileManager) CreateDirectory(dirPath string) error {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (f fileManager) Exists(itemPath string) (bool, error) {
	_, err := os.Stat(itemPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (f fileManager) GetTempDirectory() (string, error) {
	dirPath := path.Join(os.TempDir(), "sam")
	exists, err := f.Exists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		return dirPath, nil
	}
	if err := f.CreateDirectory(dirPath); err != nil {
		return "", err
	}
	return dirPath, nil
}

// Zip compresses one or many files into a single zip archive file.
// Param 1: zipFilePath is the output zip file's absolute path.
// Param 2: files is a list of files absolute path to add to the zip.
// From: https://golangcode.com/create-zip-files-in-go/
func (f fileManager) Zip(zipFilePath string, files []string) error {

	newZipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer func(newZipFile *os.File) {
		err := newZipFile.Close()
		if err != nil {
			panic(err)
		}
	}(newZipFile)

	zipWriter := zip.NewWriter(newZipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(zipWriter)

	// Add files to zip
	for _, file := range files {
		if err = f.addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func (f fileManager) addFileToZip(zipWriter *zip.Writer, filePath string) error {

	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			panic(err)
		}
	}(fileToZip)

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

package utils

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filePath string) {
	file, error := os.Create(filePath)

	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
}

func CheckFileExists(filePath string) bool {
	_, error := os.Stat(filePath)

	if errors.Is(error, os.ErrNotExist) {
		log.Fatal(error)
	}
	return true
}

func DeleteFile(filePath string) {
	CheckFileExists(filePath)

	error := os.Remove(filePath)
	if error != nil {
		log.Fatal(error)
	}
}

func WriteFileContent(filePath string, content string) {

	file, error := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()

	if _, error := file.WriteString(content); error != nil {
		log.Fatal(error)
	}
}

func GetFileStats(filePath string) fs.FileInfo {
	CheckFileExists(filePath)
	fileStats, error := os.Stat(filePath)

	if error != nil {
		log.Fatal(error)
	}

	return fileStats
}

func ReadFileContent(filePath string) string {
	CheckFileExists(filePath)
	content, error := ioutil.ReadFile(filePath)

	if error != nil {
		log.Fatal(error)
	}
	return string(content)
}

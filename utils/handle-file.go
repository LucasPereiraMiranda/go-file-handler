package utils

import (
	"errors"
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

package utils

import (
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

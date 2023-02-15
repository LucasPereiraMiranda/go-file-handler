package main

import (
	"fmt"

	"github.com/LucasPereiraMiranda/go-file-handler/utils"
)

func main() {

	filePath := "/home/lucas/dev/personal/go-file-handler/files/data.txt"
	utils.CreateFile(filePath)

	utils.CheckFileExists(filePath)

	utils.WriteFileContent(filePath, "data to add in file\n")
	content := utils.ReadFileContent(filePath)

	fmt.Println("content:", content)

	fileStats := utils.GetFileStats(filePath)

	fmt.Println("file name:", fileStats.Name(), "file size:", fileStats.Size())

	utils.DeleteFile(filePath)

}

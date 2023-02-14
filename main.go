package main

import (
	"github.com/LucasPereiraMiranda/go-file-handler/utils"
)

func main() {

	filePath := "/home/lucas/dev/personal/go-file-handler/files/data.txt"
	utils.CreateFile(filePath)

}

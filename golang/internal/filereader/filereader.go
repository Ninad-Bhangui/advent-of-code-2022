package filereader

import (
	"log"
	"os"
)

func ReadFileToString(inputFilepath string) string {
	byteData, err := os.ReadFile(inputFilepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteData)
}

package utils

import (
	"os"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}


func SaveFile(fileData []byte, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	PanicIfError(err)

	file.Write(fileData)
}


func Unzip(fileData []byte, password string) {
	
}

// Generic SampleInfo struct dump using reflection

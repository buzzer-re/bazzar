package utils

import (
	"os"
	"bytes"
	"io/ioutil"
	"fmt"

	"github.com/alexmullins/zip"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}


func SaveFile(fileData []byte, fileName string) {
	file, err := os.Create(fileName)
	PanicIfError(err)
	defer file.Close()
	file.Write(fileData)
}

func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error: %s", err)
		os.Exit(1)
	}
}

func Unzip(fileData []byte, password string) ([]byte, string) {
	zipReader, err := zip.NewReader(bytes.NewReader(fileData), int64(len(fileData)))
	ExitIfError(err)
	
    for _, zipFile := range zipReader.File {
		if zipFile.IsEncrypted() {
			zipFile.SetPassword(password)
		}
		
		f, err := zipFile.Open()
		ExitIfError(err)
		defer f.Close()

		unzippedFileBytes, err := ioutil.ReadAll(f)
		ExitIfError(err)

       return unzippedFileBytes, zipFile.Name // this is unzipped file bytes
	}

	return nil, ""
}

// Generic SampleInfo struct dump using reflection

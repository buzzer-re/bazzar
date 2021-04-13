package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/alexmullins/zip"
)

const (
	HOST_RUlE     = `(([a-zA-Z]){0,5}:\/\/)?.+(\.[a-zA-Z0-9_@./#&+-].+?).+?(\/|)`
	URL_PATH_RULE = `(([a-zA-Z]){0,5}:\/\/)?.+(\.[a-zA-Z0-9_+-]+){1,}(:[0-9]{1,5})?(\/[a-zA-Z0-9_@.?/#&+-]+)+\/?$`
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
		fmt.Fprintf(os.Stderr, "Error: %s", err)
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

func IsHost(url string) bool {
	match, _ := regexp.MatchString(HOST_RUlE, url)
	return match
}

func IsFullUrl(url string) bool {
	match, _ := regexp.MatchString(URL_PATH_RULE, url)
	return match
}

func CleanHost(hostname string) string {
	hostname = strings.ReplaceAll(hostname, "http://", "")
	hostname = strings.ReplaceAll(hostname, "https://", "")
	hostname = strings.ReplaceAll(hostname, "ftp://", "")
	hostname = strings.ReplaceAll(hostname, "/", "")

	return hostname
}

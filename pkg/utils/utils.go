package utils

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}


func SaveFile(fileData []byte, fileName string) {
	
}
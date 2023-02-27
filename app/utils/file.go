package utils

import (
	"io"
	"log"
	"os"
)

func CraeteDirectory(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func CreateFileForce(path string, filename string) *os.File {
	file, err := os.Create(path + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func CreateFile(path string, filename string) *os.File {

	if _, err := os.Stat(path + filename); err == nil {
		file, err := os.OpenFile(path+filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

		if err != nil {
			log.Fatal(err)
		}

		return file
	}

	file := CreateFileForce(path, filename)
	return file
}

func MultiWrite(out *os.File, file *os.File) io.Writer {
	multiOutput := io.MultiWriter(out, file)
	return multiOutput
}

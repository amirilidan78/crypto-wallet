package logger

import (
	"log"
	"os"
)

func openFile(file string) *os.File {

	f, err := os.OpenFile("./logs/"+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return f
}

func closeFile(f *os.File) {
	f.Close()
}

func Write(file string, message string) {

	f := openFile(file)

	defer closeFile(f)

	log.SetOutput(f)

	log.Println(message)
}

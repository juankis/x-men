package utils

import (
	"log"
	"os"
)

//WriteInLog save erro in log
func WriteInLog(error string) {
	f, err := os.OpenFile("logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	log.Println(error)
	defer f.Close()
}

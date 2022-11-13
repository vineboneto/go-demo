package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func fakeError() (string, error) {
	return "", errors.New("Testing Error")
}

func main() {
	_, err := fakeError()

	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	defer file.Close()

	ErrorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger.Println(err)

	fmt.Println("Hello World")
}

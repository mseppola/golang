package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("/home/mseppola/go/log/logging", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("Logging to a file in Go!")
}

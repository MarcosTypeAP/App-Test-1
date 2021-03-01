package main

import (
	"fmt"
	"os"
)

//ErrorPrinter prints the error
func ErrorPrinter(err error) {
	if err != nil {
		fmt.Println("error ->", err)
	}
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "1700"
	}

	RunServer("", port)
}

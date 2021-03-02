package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//ErrorPrinter prints the error
func ErrorPrinter(err error) {
	if err != nil {
		fmt.Println("error ->", err)
	}
}

type xd struct {
	Name string `json:"name,omitempty"`
	XD   string `json:"xd,omitempty"`
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	x := xd{Name: "namexd"}
	d, _ := json.Marshal(x)
	fmt.Println(string(d))
	RunServer("", port)
}

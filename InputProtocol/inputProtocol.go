package inputprotocol

import (
	"fmt"
	"log"
	"os"
)

// CreateInputProtocol creates ModuleNameInput.swift
func CreateInputProtocol(moduleName string) *os.File {
	input, err := os.Create(moduleName + "Input" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	inputBody := "protocol " + moduleName + "Input" + " { \n	<#T##Input's protocol code#> \n}"
	fmt.Println(inputBody)
	_, err = input.WriteString(inputBody)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

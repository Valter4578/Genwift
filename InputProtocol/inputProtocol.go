package inputprotocol

import (
	"log"
	"os"
)

// CreateInputProtocol creates ModuleNameInput.swift
func CreateInputProtocol(moduleName string) *os.File {
	input, err := os.Create(moduleName + "/Presenter/Protocols/" + moduleName + "Input" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	inputBody := "protocol " + moduleName + "Input" + " { \n	<#T##Input's protocol code#> \n}"
	_, err = input.WriteString(inputBody)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

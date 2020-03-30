package inputprotocol

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateInputProtocol creates ModuleNameInput.swift
func CreateInputProtocol(moduleName string) *os.File {
	fileName := moduleName + "/Presenter/Protocols/" + moduleName + "Input" + ".swift"
	input, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	color.Yellow("Sucsessfully create directory: %v", fileName)

	inputBody := "protocol " + moduleName + "Input" + " { \n	<#T##Input's protocol code#> \n}"
	_, err = input.WriteString(inputBody)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

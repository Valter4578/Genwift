package outputprotocol

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateOutputProtocol create ModuleNameInputProtocol.swift
func CreateOutputProtocol(moduleName string) *os.File {
	fileName := moduleName + "/Presenter/Protocols/" + moduleName + "Output" + ".swift"
	output, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	color.Yellow("Sucsessfully create directory: %v", fileName)

	outputBody := "protocol " + moduleName + "Output" + " { \n	<#T##Output's protocol code#> \n}"
	_, err = output.WriteString(outputBody)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

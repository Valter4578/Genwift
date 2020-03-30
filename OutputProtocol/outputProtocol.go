package outputprotocol

import (
	"log"
	"os"
)

// CreateOutputProtocol create ModuleNameInputProtocol.swift
func CreateOutputProtocol(moduleName string) *os.File {
	output, err := os.Create(moduleName + "/Presenter/Protocols/" + moduleName + "Output" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	outputBody := "protocol " + moduleName + "Output" + " { \n	<#T##Output's protocol code#> \n}"
	_, err = output.WriteString(outputBody)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

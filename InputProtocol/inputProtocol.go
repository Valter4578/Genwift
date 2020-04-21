package inputprotocol

import (
	"fmt"
	"log"
	"os"
)

// CreateInputProtocol creates ModuleNameInput.swift
func CreateInputProtocol(moduleName string) *os.File {
	fileName := moduleName + "Module" + "/Presenter/Protocols/" + moduleName + "Input" + ".swift"
	input, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	body := fmt.Sprintf(`
protocol %vInput: class {
	<#T##Input's protocol code#>
}
`, moduleName)
	_, err = input.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

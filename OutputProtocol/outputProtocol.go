package outputprotocol

import (
	"fmt"
	"log"
	"os"
)

// CreateOutputProtocol create ModuleNameInputProtocol.swift
func CreateOutputProtocol(moduleName string) *os.File {
	output, err := os.Create(moduleName + "/Presenter/Protocols/" + moduleName + "Output" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	body := fmt.Sprintf(`
protocol %vOutput {
	<#T##Output's protocol code#>

}`, moduleName)

	_, err = output.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

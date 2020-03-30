package viewcontroller

import (
	"log"
	"os"
)

// CreateViewController creates ModuleNameViewController.swift
func CreateViewController(moduleName string) *os.File {
	vc, err := os.Create(moduleName + "/View/" + moduleName + "ViewController" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	vcBody := "class " + moduleName + "ViewController: " + moduleName + "Input { \n	<#T##ViewController's code#> \n}"

	_, err = vc.WriteString(vcBody)
	if err != nil {
		log.Fatal(err)
	}

	return vc
}

package viewcontroller

import (
	"log"
	"os"
)

// CreateViewController creates ModuleNameViewController.swift
func CreateViewController(moduleName string) *os.File {
	fileName := moduleName + "Module" + "/View/" + moduleName + "ViewController" + ".swift"
	vc, err := os.Create(fileName)
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

package viewcontroller

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateViewController creates ModuleNameViewController.swift
func CreateViewController(moduleName string) *os.File {
	fileName := moduleName + "/View/" + moduleName + "ViewController" + ".swift"
	vc, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	color.Yellow("Sucsessfully create directory: %v", fileName)

	vcBody := "class " + moduleName + "ViewController: " + moduleName + "Input { \n	<#T##ViewController's code#> \n}"

	_, err = vc.WriteString(vcBody)
	if err != nil {
		log.Fatal(err)
	}

	return vc
}

package viewcontroller

import (
	"fmt"
	"log"
	"os"
)

// CreateViewController creates ModuleNameViewController.swift
func CreateViewController(moduleName string) *os.File {
	vc, err := os.Create(moduleName + "ViewController" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	vcBody := "class " + moduleName + "ViewController: " + moduleName + "Input { \n	<#T##Presenter's code#> \n}"
	fmt.Println(vcBody)
	_, err = vc.WriteString(vcBody)
	if err != nil {
		log.Fatal(err)
	}

	return vc
}

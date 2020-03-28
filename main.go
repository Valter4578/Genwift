package main

import (
	"fmt"

	// "io/ioutil"
	"log"
	"os"
)

func main() {
	arguments := os.Args
	moduleName := os.Args[1]
	fmt.Println(arguments)
	fmt.Println(moduleName)

	presenterName := moduleName + "Presenter.swift"
	presenter, err := os.Create(presenterName)
	if err != nil {
		log.Fatal(err)
	}

	presenterBody := "class " + moduleName + "Presenter: " + moduleName + "Output { \n<#T##Presenter's code#> \n}"
	fmt.Println(presenterBody)
	_, err = presenter.WriteString(presenterBody)
	if err != nil {
		log.Fatal(err)
	}

	vc, err := os.Create(moduleName + "ViewController" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	vcBody := "class " + moduleName + "ViewController: " + moduleName + "Input { \n<#T##Presenter's code#> \n}"
	fmt.Println(vcBody)
	_, err = vc.WriteString(vcBody)
	if err != nil {
		log.Fatal(err)
	}

	input, err := os.Create(moduleName + "Input" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	inputBody := "protocol " + moduleName + "Input" + " { \n<#T##Input's protocol code#> \n}"
	fmt.Println(inputBody)
	_, err = input.WriteString(inputBody)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create(moduleName + "Output" + ".swift")
	if err != nil {
		log.Fatal(err)
	}

	outputBody := "protocol " + moduleName + "Output" + " { \n<#T##Output's protocol code#> \n}"
	fmt.Println(inputBody)
	_, err = output.WriteString(outputBody)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(presenter)
	log.Println(vc)
	log.Println(input)
	log.Println(output)

}

package presenter

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// CreatePresenter creates ModuleNamePresenter.swift file
func CreatePresenter(moduleName string) *os.File {
	fileName := moduleName + "/Presenter/" + moduleName + "Presenter.swift"
	presenter, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	color.Yellow("Sucsessfully create directory: %v", fileName)

	presenterBody := "class " + moduleName + "Presenter: " + moduleName + "Output {\n	weak var view:" + moduleName + "Input? \n	<#T##Presenter's code#> \n}\ninit(view: " + moduleName + ") {\n	self.view = view \n}"

	_, err = presenter.WriteString(presenterBody)
	if err != nil {
		log.Fatal(err)
	}
	return presenter
}

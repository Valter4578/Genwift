package presenter

import (
	"fmt"
	"log"
	"os"
)

// CreatePresenter creates ModuleNamePresenter.swift file
func CreatePresenter(moduleName string) *os.File {
	presenterName := moduleName + "Presenter.swift"
	presenter, err := os.Create(presenterName)
	if err != nil {
		log.Fatal(err)
	}

	presenterBody := "class " + moduleName + "Presenter: " + moduleName + "Output {\n	weak var view:" + moduleName + "Input? \n	<#T##Presenter's code#> \n}\ninit(view: " + moduleName + ") {\n	self.view = view \n}"
	fmt.Println(presenterBody)
	_, err = presenter.WriteString(presenterBody)
	if err != nil {
		log.Fatal(err)
	}
	return presenter
}
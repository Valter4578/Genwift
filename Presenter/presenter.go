package presenter

import (
	"log"
	"os"
)

// CreatePresenter creates ModuleNamePresenter.swift file
func CreatePresenter(moduleName string) *os.File {
	fileName := moduleName + "Module/" + "/Presenter/" + moduleName + "Presenter.swift"
	presenter, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	body := "	weak var view: " + moduleName + "Input? \n 	<#T##Presenter's code#> \n\n 	init(view: " + moduleName + "Input) {\n		self.view = view\n	}\n"
	presenterBody := "class " + moduleName + "Presenter: " + moduleName + "Output {\n" + body + "}"

	_, err = presenter.WriteString(presenterBody)
	if err != nil {
		log.Fatal(err)
	}
	return presenter
}

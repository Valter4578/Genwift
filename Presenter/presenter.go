package presenter

import (
	"fmt"
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

	body := fmt.Sprintf(`
class %vPresenter: %vOutput {
	// MARK:- Dependencies 
	weak var view: %vInput!
	// MARK:- <#T## mark's name #> 
	<#T##Presenter's code#>

	// MARK:- Initializers 
	init(view: %vInput) {
		self.view = view
	}
}`, moduleName, moduleName, moduleName, moduleName)

	_, err = presenter.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}
	return presenter
}

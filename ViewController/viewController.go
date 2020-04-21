package viewcontroller

import (
	"fmt"
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

	body := fmt.Sprintf(`
class %vViewController: UIViewController { 
	// MARK:- Dependencies 
	var presenter: %vOutput! 

	// MARK:- Lifecycle 
	override func viewDidLoad() {
		super.viewDidLoad()

	}

	// MARK:- <#T## mark's name#>
}

// MARK: %vInput Implementation
extension %vViewController: LoginInput {

}
`, moduleName, moduleName, moduleName, moduleName)

	_, err = vc.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}

	return vc
}

package viewcontroller

import (
	"fmt"
	"log"
	"os"
)

// CreateViewController creates ModuleNameViewController.swift
func CreateViewController(moduleName string) *os.File {
	vc, err := os.Create(moduleName + "/View/" + moduleName + "ViewController" + ".swift")
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
`, moduleName, moduleName)

	_, err = vc.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}

	return vc
}

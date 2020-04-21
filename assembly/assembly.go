package assembly

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateAssembly creates ModuleNameAssembly.swift
func CreateAssembly(moduleName string) *os.File {
	fileName := moduleName + "Module" + "/" + moduleName + "Assembly.swift"
	assembly, err := os.Create(fileName)
	if err != nil {
		color.Red("Can not create assembly for %v", moduleName)
		log.Fatal(err)
	}

	body := fmt.Sprintf(`
class %vAssembly {
	static func configureModule() -> %vViewController {
		let view = %vViewController()
		let presenter = %vPresenter(view: view)
		view.presenter = presenter 
		return view 
	}
}
`, moduleName, moduleName, moduleName, moduleName)

	_, err = assembly.WriteString(body)
	if err != nil {
		log.Fatal(err)
	}

	return assembly
}

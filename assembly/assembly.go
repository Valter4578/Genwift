package assembly

import (
	"log"
	"os"

	"github.com/fatih/color"
)

// CreateAssembly creates ModuleNameAssembly.swift
func CreateAssembly(moduleName string) *os.File {
	fileName := moduleName + "/" + moduleName + "Assembly.swift"
	assembly, err := os.Create(fileName)
	if err != nil {
		color.Red("Can not create assembly for %v", moduleName)
		log.Fatal(err)
	}
	color.Yellow("Sucsessfully create directory: %v", fileName)
	body := "		let view = " + moduleName + "ViewController()\n		let presenter = " + moduleName + "Presenter(view: view)\n		view.presenter = presenter\n		return view"
	assemblyBody := "class " + moduleName + "Assembly" + " { \n class func configureModule() -> " + moduleName + "ViewController {\n" + body + "\n	}\n}"

	_, err = assembly.WriteString(assemblyBody)
	if err != nil {
		log.Fatal(err)
	}

	return assembly
}

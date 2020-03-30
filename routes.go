package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func createDirectories(moduleName string) {
	err := os.Mkdir(moduleName+"Module", 0777)
	if err != nil {
		color.Red("Can't create directory:")
		log.Fatal(err)
	}
	err = os.Mkdir(moduleName+"Module"+"/Presenter", 0777)
	if err != nil {
		color.Red("Can't create directory:")
		log.Fatal(err)
	}

	err = os.Mkdir(moduleName+"Module"+"/View", 0777)
	if err != nil {
		color.Red("Can't create directory:")
		log.Fatal(err)
	}

	err = os.Mkdir(moduleName+"Module"+"/Presenter/Protocols", 0777)
	if err != nil {
		color.Red("Can't create directory:")
		log.Fatal(err)
	}
}

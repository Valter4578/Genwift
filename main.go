package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"

	"genwift/assembly"
	"genwift/inputprotocol"
	"genwift/outputprotocol"
	"genwift/presenter"
	"genwift/viewcontroller"
)

func main() {
	moduleName := os.Args[1]
	architecture := os.Args[2]
	fmt.Println(architecture)

	if architecture == "MVP" {
		// Create file stucture
		err := os.Mkdir(moduleName, 0777)

		err = os.Mkdir(moduleName+"/Presenter", 0777)
		if err != nil {
			color.Red("Can't create directory:")
			log.Fatal(err)
		}

		err = os.Mkdir(moduleName+"/View", 0777)
		if err != nil {
			color.Red("Can't create directory:")
			log.Fatal(err)
		}

		err = os.Mkdir(moduleName+"/Presenter/Protocol", 0777)
		if err != nil {
			color.Red("Can't create directory:")
			log.Fatal(err)
		}

		cyan := color.New(color.FgCyan)
		cyan.Printf("Creating module %v....\n", moduleName)

		presenter := presenter.CreatePresenter(moduleName)
		color.Magenta("Presenter was created")

		vc := viewcontroller.CreateViewController(moduleName)
		color.Magenta("View Controller was created")

		input := inputprotocol.CreateInputProtocol(moduleName)
		color.Magenta("Input was created")

		output := outputprotocol.CreateOutputProtocol(moduleName)
		color.Magenta("Output was created")

		assembly := assembly.CreateAssembly(moduleName)
		color.Magenta("Assembly was created")

		color.Green("Sucsessfully create MVP module %v:", moduleName)

		log.Println(presenter)
		log.Println(vc)
		log.Println(input)
		log.Println(output)
		log.Println(assembly)
	} else {
		color.Red("ERROR: Unknown architecture %v", architecture)
	}

}

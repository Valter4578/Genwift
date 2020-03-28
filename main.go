package main

import (
	"fmt"
	"log"
	"os"

	"genwift/inputprotocol"
	"genwift/outputprotocol"
	"genwift/presenter"
	"genwift/viewcontroller"
)

func main() {
	arguments := os.Args
	moduleName := os.Args[1]
	fmt.Println(arguments)
	fmt.Println(moduleName)

	presenter := presenter.CreatePresenter(moduleName)
	vc := viewcontroller.CreateViewController(moduleName)
	input := inputprotocol.CreateInputProtocol(moduleName)
	output := outputprotocol.CreateOutputProtocol(moduleName)

	log.Println(presenter)
	log.Println(vc)
	log.Println(input)
	log.Println(output)

}

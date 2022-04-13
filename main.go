package main

import (
	"fmt"
	"log"

	"github.com/ricardoarturo12/simple_cli/commands"
)

func main() {
	fmt.Println("hola mundo")

	for {

		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}
	}

	log.Println("Closing...")
}

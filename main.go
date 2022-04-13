package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/ricardoarturo12/simple_cli/commands"
)

func main() {

	var expenses []float32
	var export string

	// flag.StringVar(&export, "export", "", "Exports the details to .txt")

	flag.StringVar(&export, "export", "", "Exports the details to .txt")

	// parsea a esa variable
	flag.Parse()

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			continue
		}

		expenses = append(expenses, float32(expense))
	}

	if export == "" {
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}
}

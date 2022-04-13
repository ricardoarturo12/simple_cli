package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ricardoarturo12/simple_cli/expenses"
)

// para leer la consola
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Print("-> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32) {

	/* string builder -> variable que captura valores
	y mantiene en buffer*/
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}

	max, min, avg, total := expensesDetails(expensesList)

	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expenses: %6.2f \n", expense))

		if i == len(expensesList)-1 {
			builder.WriteString("")
			builder.WriteString("=========================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f \n", total))
			builder.WriteString(fmt.Sprintf("Max: %6.2f \n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f \n", min))
			builder.WriteString(fmt.Sprintf("Promedio: %6.2f \n", avg))
			builder.WriteString("=========================\n")
		}

	}
	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, avg, total float32) {

	if len(expensesList) == 0 {
		return
	}

	// acá actualiza el valor
	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	avg = expenses.Average(expensesList...)
	total = expenses.Sum(expensesList...)

	return
}

func Export(fileName string, list []float32) error {
	// flags -export
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}

	// guarda el archivo
	return w.Flush()
}

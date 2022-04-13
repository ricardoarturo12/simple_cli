package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

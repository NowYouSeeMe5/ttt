package ui

import "fmt"

type ConsoleIO struct{}

func (c ConsoleIO) GetInput() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func (c ConsoleIO) Print(message string) {
	fmt.Print(message)
}

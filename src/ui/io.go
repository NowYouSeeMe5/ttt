package ui

type IO interface {
	Print(message string)
	GetInput() string
}

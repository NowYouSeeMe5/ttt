package io

import "fmt"

type Prompter func(message string) interface{}
type InputGetter func() int
type ErrorHandler func(expectedResults []int, result interface{}) bool

type TttInput struct {
	prompter     Prompter
	inputGetter  InputGetter
	errorHandler ErrorHandler
}

var expectedInputBoardSize = []int{3, 4, 5}
var expectedInputHowManyPlayers = []int{0, 1, 2}
var expectedInputWhichPlayer = []int{1, 2}

func NewTttInput() *TttInput {
	TttInput := new(TttInput)

	TttInput.prompter = TttInput.prompt
	TttInput.errorHandler = TttInput.handleErrors
	TttInput.inputGetter = TttInput.getInput

	return TttInput
}

func (t TttInput) BoardSize() int {
	return t.AskQuestion(BoardSizeQuestion, expectedInputBoardSize)
}

func (t TttInput) HowManyPlayers() int {
	return t.AskQuestion(HowManyPlayersQuestion, expectedInputHowManyPlayers)
}

func (t TttInput) AskQuestion(question string, expectedInput []int) int {
	input := t.questionPrompt(question)

	if t.errorHandler(expectedInput, input) {
		return input
	}

	return t.AskQuestion(question, expectedInput)
}

func (t TttInput) questionPrompt(question string) int {
	t.prompter(question)
	return t.inputGetter()
}

func (t TttInput) prompt(message string) interface{} {
	fmt.Println(message)
	return nil
}

func (t TttInput) getInput() int {
	var i int
	fmt.Scanf("%d", &i)
	return i
}

func (t TttInput) handleErrors(expectedResults []int, result interface{}) bool {
	for _, v := range expectedResults {
		if result == v {
			return true
		}
	}
	fmt.Println("You have entered in something unexpected!!!")
	return false
}

func (t *TttInput) InjectPrompter(prompter Prompter) {
	t.prompter = prompter
}

func (t *TttInput) InjectGetter(getter InputGetter) {
	t.inputGetter = getter
}

func (t *TttInput) InjectErrorHandler(errorHandler ErrorHandler) {
	t.errorHandler = errorHandler
}

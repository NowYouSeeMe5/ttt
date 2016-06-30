package player

import (
	"fmt"

	"ttt/src/board"
	"ttt/src/io"
)

type HumanPlayer struct{}

var Question = "Please select a move from 1-%d"

func (h *HumanPlayer) Move(board board.Board) int {
	question := FormQuestion(board)
	possibleMoves := board.PossibleMoves()

	tttInput := io.NewTttInput()

	return GetInput(question, possibleMoves, tttInput)
}

func GetInput(question string, possibleMoves []int, input io.Input) int {
	return input.AskQuestion(question, possibleMoves) - 1
}

func FormQuestion(board board.Board) string {
	spaces := board.Spaces()
	size := len(spaces)

	return fmt.Sprintf(Question, size)
}

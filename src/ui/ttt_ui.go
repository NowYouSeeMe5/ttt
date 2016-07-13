package ui

import (
	"bytes"
	"fmt"
	"strconv"
	"ttt/src/board"
)

const (
	BoardSizeQuestion      = "Choose a board size from the following: %d\n"
	HowManyPlayersQuestion = "How many human players are you? Choose from the following: %d\n"
	ErrorMessage           = "That input was incorrect.\n"
	PlayerVictory          = "Well done %s, congrats on your victory!!!\n"
	TieGame                = "That's a nice Tie!!!\n"
	GetMove                = "Please select a move from the following: %d\n"

	leftSpace = "%s"
	space     = "|%s"
)

var pieces = map[int]string{
	0:  "_",
	1:  "X",
	-1: "O",
}

type TttUI struct {
	io        IO
	validator *TttValidator
}

func NewTttUI(io IO) *TttUI {
	ui := new(TttUI)
	ui.io = io
	ui.validator = new(TttValidator)

	return ui
}

func (u TttUI) PrintBoard(board *board.TttBoard) {
	boardString := renderBoard(board)
	u.io.Print(boardString)
}

func renderBoard(board *board.TttBoard) string {
	size := board.Size()
	spaces := board.Spaces()

	var buffer bytes.Buffer

	for y := 0; y < size; y++ {
		currentY := y * size
		buffer.WriteString(formatSpace(leftSpace, spaces, currentY))
		for x := 1; x < size; x++ {
			buffer.WriteString(formatSpace(space, spaces, currentY+x))
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("\n")
	return buffer.String()
}

func formatSpace(space string, spaces []int, index int) string {
	return fmt.Sprintf(space, getPiece(spaces[index]))
}

func getPiece(squareState int) string {
	return pieces[squareState]
}

func (u TttUI) PrintEndGameMessage(winner int) {
	endGameMessage := endGameMessage(winner)
	u.io.Print(endGameMessage)
}

func endGameMessage(winner int) string {
	if winner != 0 {
		return fmt.Sprintf(PlayerVictory, getPiece(winner))
	}

	return TieGame
}

func (u TttUI) GetMove(availableMoves []int) int {
	return u.prompt(GetMove, availableMoves)
}

func (u TttUI) GetBoardSize(availableSizes []int) int {
	return u.prompt(BoardSizeQuestion, availableSizes)
}

func (u TttUI) GetNumberOfPlayers(availablePlayers []int) int {
	return u.prompt(HowManyPlayersQuestion, availablePlayers)
}

func (u TttUI) prompt(message string, validInput []int) int {
	question := fmt.Sprintf(message, validInput)
	u.io.Print(question)

	response, _ := strconv.Atoi(u.io.GetInput())

	if !u.validator.Validate(response, validInput) {
		u.prompt(message, validInput)
	}

	return response
}

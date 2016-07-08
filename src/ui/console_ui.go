package ui

import (
	"bytes"
	"fmt"
	"ttt/src/board"
)

const (
	BoardSizeQuestion      = "What size board would you like to play on? Enter 3, 4, or 5"
	HowManyPlayersQuestion = "How many human players are you? Enter 0, 1, or 2"
	ErrorMessage           = "Please enter the appropriate information"
	PlayerVictory          = "Well done %s, congrats on your victory!!!"
	TieGame                = "That's a nice Tie!!!"
	GetMove                = "Please select a move from 1-%d"

	leftSpace = "%s"
	space     = "|%s"
)

var pieces = map[int]string{
	0:  "_",
	1:  "X",
	-1: "O",
}

type consoleUI struct {
	io        IO
	validator InputValidator
}

func NewConsoleUI(io IO) *consoleUI {
	ui := new(consoleUI)
	ui.io = io

	return ui
}

func (c consoleUI) PrintBoard(board *board.TttBoard) {
	boardString := renderBoard(board)
	c.io.Print(boardString)
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
	return buffer.String()
}

func formatSpace(space string, spaces []int, index int) string {
	return fmt.Sprintf(space, getPiece(spaces[index]))
}

func getPiece(squareState int) string {
	return pieces[squareState]
}

func prompt(message, acceptedInput) int {

}

//func (c consoleUI) PrintEndGameMessage(winner int) {
//endGameMessage := c.endGameMessage(winner)
//c.io.Print(endGameMessage)
//}

//func (c consoleUI) endGameMessage(winner int) string {
//if winner != 0 {
//return fmt.Sprintf(PlayerVictory, getPiece(winner))
//}
//return TieGame
//}

//func (c consoleUI) GetMove(availableMoves int[]) int {
//c.io.Print(GetMove)

//response := c.io.GetInput()

//err := c.validator(response, availableMoves)

//if err != nil {
//c.io.Print(errorMessage)
//}

//return response
//}

//func (c consoleUI) GetBoardSize(availableSizes int[]) int {}
//func (c consoleUI) GetNumberOfPlayers() int {}

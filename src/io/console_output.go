package io

import (
	"bytes"
	"fmt"
	"math"
	. "ttt/src/board"
)

type TttOutput struct{}

const (
	BoardSizeQuestion      = "What size board would you like to play on? Enter 3 or higher"
	HowManyPlayersQuestion = "How many human players are you? Enter 0, 1, or 2"
	ErrorMessage           = "Please enter the appropriate information"
	PlayerVictory          = "Well done %s, congrats on your victory!!!"
	TieGame                = "That's a nice Tie!!!"

	leftSpace = "%s"
	space     = "|%s"
)

var pieces = map[int]string{
	0:  "_",
	1:  "X",
	-1: "O",
}

func (t TttOutput) Print(message string) {
	fmt.Println(message)
}

func (t TttOutput) PrintBoard(board Board) {
	boardString := t.RenderBoard(board)
	t.Print(boardString)
}

func (t TttOutput) RenderBoard(board Board) string {
	spaces := board.Spaces()
	lines := renderLines(spaces)

	return lines
}

func renderLines(spaces []int) string {
	size := size(spaces)
	var buffer bytes.Buffer

	for y := 0; y < size; y++ {
		currentY := y * size
		buffer.WriteString(fmt.Sprintf(leftSpace, getPiece(spaces[currentY])))
		for x := 1; x < size; x++ {
			buffer.WriteString(fmt.Sprintf(space, getPiece(spaces[currentY+x])))
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func getPiece(squareState int) string {
	return pieces[squareState]
}

func (t TttOutput) PrintEndGameMessage(winner int) {
	endGameMessage := t.EndGameMessage(winner)
	t.Print(endGameMessage)
}

func (t TttOutput) EndGameMessage(winner int) string {
	if winner != 0 {
		return fmt.Sprintf(PlayerVictory, getPiece(winner))
	}
	return TieGame
}

func size(spaces []int) int {
	squares := float64(len(spaces))
	return int(math.Sqrt(squares))
}

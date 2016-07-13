package player

import (
	"github.com/NowYouSeeMe5/ttt/src/board"
	"github.com/NowYouSeeMe5/ttt/src/ui"
)

var humanReadabilityPadding = 1

type HumanPlayer struct {
	ui *ui.TttUI
}

func NewHumanPlayer(ui *ui.TttUI) *HumanPlayer {
	newPlayer := new(HumanPlayer)
	newPlayer.ui = ui

	return newPlayer
}

func (h *HumanPlayer) Move(board *board.TttBoard) int {
	possibleMoves := possibleMoves(board)

	return h.ui.GetMove(possibleMoves) - humanReadabilityPadding
}

func possibleMoves(board *board.TttBoard) []int {
	var moves []int
	spaces := board.Spaces()

	for i, v := range spaces {
		if v == 0 {
			moves = append(moves, i+humanReadabilityPadding)
		}
	}
	return moves
}

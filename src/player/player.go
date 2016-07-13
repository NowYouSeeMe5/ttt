package player

import (
	"ttt/src/board"
)

type Player interface {
	Move(board *board.TttBoard) int
}

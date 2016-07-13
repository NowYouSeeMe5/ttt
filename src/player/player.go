package player

import (
	"github.com/NowYouSeeMe5/ttt/src/board"
)

type Player interface {
	Move(board *board.TttBoard) int
}

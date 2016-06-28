package game

import . "ttt/src/board"

type Game interface {
	Play()
}

type TttInput interface {
	HowManyPlayers() int
}

type TttOutput interface {
	PrintBoard(board Board)
	PrintEndGameMessage(winner int)
}

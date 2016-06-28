package game_test

import . "ttt/src/board"

type MockTttInput struct{}
type MockTttOutput struct{}

func (i MockTttInput) HowManyPlayers() int             { return 1 }
func (o MockTttOutput) PrintBoard(board Board)         {}
func (o MockTttOutput) PrintEndGameMessage(winner int) {}

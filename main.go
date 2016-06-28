package main

import (
	"ttt/src/board"
	"ttt/src/evaluator"
	"ttt/src/game"
	"ttt/src/io"
)

func main() {
	board := new(board.TttBoard)
	evaluator := new(evaluator.TttEvaluator)
	input := io.NewTttInput()
	output := new(io.TttOutput)

	game.NewTttGame(input, output, board, evaluator).Play()
}

package main

import (
	"ttt/src/game"
	"ttt/src/ui"
)

func main() {
	io := new(ui.ConsoleIO)

	game.NewTttGame(io).Play()
}

package main

import (
	"github.com/NowYouSeeMe5/ttt/src/game"
	"github.com/NowYouSeeMe5/ttt/src/ui"
)

func main() {
	io := new(ui.ConsoleIO)

	game.NewTttGame(io).Play()
}

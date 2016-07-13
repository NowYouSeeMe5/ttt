package game

import (
	"github.com/NowYouSeeMe5/ttt/src/board"
	"github.com/NowYouSeeMe5/ttt/src/evaluator"
	"github.com/NowYouSeeMe5/ttt/src/player"
	"github.com/NowYouSeeMe5/ttt/src/ui"
)

type TttGame struct {
	Player1   player.Player
	Player2   player.Player
	Board     *board.TttBoard
	evaluator *evaluator.TttEvaluator
	ui        *ui.TttUI
}

var (
	validBoardSizes      = []int{3, 4, 5, 6}
	validNumberOfPlayers = []int{0, 1, 2}
)

func NewTttGame(io ui.IO) *TttGame {
	TttGame := new(TttGame)

	TttGame.ui = ui.NewTttUI(io)
	TttGame.evaluator = new(evaluator.TttEvaluator)

	return TttGame
}

func (t *TttGame) Play() {
	t.Setup()
	t.PlayGame()
	t.End()
}

func (t *TttGame) Setup() {
	boardSize := t.ui.GetBoardSize(validBoardSizes)
	players := t.ui.GetNumberOfPlayers(validNumberOfPlayers)

	t.Board = board.NewTttBoard(boardSize)
	t.setPlayers(players)
}

func (t *TttGame) setPlayers(players int) {
	aiPlayer := player.NewAiPlayer()
	humanPlayer := player.NewHumanPlayer(t.ui)

	if players == 0 {
		t.Player1, t.Player2 = aiPlayer, aiPlayer
	} else if players == 1 {
		t.Player1, t.Player2 = humanPlayer, aiPlayer
	} else {
		t.Player1, t.Player2 = humanPlayer, humanPlayer
	}
}

func (t *TttGame) PlayGame() {
	board := t.Board

	for !t.evaluator.IsOver(board) {
		t.ui.PrintBoard(board)

		whoseTurn := t.evaluator.WhoseTurn(board)
		currentPlayer := t.getPlayer(whoseTurn)
		move := currentPlayer.Move(board)

		board.SetSpace(whoseTurn, move)
	}
}

func (t TttGame) getPlayer(whoseTurn int) player.Player {
	if whoseTurn == 1 {
		return t.Player1
	} else {
		return t.Player2
	}
}

func (t *TttGame) End() {
	t.ui.PrintBoard(t.Board)

	winner := t.evaluator.Winner(t.Board)
	t.ui.PrintEndGameMessage(winner)
}

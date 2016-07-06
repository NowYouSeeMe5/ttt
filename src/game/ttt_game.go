package game

import (
	"ttt/src/board"
	"ttt/src/evaluator"
	"ttt/src/io"
	"ttt/src/player"
)

const (
	Setup    = "setup"
	Playing  = "playing"
	End      = "end"
	Finished = "finished"
)

type GameSetup func()
type GamePlayer func()
type GameEnd func()

type TttGame struct {
	Board   board.Board
	Player1 player.Player
	Player2 player.Player

	CurrentState string

	Evaluator evaluator.Evaluator
	Input     io.GameInput
	Output    io.GameOutput

	gameSetup  GameSetup
	gamePlayer GamePlayer
	gameEnd    GameEnd
}

func NewTttGame(input io.GameInput, output io.GameOutput, board board.Board, evaluator evaluator.Evaluator) *TttGame {
	TttGame := new(TttGame)

	TttGame.Input = input
	TttGame.Output = output
	TttGame.Board = board
	TttGame.Evaluator = evaluator

	TttGame.gameSetup = TttGame.Setup
	TttGame.gamePlayer = TttGame.PlayGame
	TttGame.gameEnd = TttGame.End

	TttGame.CurrentState = Setup
	return TttGame
}

func (t *TttGame) Play() {

	for t.CurrentState != Finished {
		switch t.CurrentState {
		case Setup:
			t.CurrentState = Playing
			t.gameSetup()
		case Playing:
			t.CurrentState = End
			t.gamePlayer()
		case End:
			t.CurrentState = Finished
			t.gameEnd()
		}
	}
}

func (t *TttGame) Setup() {
	boardSize := t.Input.BoardSize()
	players := t.Input.HowManyPlayers()

	t.SetBoardSize(boardSize)
	t.SetPlayers(players)
}

func (t *TttGame) PlayGame() {
	board := t.Board
	for !t.IsOver(board) {
		t.printBoard()

		whoseTurn := t.Board.WhoseTurn()
		currentPlayer := t.GetPlayer(whoseTurn)
		move := currentPlayer.Move(t.Board)

		t.MakeMove(whoseTurn, move)
	}
}

func (t *TttGame) End() {
	t.printBoard()
	t.printWinner()
}

func (t TttGame) printBoard() {
	spaces := t.Board.Spaces()
	t.Output.PrintBoard(spaces)
}

func (t TttGame) printWinner() {
	spaces := t.Board.Spaces()
	winner := t.Evaluator.Winner(spaces)
	t.Output.PrintEndGameMessage(winner)
}

func (t *TttGame) SetBoardSize(boardSize int) {
	t.Board.ResetBoard(boardSize)
}

func (t *TttGame) SetPlayers(players int) {
	if players == 0 {
		t.Player1 = player.NewAiPlayer(t.Evaluator)
		t.Player2 = player.NewAiPlayer(t.Evaluator)
	} else if players == 1 {
		t.Player1 = new(player.HumanPlayer)
		t.Player2 = player.NewAiPlayer(t.Evaluator)
	} else {
		t.Player1 = new(player.HumanPlayer)
		t.Player2 = new(player.HumanPlayer)
	}
}

func (t TttGame) IsOver(board board.Board) bool {
	spaces := board.Spaces()

	winner := t.Evaluator.Winner(spaces)
	currentDepth := t.Board.CurrentDepth()

	return winner != 0 || currentDepth == 0
}

func (t TttGame) GetPlayer(whoseTurn int) player.Player {
	if whoseTurn == 1 {
		return t.Player1
	} else {
		return t.Player2
	}
}

func (t *TttGame) MakeMove(currentPlayer int, position int) {
	t.Board.MakeMove(currentPlayer, position)
}

func (t *TttGame) InjectSetup(gameSetup GameSetup) {
	t.gameSetup = gameSetup
}

func (t *TttGame) InjectPlayer(gamePlayer GamePlayer) {
	t.gamePlayer = gamePlayer
}

func (t *TttGame) InjectEnd(gameEnd GameEnd) {
	t.gameEnd = gameEnd
}

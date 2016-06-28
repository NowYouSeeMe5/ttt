package game

import (
	. "ttt/src/board"
	. "ttt/src/evaluator"
	. "ttt/src/player"
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
	Board   Board
	Player1 Player
	Player2 Player

	CurrentState string

	Evaluator Evaluator
	Input     TttInput
	Output    TttOutput

	gameSetup  GameSetup
	gamePlayer GamePlayer
	gameEnd    GameEnd
}

func NewTttGame(input TttInput, output TttOutput, board Board, evaluator Evaluator) *TttGame {
	TttGame := new(TttGame)

	TttGame.Input = input
	TttGame.Output = output
	TttGame.Board = board
	TttGame.Evaluator = evaluator

	TttGame.Board.ResetBoard(3)

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
	players := t.Input.HowManyPlayers()
	t.SetPlayers(players)
}

func (t *TttGame) SetPlayers(players int) {
	if players == 0 {
		t.Player1 = new(AiPlayer)
		t.Player2 = new(AiPlayer)
	} else if players == 1 {
		t.Player1 = new(HumanPlayer)
		t.Player2 = new(AiPlayer)
	} else {
		t.Player1 = new(HumanPlayer)
		t.Player2 = new(HumanPlayer)
	}
}

func (t *TttGame) PlayGame() {
	board := t.Board
	for !t.IsOver(board) {
		t.Output.PrintBoard(t.Board)

		whoseTurn := t.Board.WhoseTurn()
		currentPlayer := t.GetPlayer(whoseTurn)
		move := currentPlayer.Move(t.Board)

		t.MakeMove(whoseTurn, move)
	}
}

func (t TttGame) IsOver(board Board) bool {
	spaces := board.Spaces()

	winner := t.Evaluator.Winner(spaces)
	currentDepth := t.Board.CurrentDepth()

	return winner != 0 || currentDepth == 0
}

func (t TttGame) GetPlayer(whoseTurn int) Player {
	if whoseTurn == 1 {
		return t.Player1
	} else {
		return t.Player2
	}
}

func (t *TttGame) MakeMove(currentPlayer int, position int) {
	t.Board.MakeMove(currentPlayer, position)
}

func (t *TttGame) End() {
	t.Output.PrintBoard(t.Board)

	spaces := t.Board.Spaces()
	winner := t.Evaluator.Winner(spaces)

	t.Output.PrintEndGameMessage(winner)
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

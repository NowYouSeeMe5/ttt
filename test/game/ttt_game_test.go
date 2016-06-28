package game_test

import (
	"reflect"
	. "ttt/src/board"
	. "ttt/src/evaluator"
	. "ttt/src/game"
	. "ttt/src/player"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var mockCurrentState = Finished
var currentGameState = Setup

var input = new(MockTttInput)
var output = new(MockTttOutput)
var board = new(TttBoard)
var evaluator = new(TttEvaluator)

var tttGame = NewTttGame(input, output, board, evaluator)

func mockSetup() {
	currentGameState = tttGame.CurrentState
	mockCurrentState = Setup
	tttGame.CurrentState = Finished
}

func mockPlayGame() {
	currentGameState = tttGame.CurrentState
	mockCurrentState = Playing
	tttGame.CurrentState = Finished
}

func mockEndGame() {
	currentGameState = tttGame.CurrentState
	mockCurrentState = End
	tttGame.CurrentState = Finished
}

func startPlay(startingState string) {
	tttGame.CurrentState = startingState
	tttGame.Play()
}

var _ = Describe("TttGame", func() {

	tttGame.InjectSetup(mockSetup)
	tttGame.InjectPlayer(mockPlayGame)
	tttGame.InjectEnd(mockEndGame)

	Describe("Play", func() {
		It("Uses the gameSetup when the current state is Setup", func() {
			startPlay(Setup)
			Expect(Setup).To(Equal(mockCurrentState))
		})

		It("Always sets the current state of the game to Playing when the state is Setup", func() {
			startPlay(Setup)
			Expect(Playing).To(Equal(currentGameState))
		})

		It("Uses the gamePlayer when the current state is Playing", func() {
			startPlay(Playing)
			Expect(Playing).To(Equal(mockCurrentState))
		})

		It("Always sets the current state of the game to End when the state is Playing", func() {
			startPlay(Playing)
			Expect(End).To(Equal(currentGameState))
		})

		It("Uses gameEnd when the current state is End", func() {
			startPlay(End)
			Expect(End).To(Equal(mockCurrentState))
		})

		It("Always sets the current state of the game to Finished when the state is End", func() {
			startPlay(End)
			Expect(Finished).To(Equal(currentGameState))
		})
	})

	Describe("Setup", func() {

		humanType := reflect.TypeOf(new(HumanPlayer)).String()
		aiType := reflect.TypeOf(new(AiPlayer)).String()

		It("Sets both players to ai when the players count is 0", func() {
			tttGame.SetPlayers(0)

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(aiType).To(Equal(player1Type))
			Expect(aiType).To(Equal(player2Type))
		})

		It("Sets player 1 to ai and player 2 to human when the players count is 1", func() {
			tttGame.SetPlayers(1)

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(humanType).To(Equal(player1Type))
			Expect(aiType).To(Equal(player2Type))
		})

		It("Sets player 1 and player 2 to human when the players count is 2", func() {
			tttGame.SetPlayers(2)

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(humanType).To(Equal(player1Type))
			Expect(humanType).To(Equal(player2Type))
		})
	})

	Describe("Is over", func() {
		board := tttGame.Board

		It("Is not over if there are spaces left or if nobody has won", func() {
			spaces := []int{0, 1, -1, 1, -1, 1, 0, -1, 1}
			board.SetSpaces(spaces)
			Expect(false).To(Equal(tttGame.IsOver(board)))
		})

		It("Is over if there are no spaces left", func() {
			spaces := []int{1, -1, 1, -1, 1, -1, -1, 1, -1}
			board.SetSpaces(spaces)
			Expect(true).To(Equal(tttGame.IsOver(board)))
		})

		It("Is over if someone has won", func() {
			spaces := []int{1, 1, 1, -1, -1, 0, 0, 0, 0}
			board.SetSpaces(spaces)
			Expect(true).To(Equal(tttGame.IsOver(board)))
		})
	})

	Describe("Get move", func() {
		tttGame.SetPlayers(1)

		It("uses the player stored in Player1 if the current player equals 1", func() {
			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player := tttGame.GetPlayer(1)
			playerType := reflect.TypeOf(player).String()

			Expect(player1Type).To(Equal(playerType))
		})

		It("uses the player stored in Player2 if the current player equals 2", func() {
			player2Type := reflect.TypeOf(tttGame.Player2).String()
			player := tttGame.GetPlayer(2)
			playerType := reflect.TypeOf(player).String()

			Expect(player2Type).To(Equal(playerType))
		})
	})

	Describe("Make move", func() {
		It("Mutates the game board when given a player and position", func() {
			tttGame.Board.ResetBoard(3)
			tttGame.MakeMove(1, 1)

			expectedBoard := []int{0, 1, 0, 0, 0, 0, 0, 0, 0}
			Expect(expectedBoard).To(Equal(tttGame.Board.Spaces()))
		})
	})
})

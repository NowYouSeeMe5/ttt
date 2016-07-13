package game_test

import (
	"reflect"

	. "ttt/src/board"
	. "ttt/src/evaluator"
	. "ttt/src/game"
	. "ttt/src/player"
	. "ttt/src/ui"
	. "ttt/src/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var mockIO = new(MockGameIO)
var ui = NewTttUI(mockIO)
var board = NewTttBoard(3)

var tttGame = NewTttGame(mockIO)

var _ = Describe("TttGame", func() {

	humanType := reflect.TypeOf(new(HumanPlayer)).String()
	aiType := reflect.TypeOf(new(AiPlayer)).String()

	Describe("Set up", func() {

		It("Sets both players to ai players when the number of players input is 0", func() {
			mockIO.BoardSize = "3"
			mockIO.HowManyPlayers = "0"

			tttGame.Setup()

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(aiType).To(Equal(player1Type))
			Expect(aiType).To(Equal(player2Type))
		})

		It("Sets player 1 to ai and player 2 to human when the number of players input is 1", func() {
			mockIO.HowManyPlayers = "1"

			tttGame.Setup()

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(humanType).To(Equal(player1Type))
			Expect(aiType).To(Equal(player2Type))
		})

		It("Sets player 1 and player 2 to human when the number of players input is 2", func() {
			mockIO.HowManyPlayers = "2"

			tttGame.Setup()

			player1Type := reflect.TypeOf(tttGame.Player1).String()
			player2Type := reflect.TypeOf(tttGame.Player2).String()

			Expect(humanType).To(Equal(player1Type))
			Expect(humanType).To(Equal(player2Type))
		})

		It("Sets the board size when the user inputs 4", func() {
			mockIO.BoardSize = "4"

			tttGame.Setup()

			Expect(4).To(Equal(tttGame.Board.Size()))
		})
	})

	Describe("Play game", func() {
		It("Plays through a game when both players are ai", func() {
			mockIO.BoardSize = "3"
			mockIO.HowManyPlayers = "0"

			tttGame.Play()

			Expect(0).To(Equal(tttGame.Board.CurrentDepth()))
		})
	})

	Describe("End", func() {
		It("Prints the tie game message when r", func() {
			tttGame.End()

			Expect(TieGame).To(Equal(mockIO.PrintedString))
		})
	})
})

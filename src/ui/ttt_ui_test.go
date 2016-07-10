package ui

import (
	"fmt"
	. "ttt/src/board"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mockIO struct {
	printedString string
}

func (m *mockIO) Print(message string) {
	m.printedString = message
}

func (m mockIO) GetInput() string {
	return "1"
}

func setSpaces(spaces []int, board *TttBoard) {
	for i, v := range spaces {
		board.SetSpace(v, i)
	}
}

var _ = Describe("TttUI", func() {

	board := NewTttBoard(3)
	mockIO := new(mockIO)
	tttUI := NewTttUI(mockIO)

	Describe("Print board", func() {
		It("Prints an empty board when given an empty 3X3 board", func() {
			empty3X3Board := "_|_|_\n_|_|_\n_|_|_\n"

			tttUI.PrintBoard(board)

			Expect(empty3X3Board).To(Equal(mockIO.printedString))
		})

		It("Prints an empty board when given an empty 4X4 board", func() {
			empty4X4Board := "_|_|_|_\n_|_|_|_\n_|_|_|_\n_|_|_|_\n"

			board = NewTttBoard(4)
			tttUI.PrintBoard(board)

			Expect(empty4X4Board).To(Equal(mockIO.printedString))
		})

		It("Renders the expected board when given a non empty 3X3 board", func() {
			filled3X3Board := "X|O|X\nO|X|O\nO|X|O\n"
			spaces := []int{1, -1, 1, -1, 1, -1, -1, 1, -1}

			board = NewTttBoard(3)
			setSpaces(spaces, board)
			tttUI.PrintBoard(board)

			Expect(filled3X3Board).To(Equal(mockIO.printedString))
		})

		It("Renders the expected board when given a non empty 4X4 board", func() {
			filled4X4Board := "X|_|X|_\n_|O|_|O\nX|_|X|_\n_|O|_|_\n"
			spaces := []int{1, 0, 1, 0, 0, -1, 0, -1, 1, 0, 1, 0, 0, -1, 0, 0}

			board = NewTttBoard(4)
			setSpaces(spaces, board)
			tttUI.PrintBoard(board)

			Expect(filled4X4Board).To(Equal(mockIO.printedString))
		})
	})

	Describe("Print end game message", func() {
		It("Prints tie game message when there is no winner", func() {
			tttUI.PrintEndGameMessage(0)
			Expect(TieGame).To(Equal(mockIO.printedString))
		})

		It("Prints the player 1 winner message when player 1 wins", func() {
			winner := 1
			tttUI.PrintEndGameMessage(winner)
			expectedMessage := fmt.Sprintf(PlayerVictory, getPiece(winner))

			Expect(expectedMessage).To(Equal(mockIO.printedString))
		})

		It("Prints the player 2 winner message when player 2 wins", func() {
			winner := 2
			tttUI.PrintEndGameMessage(winner)
			expectedMessage := fmt.Sprintf(PlayerVictory, getPiece(winner))

			Expect(expectedMessage).To(Equal(mockIO.printedString))
		})
	})

	Describe("Get move", func() {
		validInput := []int{1, 2, 3}

		It("Returns the move typed in if it is in the valid input", func() {
			Expect(1).To(Equal(tttUI.GetMove(validInput)))
		})

		It("Asks for the move based on the valid available moves", func() {
			expectedQuestion := fmt.Sprintf(GetMove, validInput)

			Expect(expectedQuestion).To(Equal(mockIO.printedString))
		})
	})

	Describe("Get board size", func() {
		validInput := []int{1, 2, 3}

		It("Returns the input if it is in the validInput", func() {
			Expect(1).To(Equal(tttUI.GetBoardSize(validInput)))
		})

		It("Asks for the size based on the valid available moves", func() {
			expectedQuestion := fmt.Sprintf(BoardSizeQuestion, validInput)

			Expect(expectedQuestion).To(Equal(mockIO.printedString))
		})
	})

	Describe("Get number of players", func() {
		It("Returns the input if it is in the validInput", func() {
			Expect(1).To(Equal(tttUI.GetNumberOfPlayers()))
		})

		It("Asks for the number of players", func() {
			expectedQuestion := fmt.Sprintf(HowManyPlayersQuestion, []int{0, 1, 2})

			Expect(expectedQuestion).To(Equal(mockIO.printedString))
		})
	})
})

package ui

import (
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
	return ""
}

func setSpaces(spaces []int, board *TttBoard) {
	for i, v := range spaces {
		board.SetSpace(v, i)
	}
}

var _ = Describe("ConsoleUi", func() {

	board := NewTttBoard(3)
	consoleUI := NewConsoleUI(new(mockIO))

	Describe("Print board", func() {
		It("Prints an empty board when given an empty 3X3 board", func() {
			empty3X3Board := "_|_|_\n_|_|_\n_|_|_\n"

			consoleUI.PrintBoard(board)

			Expect(empty3X3Board).To(Equal(consoleUI.io.printedString))
		})

		It("Prints an empty board when given an empty 4X4 board", func() {
			empty4X4Board := "_|_|_|_\n_|_|_|_\n_|_|_|_\n_|_|_|_\n"

			board = NewTttBoard(4)
			consoleUI.PrintBoard(board)

			Expect(empty4X4Board).To(Equal(consoleUI.io.printedString))
		})

		It("Renders the expected board when given a non empty 3X3 board", func() {
			filled3X3Board := "X|O|X\nO|X|O\nO|X|O\n"
			spaces := []int{1, -1, 1, -1, 1, -1, -1, 1, -1}

			setSpaces(spaces, board)
			consoleUI.PrintBoard(board)

			Expect(filled3X3Board).To(Equal(consoleUI.io.printedString))
		})

		It("Renders the expected board when given a non empty 4X4 board", func() {
			filled4X4Board := "X|_|X|_\n_|O|_|O\nX|_|X|_\n_|O|_|_\n"
			spaces := []int{1, 0, 1, 0, 0, -1, 0, -1, 1, 0, 1, 0, 0, -1, 0, 0}

			setSpaces(spaces, board)
			consoleUI.PrintBoard(board)

			Expect(filled4X4Board).To(Equal(consoleUI.io.printedString))
		})
	})
})

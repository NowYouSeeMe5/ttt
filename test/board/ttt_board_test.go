package board_test

import (
	. "ttt/src/board"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TttBoard", func() {

	n := 3
	board := new(TttBoard)
	board.ResetBoard(n)
	spaces := board.GetSpaces()

	Describe("New Board", func() {
		It("Should be an array of size n-squared when given n", func() {
			Expect(n * n).To(Equal(len(spaces)))
		})

		It("Should be made up entirely of empty space", func() {
			for i := range spaces {
				Expect(0).To(Equal(spaces[i]))
			}
		})
	})

	Describe("Current depth", func() {
		It("returns the remaining number of zeroes left on a board", func() {
			board.SetSpaces([]int{0, 1, 1, 1, 2, 3, 2, 0, 1})
			Expect(2).To(Equal(board.CurrentDepth()))
		})
	})

	Describe("Make move", func() {
		It("returns a board with the player's number in the appropriate space", func() {
			board.ResetBoard(3)
			board.MakeMove(1, 3)
			spaces := board.GetSpaces()
			Expect(1).To(Equal(spaces[3]))
		})
	})
})
package board

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func setSpaces(spaces []int, board *TttBoard) {
	for i, v := range spaces {
		board.SetSpace(v, i)
	}
}

var _ = Describe("TttBoard", func() {

	n := 3

	var (
		board  *TttBoard
		spaces []int
	)

	BeforeEach(func() {
		board = NewTttBoard(n)
		spaces = board.Spaces()
	})

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

	Describe("Set Space", func() {
		It("sets a space with the given marker at the given position", func() {
			newSpaces := []int{0, 0, 1, 0, 0, -1, 1, 1, -1}

			setSpaces(newSpaces, board)
			board.SetSpace(1, 3)

			Expect(1).To(Equal(spaces[3]))
		})
	})

	Describe("Size", func() {
		It("returns the 1 dimensional size of a 3X3 board", func() {
			Expect(n).To(Equal(board.Size()))
		})

		It("returns the 1 dimensional size of a 4X4 board", func() {
			board = NewTttBoard(4)

			Expect(4).To(Equal(board.Size()))
		})
	})

	Describe("Whose turn", func() {
		It("returns 1 when the board adds up to 0", func() {
			newSpaces := []int{0, 0, 1, -1, 0, -1, 1, 1, -1}
			setSpaces(newSpaces, board)
			Expect(1).To(Equal(board.WhoseTurn()))
		})

		It("returns -1 when the board adds up to 1", func() {
			newSpaces := []int{0, 0, 1, 1, 0, 1, 0, -1, 0}
			setSpaces(newSpaces, board)
			Expect(-1).To(Equal(board.WhoseTurn()))
		})
	})

	Describe("Possible moves", func() {
		It("returns a slice containing the possible moves on a board", func() {
			newSpaces := []int{0, 0, 0, 1, 0, -1, 0, -1, 0}
			setSpaces(newSpaces, board)
			Expect([]int{1, 2, 3, 5, 7, 9}).To(Equal(board.PossibleMoves()))
		})
	})
})

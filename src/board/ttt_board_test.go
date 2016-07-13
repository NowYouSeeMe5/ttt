package board

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

	Describe("Copy", func() {
		It("returns a new TttBoard with the same spaces and different addresses", func() {
			board = NewTttBoard(3)
			board.SetSpace(1, 8)

			copy := board.Copy()

			Expect(board.Spaces()).To(Equal(copy.Spaces()))
			Expect(&board).To(Not(Equal(copy)))
		})
	})
})

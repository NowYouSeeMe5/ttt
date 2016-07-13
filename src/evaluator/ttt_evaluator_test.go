package evaluator

import (
	. "ttt/src/board"
	. "ttt/src/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TttEvaluator", func() {

	board := NewTttBoard(3)
	tttEvaluator := new(TttEvaluator)
	spaces := []int{}

	Describe("Is over", func() {
		board := NewTttBoard(3)

		It("Is not over if there are spaces left or if nobody has won", func() {
			spaces := []int{0, 1, -1, 1, -1, 1, 0, -1, 1}

			SetSpaces(spaces, board)

			Expect(false).To(Equal(tttEvaluator.IsOver(board)))
		})

		It("Is over if there are no spaces left", func() {
			spaces := []int{1, -1, 1, -1, 1, -1, -1, 1, -1}

			SetSpaces(spaces, board)

			Expect(true).To(Equal(tttEvaluator.IsOver(board)))
		})

		It("Is over if someone has won", func() {
			spaces := []int{1, 1, 1, -1, -1, 0, 0, 0, 0}

			SetSpaces(spaces, board)

			Expect(true).To(Equal(tttEvaluator.IsOver(board)))
		})
	})

	Describe("Check for winner", func() {

		It("returns a 0 if there is no winner", func() {
			Expect(0).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns a 0 if there is no winner", func() {
			spaces = []int{1, 1, -1, 1, -1, 0, 0, 0, -1}

			SetSpaces(spaces, board)

			Expect(0).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns a 0 if there is a tie", func() {
			spaces = []int{1, -1, 1, -1, 1, -1, -1, 1, -1}

			SetSpaces(spaces, board)

			Expect(0).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns -1 if player 2 has won vertically", func() {
			spaces = []int{-1, 0, 0, -1, 0, 0, -1, 0, 0}

			SetSpaces(spaces, board)

			Expect(-1).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns 1 if player 1 has won horizontally", func() {
			spaces = []int{1, 1, 1, 0, 0, 0, 0, 0, 0}

			SetSpaces(spaces, board)

			Expect(1).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns 1 on a 4X4 board where player 1 has won horizontally", func() {
			board = NewTttBoard(4)
			spaces = []int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

			SetSpaces(spaces, board)

			Expect(1).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns -1 on a 4X4 board where player 2 has won vertically", func() {
			spaces = []int{-1, 0, 0, 0, -1, 0, 0, 0, -1, 0, 0, 0, -1, 0, 0, 0}

			SetSpaces(spaces, board)

			Expect(-1).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns 1 on a 4X4 board where player 1 has won diagonally left to right", func() {
			spaces = []int{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}

			SetSpaces(spaces, board)

			Expect(1).To(Equal(tttEvaluator.Winner(board)))
		})

		It("returns -1 on a 3X3 board where player 2 has won diagonally right to left", func() {
			board = NewTttBoard(3)
			spaces = []int{0, 0, -1, 0, -1, 0, -1, 0, 0}

			SetSpaces(spaces, board)

			Expect(-1).To(Equal(tttEvaluator.Winner(board)))
		})
	})

	Describe("Whose turn", func() {
		It("returns 1 when the board adds up to 0", func() {
			newSpaces := []int{0, 0, 1, -1, 0, -1, 1, 1, -1}

			SetSpaces(newSpaces, board)

			Expect(1).To(Equal(tttEvaluator.WhoseTurn(board)))
		})

		It("returns -1 when the board adds up to 1", func() {
			newSpaces := []int{0, 0, 1, 1, 0, 1, 0, -1, 0}

			SetSpaces(newSpaces, board)

			Expect(-1).To(Equal(tttEvaluator.WhoseTurn(board)))
		})
	})
})

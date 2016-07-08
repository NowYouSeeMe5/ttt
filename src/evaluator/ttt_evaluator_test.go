package evaluator_test

import (
	. "ttt/src/evaluator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TttEvaluator", func() {

	Describe("Check for winner", func() {

		tttEvaluator := new(TttEvaluator)
		spaces := []int{E, E, E, E, E, E, E, E, E}

		It("returns a 0 if there is no winner", func() {
			Expect(E).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns a 0 if there is no winner", func() {
			spaces = []int{P1, P1, P2, P1, P2, 0, 0, 0, P2}
			Expect(E).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns a 0 if there is a tie", func() {
			spaces = []int{P1, P2, P1, P2, P1, P2, P2, P1, P2}
			Expect(E).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P2 if Player 2 has won vertically", func() {
			spaces = []int{P2, E, E, P2, E, E, P2, E, E}
			Expect(P2).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P1 if Player 1 has won horizontally", func() {
			spaces = []int{P1, P1, P1, E, E, E, E, E, E}
			Expect(P1).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P1 on a 4X4 board where P1 has won horizontally", func() {
			spaces = []int{P1, P1, P1, P1, E, E, E, E, E, E, E, E, E, E, E, E}
			Expect(P1).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P2 on a 4X4 board where P2 has won vertically", func() {
			spaces = []int{P2, E, E, E, P2, E, E, E, P2, E, E, E, P2, E, E, E}
			Expect(P2).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P1 on a 4X4 board where P1 has won diagonally left to right", func() {
			spaces = []int{P1, E, E, E, E, P1, E, E, E, E, P1, E, E, E, E, P1}
			Expect(P1).To(Equal(tttEvaluator.Winner(spaces)))
		})

		It("returns P2 on a 3X3 board where P2 has won diagonally right to left", func() {
			spaces = []int{E, E, P2, E, P2, E, P2, E, E}
			Expect(P2).To(Equal(tttEvaluator.Winner(spaces)))
		})
	})
})

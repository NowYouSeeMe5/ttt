package io_test

import (
	"fmt"

	. "ttt/src/io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConsoleUi", func() {

	tttOutput := new(TttOutput)

	Describe("Render board", func() {

		It("Renders the expected board when given an empty 3X3 board", func() {
			empty3X3Board := "_|_|_\n_|_|_\n_|_|_\n"
			spaces := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

			Expect(empty3X3Board).To(Equal(tttOutput.RenderBoard(spaces)))
		})

		It("Renders the expected board when given an empty 4X4 board", func() {
			empty4X4Board := "_|_|_|_\n_|_|_|_\n_|_|_|_\n_|_|_|_\n"
			spaces := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

			Expect(empty4X4Board).To(Equal(tttOutput.RenderBoard(spaces)))
		})

		It("Renders the expected board when given a non empty 3X3 board", func() {
			filled3X3Board := "X|O|X\nO|X|O\nO|X|O\n"
			spaces := []int{1, -1, 1, -1, 1, -1, -1, 1, -1}

			Expect(filled3X3Board).To(Equal(tttOutput.RenderBoard(spaces)))
		})

		It("Renders the expected board when given a non empty 4X4 board", func() {
			filled4X4Board := "X|_|X|_\n_|O|_|O\nX|_|X|_\n_|O|_|_\n"
			spaces := []int{1, 0, 1, 0, 0, -1, 0, -1, 1, 0, 1, 0, 0, -1, 0, 0}

			Expect(filled4X4Board).To(Equal(tttOutput.RenderBoard(spaces)))
		})
	})

	Describe("End game message", func() {
		It("Prints the victory message with X when player 1 has won", func() {
			expectedMessage := fmt.Sprintf(PlayerVictory, "X")
			Expect(expectedMessage).To(Equal(tttOutput.EndGameMessage(1)))
		})

		It("Prints the victory message with O when player 2 has won", func() {
			expectedMessage := fmt.Sprintf(PlayerVictory, "O")
			Expect(expectedMessage).To(Equal(tttOutput.EndGameMessage(-1)))
		})

		It("Prints the tie game message neither player has won", func() {
			Expect(TieGame).To(Equal(tttOutput.EndGameMessage(0)))
		})
	})
})

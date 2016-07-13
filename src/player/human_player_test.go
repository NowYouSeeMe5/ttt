package player

import (
	"fmt"

	. "ttt/src/board"
	. "ttt/src/ui"
	. "ttt/src/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HumanPlayer", func() {

	mockIO := new(MockIO)
	tttInput := NewTttUI(mockIO)
	player := NewHumanPlayer(tttInput)

	board := NewTttBoard(3)

	Describe("Move", func() {
		It("Subtracts one from the input returned from the human user", func() {
			Expect(0).To(Equal(player.Move(board)))
		})
	})

	Describe("Possible moves", func() {
		It("returns a slice containing the possible moves on a board", func() {
			newSpaces := []int{0, 0, 0, 1, 0, -1, 0, -1, 0}
			SetSpaces(newSpaces, board)

			expectedQuestion := fmt.Sprintf(GetMove, []int{1, 2, 3, 5, 7, 9})
			player.Move(board)

			Expect(expectedQuestion).To(Equal(mockIO.PrintedString))
		})
	})
})

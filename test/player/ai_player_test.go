package player_test

import (
	. "ttt/src/board"
	. "ttt/src/player"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiPlayer", func() {

	Describe("Move", func() {
		aiPlayer := new(AiPlayer)
		board := new(TttBoard)

		//It("", func() {
		//board.ResetBoard(3)
		//Expect(0).To(Equal(aiPlayer.Move(board)))
		//})

		It("goes in the middle space when it has the chance to win", func() {
			spaces := []int{-1, 0, 1, 1, 0, 0, 1, -1, -1}
			board.SetSpaces(spaces)
			Expect(4).To(Equal(aiPlayer.Move(board)))
		})

		It("makes a move", func() {
			spaces := []int{1, 1, -1, 1, 0, 0, 0, 0, -1}
			board.SetSpaces(spaces)
			Expect(5).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 4 when spaces 1 and 7 are taken", func() {
			spaces := []int{1, -1, 1, 0, -1, 0, 0, 0, 0}
			board.SetSpaces(spaces)
			Expect(7).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 10 when there is an imminent win", func() {
			spaces := []int{1, -1, 1, -1, -1, 1, 1, -1, 0, -1, 0, -1, 0, -1, 1, 1}
			board.SetSpaces(spaces)
			Expect(10).To(Equal(aiPlayer.Move(board)))
		})
	})
})

package player

import (
	. "ttt/src/board"
	. "ttt/src/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AiPlayer", func() {

	Describe("Move", func() {
		aiPlayer := NewAiPlayer()
		board := NewTttBoard(3)

		It("goes in space 4 when it has the chance to win", func() {
			spaces := []int{
				-1, 0, 1,
				1, 0, 0,
				1, -1, -1}
			SetSpaces(spaces, board)
			Expect(4).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 4 when spaces 1 and 7 are taken", func() {
			spaces := []int{
				1, -1, 1,
				0, 0, 0,
				0, -1, 0}
			SetSpaces(spaces, board)
			Expect(4).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 3 when there is an imminent win in a 4X4 game", func() {
			board = NewTttBoard(4)
			spaces := []int{
				1, 1, 1, 0,
				-1, -1, 0, 0,
				-1, 1, 0, 0,
				-1, -1, 1, 1}
			SetSpaces(spaces, board)
			Expect(3).To(Equal(aiPlayer.Move(board)))
		})

		It("wins on 20 in a 5X5 game", func() {
			board = NewTttBoard(5)
			spaces := []int{
				1, -1, 0, -1, -1,
				1, 1, -1, 0, -1,
				1, -1, 1, 0, 0,
				1, -1, 1, -1, -1,
				0, 1, 1, -1, 1}
			SetSpaces(spaces, board)
			Expect(20).To(Equal(aiPlayer.Move(board)))
		})
	})
})

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

		It("goes in the middle space when it has the chance to win", func() {
			spaces := []int{
				-1, 0, 1,
				1, 0, 0,
				1, -1, -1}
			board.SetSpaces(spaces)
			Expect(4).To(Equal(aiPlayer.Move(board)))
		})

		It("makes a move", func() {
			spaces := []int{
				1, 1, -1,
				1, 0, 0,
				0, 0, -1}
			board.SetSpaces(spaces)
			Expect(5).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 4 when spaces 1 and 7 are taken", func() {
			spaces := []int{
				1, -1, 1,
				0, -1, 0,
				0, 0, 0}
			board.SetSpaces(spaces)
			Expect(7).To(Equal(aiPlayer.Move(board)))
		})

		It("blocks on 3 when there is an imminent win in a 4X4 game", func() {
			spaces := []int{
				1, 1, 1, 0,
				-1, -1, 0, 0,
				-1, 1, 0, 0,
				-1, -1, 1, 1}
			board.SetSpaces(spaces)
			Expect(3).To(Equal(aiPlayer.Move(board)))
		})

		It("wins on 20 in a 5X5 game", func() {
			spaces := []int{
				1, -1, 0, -1, -1,
				1, 1, -1, 0, -1,
				1, -1, 1, 0, 0,
				1, -1, 1, -1, -1,
				0, 1, 1, -1, 1}
			board.SetSpaces(spaces)
			Expect(20).To(Equal(aiPlayer.Move(board)))
		})
	})
})

package player_test

import (
	"fmt"

	. "ttt/src/board"
	. "ttt/src/io"
	. "ttt/src/player"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func mockPrompter(message string) interface{} {
	return message
}

func mockInputGetter() int {
	return 2
}

func mockErrorHandler(expectedResults []int, result interface{}) bool {
	return true
}

var _ = Describe("HumanPlayer", func() {

	tttInput := NewTttInput()
	tttInput.InjectPrompter(mockPrompter)
	tttInput.InjectErrorHandler(mockErrorHandler)
	tttInput.InjectGetter(mockInputGetter)

	board := new(TttBoard)
	spaces := []int{0, 1, 1, 1, -1, 1, -1, 0, 1}
	board.SetSpaces(spaces)

	Describe("Form question", func() {
		It("Forms a question out of the size of a board", func() {
			expectedQuestion := fmt.Sprintf(Question, 9)
			Expect(expectedQuestion).To(Equal(FormQuestion(board)))
		})
	})

	Describe("Get input", func() {
		It("Subtracts one from the input returned from the human user", func() {
			Expect(1).To(Equal(GetInput("test", []int{}, tttInput)))
		})
	})
})

package io_test

import (
	. "ttt/src/io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errorCount = 0

func mockPrompter(message string) interface{} {
	return message
}

func mockInputGetter() int {
	return 2
}

func mockErrorHandler(expectedResults []int, result interface{}) bool {
	return true
}

func mockFailedErrorHandler(expectedResults []int, result interface{}) bool {
	if errorCount == 0 {
		errorCount += 1
		return false
	}

	return true
}

var _ = Describe("ConsoleInput", func() {

	tttInput := NewTttInput()
	tttInput.InjectPrompter(mockPrompter)
	tttInput.InjectErrorHandler(mockErrorHandler)
	tttInput.InjectGetter(mockInputGetter)

	Describe("Size board", func() {
		It("Uses the input getter to retrieve the size of the board to be played on", func() {
			Expect(2).To(Equal(tttInput.BoardSize()))
		})

		It("Retries to get the input when anything but the expected response is returned", func() {
			errorCount = 0
			tttInput.InjectErrorHandler(mockFailedErrorHandler)
			Expect(2).To(Equal(tttInput.BoardSize()))
			Expect(1).To(Equal(errorCount))
		})
	})

	Describe("How many players", func() {
		It("Uses the input getter to retrieve the number of players", func() {
			Expect(2).To(Equal(tttInput.HowManyPlayers()))
		})

		It("Retries to get the input when anything but the expected response is returned", func() {
			errorCount = 0
			tttInput.InjectErrorHandler(mockFailedErrorHandler)
			Expect(2).To(Equal(tttInput.HowManyPlayers()))
			Expect(1).To(Equal(errorCount))
		})
	})
})

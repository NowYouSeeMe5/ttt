package ui

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InputValidator", func() {

	validInput := []int{1, 2, 3}
	validator := new(TttValidator)

	Describe("Validate input", func() {
		It("Returns false if the input is not in the expected input", func() {
			wrongInput := 4
			Expect(false).To(Equal(validator.Validate(wrongInput, validInput)))
		})

		It("Returns true if the input is in the expected input", func() {
			correctInput := 3
			Expect(true).To(Equal(validator.Validate(correctInput, validInput)))
		})
	})
})

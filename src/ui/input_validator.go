package ui

type TttValidator struct{}

func (v TttValidator) Validate(input int, validInput []int) bool {
	for _, v := range validInput {
		if v == input {
			return true
		}
	}
	return false
}

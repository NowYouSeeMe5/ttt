package evaluator

import "math"

const (
	P1 = 1
	P2 = -1
	E  = 0
)

type TttEvaluator struct{}

func (t TttEvaluator) Winner(spaces []int) int {
	size := size(spaces)

	verticalWinner := verticalWinner(size, spaces)
	horizontalWinner := horizontalWinner(size, spaces)
	diagonalLeftRightWinner := diagonalLeftRightWinner(size, spaces)
	diagonalRightLeftWinner := diagonalRightLeftWinner(size, spaces)

	winners := []int{verticalWinner, horizontalWinner, diagonalLeftRightWinner, diagonalRightLeftWinner}

	return decideWinner(winners)
}

func decideWinner(winners []int) int {
	for _, v := range winners {
		if v != 0 {
			return v
		}
	}
	return 0
}

func verticalWinner(size int, spaces []int) int {
	winner := E

	for column := 0; column < size; column++ {
		verticalScore := 0
		lastIndexInColumn := column + (size * (size - 1))
		for row := column; row <= lastIndexInColumn; row += size {
			verticalScore += spaces[row]
		}

		winner = checkForWinner(size, verticalScore)
		if winner != 0 {
			return winner
		}
	}
	return winner
}

func horizontalWinner(size int, spaces []int) int {
	winner := E

	lastIndex := size * (size - 1)
	for row := 0; row <= lastIndex; row += size {
		rowScore := 0
		for column := row; column < row+size; column++ {
			rowScore += spaces[column]
		}

		winner = checkForWinner(size, rowScore)
		if winner != 0 {
			return winner
		}
	}
	return winner
}

func diagonalLeftRightWinner(size int, spaces []int) int {
	winner := E

	diagonalScore := 0
	for i := 0; i < size*size; i += (size + 1) {
		diagonalScore += spaces[i]
	}

	winner = checkForWinner(size, diagonalScore)

	return winner
}

func diagonalRightLeftWinner(size int, spaces []int) int {
	winner := E

	diagonalScore := 0
	for i := size - 1; i <= size*size-(size-1); i += (size - 1) {
		diagonalScore += spaces[i]
	}

	winner = checkForWinner(size, diagonalScore)

	return winner
}

func checkForWinner(size int, score int) int {
	winner := E

	if float64(size) == math.Abs(float64(score)) {
		winner = score / size
	}

	return winner
}

func size(spaces []int) int {
	squares := float64(len(spaces))
	return int(math.Sqrt(squares))
}

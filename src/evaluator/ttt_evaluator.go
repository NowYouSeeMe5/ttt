package evaluator

import (
	"math"
	"ttt/src/board"
)

type TttEvaluator struct{}

func (t TttEvaluator) IsOver(board *board.TttBoard) bool {
	winner := t.Winner(board)
	currentDepth := board.CurrentDepth()

	return winner != 0 || currentDepth == 0
}

func (t TttEvaluator) WhoseTurn(board *board.TttBoard) int {
	spaces := board.Spaces()

	total := 0

	for i := range spaces {
		total += spaces[i]
	}

	if total == 0 {
		return 1
	}
	return -1
}

func (t TttEvaluator) Winner(board *board.TttBoard) int {
	spaces := board.Spaces()
	size := board.Size()

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
	winner := 0

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
	winner := 0

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
	winner := 0

	diagonalScore := 0
	for i := 0; i < size*size; i += (size + 1) {
		diagonalScore += spaces[i]
	}

	winner = checkForWinner(size, diagonalScore)

	return winner
}

func diagonalRightLeftWinner(size int, spaces []int) int {
	winner := 0

	diagonalScore := 0
	for i := size - 1; i <= size*size-(size-1); i += (size - 1) {
		diagonalScore += spaces[i]
	}

	winner = checkForWinner(size, diagonalScore)

	return winner
}

func checkForWinner(size int, score int) int {
	winner := 0

	if float64(size) == math.Abs(float64(score)) {
		winner = score / size
	}

	return winner
}

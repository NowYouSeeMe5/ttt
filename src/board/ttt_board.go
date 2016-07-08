package board

import "math"

type TttBoard struct {
	spaces []int
}

func NewTttBoard(size int) *TttBoard {
	tttBoard := new(TttBoard)
	tttBoard.initializeBoard(size)

	return tttBoard
}

func (t *TttBoard) initializeBoard(size int) {
	var spaces = size * size
	t.spaces = make([]int, spaces)
}

func (t TttBoard) Size() int {
	squares := float64(len(t.spaces))
	return int(math.Sqrt(squares))
}

func (t *TttBoard) SetSpace(playerPosition int, moveSpace int) {
	t.spaces[moveSpace] = playerPosition
}

func (t TttBoard) CurrentDepth() int {
	var depth = 0
	spaces := t.spaces

	for i := range spaces {
		if spaces[i] == 0 {
			depth++
		}
	}
	return depth
}

func (t TttBoard) PossibleMoves() []int {
	var moves []int
	for i, v := range t.spaces {
		if v == 0 {
			moves = append(moves, i+1)
		}
	}
	return moves
}

func (t TttBoard) Spaces() []int {
	return t.spaces
}

func (t TttBoard) WhoseTurn() int {
	spaces := t.spaces

	total := 0

	for i := range spaces {
		total += spaces[i]
	}

	if total == 0 {
		return 1
	}
	return -1
}

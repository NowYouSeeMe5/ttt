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

func (t TttBoard) Spaces() []int {
	return t.spaces
}

func (t TttBoard) Copy() *TttBoard {
	spaces := make([]int, len(t.Spaces()))

	copy(spaces, t.Spaces())

	copy := &TttBoard{spaces}
	return copy
}

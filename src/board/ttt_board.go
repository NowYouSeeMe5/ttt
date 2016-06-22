package board

type TttBoard struct {
	spaces []int
}

func (t *TttBoard) MakeMove(playerPosition int, moveSpace int) {
	t.spaces[moveSpace] = playerPosition
}

func (t *TttBoard) ResetBoard(size int) {
	var spaces = size * size
	t.spaces = make([]int, spaces)
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

func (t TttBoard) GetSpaces() []int {
	return t.spaces
}

func (t *TttBoard) SetSpaces(spaces []int) {
	t.spaces = spaces
}

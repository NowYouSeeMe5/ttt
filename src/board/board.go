package board

type Board interface {
	MakeMove(playerPosition int, moveSpace int)
	ResetBoard(size int)
	CurrentDepth() int
	GetSpaces() []int
}

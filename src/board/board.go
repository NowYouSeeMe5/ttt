package board

type Board interface {
	CurrentDepth() int
	Spaces() []int
	SetSpaces(spaces []int)
	WhoseTurn() int
	PossibleMoves() []int
	MakeMove(playerPosition int, moveSpace int)
	ResetBoard(size int)
}

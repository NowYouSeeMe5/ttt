package io

type Input interface {
	AskQuestion(question string, expectedInput []int) int
}

type GameInput interface {
	BoardSize() int
	HowManyPlayers() int
}

type GameOutput interface {
	PrintBoard(spaces []int)
	PrintEndGameMessage(winner int)
}

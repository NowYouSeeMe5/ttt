package io

type Input interface {
	AskQuestion(question string, expectedInput []int) int
}

type Output interface {
	RenderBoard(spaces []int) string
}

package util

import (
	"strings"

	. "ttt/src/board"
)

type MockIO struct {
	PrintedString string
}

func (m *MockIO) Print(message string) {
	m.PrintedString = message
}

func (m MockIO) GetInput() string {
	return "1"
}

type MockGameIO struct {
	PrintedString  string
	BoardSize      string
	HowManyPlayers string
	GetMove        string
}

func (m *MockGameIO) Print(message string) {
	m.PrintedString = message
}

func (m MockGameIO) GetInput() string {
	message := m.PrintedString

	if strings.Contains(message, "Choose a board size") {
		return m.BoardSize
	} else if strings.Contains(message, "How many human players") {
		return m.HowManyPlayers
	} else if strings.Contains(message, "Please select a move") {
		return m.GetMove
	}
	return ""
}

func SetSpaces(spaces []int, board *TttBoard) {
	for i, v := range spaces {
		board.SetSpace(v, i)
	}
}

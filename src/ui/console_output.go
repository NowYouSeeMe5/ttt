package ui

//import (
//"bytes"
//"fmt"
//"math"
//"ttt/src/board"
//)

//type TttOutput struct{}

//const (
//BoardSizeQuestion      = "What size board would you like to play on? Enter 3, 4, or 5"
//HowManyPlayersQuestion = "How many human players are you? Enter 0, 1, or 2"
//ErrorMessage           = "Please enter the appropriate information"
//PlayerVictory          = "Well done %s, congrats on your victory!!!"
//TieGame                = "That's a nice Tie!!!"

//leftSpace = "%s"
//space     = "|%s"
//)

//var pieces = map[int]string{
//0:  "_",
//1:  "X",
//-1: "O",
//}

//func (t TttOutput) PrintBoard(board board.Board) {
//boardString := renderLines(spaces)
//t.Print(boardString)
//}

//func renderLines(spaces []int) string {
//size := size(spaces)
//var buffer bytes.Buffer

//for y := 0; y < size; y++ {
//currentY := y * size
//buffer.WriteString(formatSpace(leftSpace, spaces, currentY))
//for x := 1; x < size; x++ {
//buffer.WriteString(formatSpace(space, spaces, currentY+x))
//}
//buffer.WriteString("\n")
//}
//return buffer.String()
//}

//func formatSpace(space string, spaces []int, index int) string {
//return fmt.Sprintf(space, getPiece(spaces[index]))
//}

//func getPiece(squareState int) string {
//return pieces[squareState]
//}

//func (t TttOutput) PrintEndGameMessage(winner int) {
//endGameMessage := t.EndGameMessage(winner)
//t.Print(endGameMessage)
//}

//func (t TttOutput) EndGameMessage(winner int) string {
//if winner != 0 {
//return fmt.Sprintf(PlayerVictory, getPiece(winner))
//}
//return TieGame
//}

//func size(spaces []int) int {
//squares := float64(len(spaces))
//return int(math.Sqrt(squares))
//}

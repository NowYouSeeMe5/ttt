package player

import (
	"math"

	"github.com/NowYouSeeMe5/ttt/src/board"
	"github.com/NowYouSeeMe5/ttt/src/evaluator"
)

type AiPlayer struct {
	aiPlayer      int
	startingDepth int
	evaluator     *evaluator.TttEvaluator
}

var depthLimit = 8

func NewAiPlayer() *AiPlayer {
	player := new(AiPlayer)
	player.evaluator = new(evaluator.TttEvaluator)

	return player
}

func (p *AiPlayer) Move(board *board.TttBoard) int {
	p.aiPlayer = p.evaluator.WhoseTurn(board)

	spaces := board.Spaces()
	depth := board.CurrentDepth()
	p.startingDepth = depth
	alpha := -1
	beta := 1

	alphaBetaScores := make(map[int]int)

	for i, v := range spaces {
		if v == 0 {
			nextBoard := board.Copy()
			nextBoard.SetSpace(p.aiPlayer, i)
			otherPlayer := otherPlayer(p.aiPlayer)

			alphaBetaScores[i] = p.alphaBeta(nextBoard, depth-1, alpha, beta, otherPlayer)
		}
	}

	return bestMove(alphaBetaScores)
}

func (p *AiPlayer) alphaBeta(board *board.TttBoard, depth int, alpha int, beta int, currentPlayer int) int {
	spaces := board.Spaces()
	winner := p.evaluator.Winner(board)
	otherPlayer := otherPlayer(currentPlayer)

	if depth == 0 || winner != 0 || p.startingDepth-depth >= depthLimit {
		return p.score(winner)
	}

	if p.isMaximizingPlayer(currentPlayer) {
		score := -1000
		for i, v := range spaces {
			if v == 0 {
				nextBoard := board.Copy()
				nextBoard.SetSpace(currentPlayer, i)

				score = max(score, p.alphaBeta(nextBoard, depth-1, alpha, beta, otherPlayer))
				alpha = max(alpha, score)

				if alpha >= beta {
					return alpha
				}
			}
		}
		return score
	} else {
		score := 1000
		for i, v := range spaces {
			if v == 0 {
				nextBoard := board.Copy()
				nextBoard.SetSpace(currentPlayer, i)

				score = min(score, p.alphaBeta(nextBoard, depth-1, alpha, beta, otherPlayer))
				beta = min(beta, score)

				if beta <= alpha {
					return beta
				}
			}
		}
		return score
	}
}

func (p AiPlayer) score(winner int) int {
	if p.isMaximizingPlayer(winner) {
		return 1
	} else if winner == 0 {
		return 0
	} else {
		return -1
	}
}

func (p AiPlayer) isMaximizingPlayer(currentPlayer int) bool {
	return currentPlayer == p.aiPlayer
}

func max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x int, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func bestMove(scores map[int]int) int {
	bestKey := -1
	bestValue := -2

	for k, v := range scores {
		if v > bestValue {
			bestKey = k
			bestValue = v
		}
	}

	return bestKey
}

func otherPlayer(player int) int {
	if player == -1 {
		return 1
	} else if player == 1 {
		return -1
	} else {
		return 0
	}
}

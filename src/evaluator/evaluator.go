package evaluator

type Evaluator interface {
	Winner(spaces []int) int
}

package evaluator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEvaluator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evaluator Suite")
}

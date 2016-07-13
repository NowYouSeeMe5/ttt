package board

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBoard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Board Suite")
}

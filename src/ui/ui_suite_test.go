package ui_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ui Suite")
}

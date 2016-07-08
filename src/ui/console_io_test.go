package ui_test

import (
	"bytes"
	"io"
	"os"

	. "ttt/src/ui"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func writeMockOutput(message string, function func(string)) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	function(message)

	output := make(chan string)
	go func() {
		var buffer bytes.Buffer
		io.Copy(&buffer, r)
		output <- buffer.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-output

	return out
}

func writeMockInput(input string) {
	reader, writer, _ := os.Pipe()
	os.Stdin = reader
	writer.WriteString(input)
	writer.Close()
}

var _ = Describe("ConsoleIO", func() {

	testString := "test"
	consoleIO := new(ConsoleIO)

	Describe("Get input", func() {

		It("Gets message to stdin", func() {
			writeMockInput(testString)

			Expect(testString).To(Equal(consoleIO.GetInput()))
		})
	})

	Describe("Print", func() {
		It("Print output to stdout", func() {
			printFunction := func(s string) { consoleIO.Print(s) }
			printedString := writeMockOutput(testString, printFunction)

			Expect(testString).To(Equal(printedString))
		})
	})
})

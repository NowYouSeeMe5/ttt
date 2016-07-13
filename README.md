# ttt
tic-tac-toe implemented in the style of Go

# Install Go

Install [Go](https://golang.org/doc/install) making sure the appropriate paths are set for the [Go command](https://golang.org/doc/articles/go_command.html) to work properly.

# Running the project

Use `go get github.com/NowYouSeeMe5/ttt` to retrieve the project.
Now just run `ttt` to play a game of tic-tac-toe.

#Testing the project

Grab [ginkgo](https://github.com/onsi/ginkgo) and [gomega](https://github.com/onsi/gomega) with `go get`:

```go get github.com/onsi/ginkgo/ginkgo

   go get github.com/onsi/gomega```
   
Navigate to the ttt folder and run `ginkgo -r` to run the entire suite or navigate to each package and run `ginkgo` to test them individually.

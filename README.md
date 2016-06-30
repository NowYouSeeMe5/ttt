# ttt
tic-tac-toe implemented in the style of Go

Let me apologize in advance for the following steps, welcome to Go.

Go Instructions:

  Homebrew is the easiest way to install go: `brew install go`

  If you already have a workspace setup you can clone the project and skip these steps.

  It you don't have go there are several setup steps, this should work with a copy paste once you have go:
  
  ```
    mkdir ~/Desktop/go
    mkdir ~/Desktop/go/bin
    mkdir ~/Desktop/go/src
    mkdir ~/Desktop/go/pkg

    cd ~/Desktop/go/src

    git clone https://github.com/NowYouSeeMe5/ttt.git

    export GOPATH="$HOME/Desktop/go"
    export PATH="$PATH:$GOPATH/bin"
    export PATH="/usr/local/bin:$PATH"
    
    cd ~/Desktop/go/src/ttt
    
  
  ```

Build and Run Instructions:

From there you can build an executable using `go build` and run that executable with `./ttt`

Test Instructions:

Testing requires [ginkgo](https://github.com/onsi/ginkgo) and [gomega](https://github.com/onsi/gomega)

Grab these with 
```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```

Grab these and run the specs using `ginkgo -r`

You can run individual specs in each of the package folders by navigating to them and running `ginkgo`

Deletion Instructions:

```
rm -rf ~/Desktop/go
brew uninstall go
```


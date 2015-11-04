# Welcome to Women Who Go Boston
We're excited to start on this Go journey together, build some cool stuff, and get to know each other more through meeting and events. This repo will contain all the projects we'll be working on, along with documentation and other relevant information.

## Development Environment
You can feel free to choose any text editor, or IDE. Personally, I use Atom, the command line, and a web browser. I've found that it's all that you really need for what we're going to be doing. 

The packages that will be the most handy will be the `terminal-plus` package inside Atom which let's you have the terminal embedded into the application. To install you, you can go to `Atom > Preferences > Install` and search for the package there.

## Installing Go
We are going to install the binary version instead of building from source to keep things simple. 
Going to [https://golang.org/dl/] you will find the current binary for your system. I am using Go version 1.5.1, once you install, you can check to see what version you have with: 

`go version` 

And you should see something like:

`go version go1.5.1 darwin/amd64`

### Make sure Go is in your PATH
`echo $PATH` should return something like `/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin` If you don't see it in there, add it to your path. Something like: `PATH = $PATH:/usr/local/go/bin` (wherever you have Go installed, which you can find by running: `which go`).

### Set your GOPATH
The GOPATH lets you know where to find packages you're importing in your application. Create a new folder called `go` in your `Users` folder that will contain all of these packages. It will also contain all the projects we'll be working on here. Setting your GOPATH is exactly the same as setting your PATH:

`export GOPATH=$HOME/go`

### Additional Installation Help
Google has complete documentation incase you run into any install issues: Getting Started https://golang.org/doc/install

## Go Tools

### Fmt
`go fmt` formats your Go code so that it meets Go standards and keeps it consistent across different contributors. 

### Vet
`go vet` vets your program for mistakes.

### Imports
`go get code.google.com/p/go.tools/cmd/goimports` - this fixes all issues with imports, and will rewrite your import statement based on the strcuture of your code. 
`brew install mercurial` (on Mac, if you don't have Mercurial installed) 

Running: `goimports -w *. go` will remove all unused imports, and add imports that haven't been imported explicitly. 

## Contact information
Feel free to email me (Alex) or Jen if you have any questions or suggestions for the meetups. 

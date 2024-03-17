package main

import (
	"go-monkey-interpreter/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

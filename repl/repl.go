package repl

import (
	"bufio"
	"fmt"
	"io"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		print(line)
		io.WriteString(out, "\n")
	}
}

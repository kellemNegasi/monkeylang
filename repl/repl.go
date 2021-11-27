// Package repl contains utilities for interactive console or ReadPrintLoop(repl).
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kellemNegasi/monkeylang/token"

	"github.com/kellemNegasi/monkeylang/lexer"
)

// PROMPT defines an entry point that prompts the user to enter input.
const PROMPT = ">>"

// Start starts the repl.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

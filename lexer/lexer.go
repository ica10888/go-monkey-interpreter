package lexer

import "go-monkey-interpreter/token"

type Lexer interface {
	NextToken() token.Token
}

type lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

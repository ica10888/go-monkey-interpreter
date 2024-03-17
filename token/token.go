package token

type Type string

const (
	ILLEGAL   Type = "ILLEGAL"
	EOF       Type = "EOF"
	IDENT     Type = "IDENT"
	INT       Type = "INT"
	FLOAT     Type = "FLOAT"
	STRING    Type = "STRING"
	BANG      Type = "!"
	ASSIGN    Type = "="
	PLUS      Type = "+"
	MINUS     Type = "-"
	ASTARISK  Type = "*"
	SLASH     Type = "/"
	LT        Type = "<"
	GT        Type = ">"
	EQ        Type = "=="
	NEQ       Type = "!="
	COMMA     Type = ","
	SEMICOLON Type = ";"
	COLON     Type = ":"
	LPAREN    Type = "("
	RPAREN    Type = ")"
	LBRACE    Type = "{"
	RBRACE    Type = "}"
	LBRACKET  Type = "["
	RBRACKET  Type = "]"
	FUNCTION  Type = "FUNCTION"
	LET       Type = "LET"
	TRUE      Type = "TRUE"
	FALSE     Type = "FALSE"
	IF        Type = "IF"
	ELSE      Type = "ELSE"
	RETURN    Type = "RETURN"
	MACRO     Type = "MACRO"
)

type Token struct {
	Type    Type
	Literal string
}

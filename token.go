package main

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	IDENTIFIER
	KEYWORD
	PROPERTY
	VALUE
	TYPE

	ESCAPED_STRING
	RAW_STRING
)

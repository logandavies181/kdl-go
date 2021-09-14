package main

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	NEWLINE

	SINGLE_LINE_COMMENT
	MULTI_LINE_COMMENT

	IDENTIFIER
	KEYWORD
	PROPERTY
	VALUE
	TYPE

	ESCAPED_STRING
	RAW_STRING

	NUMBER

	SEMICOLON
)

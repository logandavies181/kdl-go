package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


type testcase struct {
	input, expected string
}

func newStringScanner(str string) *Scanner {
	reader := strings.NewReader(str)
	return NewScanner(reader)
}

func newStringParser(str string) *Parser {
	reader := strings.NewReader(str)
	return NewParser(reader)
}

func TestParse(t *testing.T) {
	teststr := `foo bar="baz"`

	parser := newStringParser(teststr)

	tok1, _ := parser.scanIgnoreWhitespace()
	tok2, _ := parser.scanIgnoreWhitespace()

	assert.Equal(t, IDENT, tok1)
	assert.Equal(t, ATTRIBUTE, tok2)
}

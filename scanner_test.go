package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanRawString(t *testing.T) {
	testcases := []testcase{
		{
			input: `r#foobarbaz#`,
			expected: `r#foobarbaz#`,
		},
		{
			input: `r#foobarbaz# 
			  `,
			expected: `r#foobarbaz#`,
		},
		{
			input: `r"foobarbaz" 
			  `,
			expected: `r"foobarbaz"`,
		},
		{
			input: `r"foobarbaz
			 "`,
			expected: `r"foobarbaz
			 "`,
		},
	}

	expectedTok := STRING

	for _, tc := range testcases {

		tok, lit := newStringScanner(tc.input).scanRawString()

		assert.Equal(t, expectedTok, tok)
		assert.Equal(t, tc.expected, lit)
	}
}

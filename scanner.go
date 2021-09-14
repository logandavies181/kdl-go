package main

import (
	"bufio"
	"bytes"
	"io"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}


// Scan returns the next token and literal value; determining its type down to Identifier, Keyword, Property, Value or
// Type.
func (s *Scanner) Scan() (tok Token, lit string) {
	var buf bytes.Buffer

	//
	// check first char
	//
	ch := s.read()
	switch {
	case ch == _EOF:
		return EOF, ""
	case ch == 'r':
		// check for raw string
		if ch2 := s.read(); ch2 == _EOF {
			return IDENTIFIER, "r"
		} else if ch2 == '#' || ch2 == '"' {
			s.unread()
			s.unread()
			return s.scanRawString()
		} else {
			s.unread()
			buf.WriteRune('r')
		}
	case ch == '"':
		s.unread()
		return s.scanEscapedString()
	case isSign(ch):
		// check if identifier or number
		ch2 := s.read()
		switch {
		case isDigit(ch2):
			s.unread()
			_, numberStr := s.scanNumber()

			return VALUE, string(ch)+numberStr
		case isIdentifierChar(ch2):
			s.unread()
			_, idStr := s.scanIdentifier()

			return IDENTIFIER, string(ch)+idStr
		}
	case isWhitespace(ch):
		s.unread()
		return s.scanWhitespace()

	case !isIdentifierChar(ch):
		// check if this case makes any sense

		// Not a valid identifier
		s.unread()
		return ILLEGAL, buf.String()
	default:
		_, _ = buf.WriteRune(ch)
	}

	//
	// check the rest of the chars
	//
	for {
		if ch := s.read(); ch == _EOF {
			break
		} else if ch == '=' {
			// found property. now scan for value
			_, value := s.scanValue()
			return PROPERTY, buf.String() + "=" + value
		} else if !isIdentifierChar(ch) {
			// TODO
			// should check for semicolon and stuff here
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	identifier := buf.String()
	if identifier == "true" || identifier == "false" || identifier == "null" {
		return KEYWORD, identifier
	}

	// Otherwise return as a regular identifier.
	return IDENTIFIER, identifier
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
// TODO update to split between ws and newlines
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == _EOF {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	// TODO: check err type
	if err != nil {
		return _EOF
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// scanValue
func (s *Scanner) scanValue() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == _EOF {
			break
		} else if ch == '"' {
			s.unread()
			_, str := s.scanEscapedString()
			buf.WriteString(str)
			break
		} else if !isLetter(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			// TODO: handle error
			_, err := buf.WriteRune(ch)
			if err != nil {
				panic(err)
			}
		}
	}

	// TODO: match keywords

	// Otherwise return as a regular identifier.
	return VALUE, buf.String()
}

func (s *Scanner) scanEscapedString() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// TODO: handle escapes, and invalid chars
	for {
		if ch := s.read(); ch == _EOF {
			break
		} else if ch == '"' {
			// TODO: handle error
			buf.WriteRune(ch)

			break
		} else {
			// TODO: handle error
			buf.WriteRune(ch)
		}
	}

	return ESCAPED_STRING, buf.String()
}

func (s *Scanner) scanRawString() (tok Token, lit string) {
	var buf bytes.Buffer

	if ch := s.read(); ch == _EOF {
		panic("RawString attempted to read but starts with wrong char")
	} else if ch == 'r' {
		buf.WriteRune(ch)
	}

	var delimiter rune
	switch ch := s.read(); ch {
	case '#', '"':
		delimiter = ch
		buf.WriteRune(ch)
	default:
		panic("RawString attempted to read but second char is wrong char")
	}

	// Scan over the raw string
	for {
		if ch := s.read(); ch == _EOF {
			break
		} else if ch == delimiter {
			buf.WriteRune(ch)

			break
		} else {
			// TODO: handle error
			buf.WriteRune(ch)
		}
	}

	return RAW_STRING, buf.String()
}

func (s *Scanner) scanType() (tok Token, lit string) {
	var buf bytes.Buffer

	if ch := s.read(); ch == '(' {
		buf.WriteRune(ch)
	} else {
		return ILLEGAL, buf.String()
	}

	_, identifier := s.scanIdentifier()

	switch ch := s.read(); ch {
	case _EOF:
		return ILLEGAL, buf.String()
	case ')':
		buf.WriteString(identifier)
		buf.WriteRune(ch)
	default:
		return ILLEGAL, buf.String()
	}

	return TYPE, buf.String()
}

// assumes that the start of the string has been checked for type and reads the rest of the identifier
func (s *Scanner) scanIdentifier() (tok Token, lit string) {
	var buf bytes.Buffer

	for {
		ch := s.read()
		switch {
		case isIdentifierChar(ch):
			buf.WriteRune(ch)
		case isWhitespace(ch):
			s.unread()
			break
		default:
			return ILLEGAL, buf.String()
		}
	}

	return IDENTIFIER, buf.String()
}

func (s *Scanner) scanNumber() (tok Token, lit string) {
	var buf bytes.Buffer

	return ILLEGAL, buf.String()
}

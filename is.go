package main

func isWhitespace(ch rune) bool {
	if isUnicodeSpace(ch) || isNewline(ch) {
		return true
	}
	return false
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isUnicodeSpace(ch rune) bool {
	switch ch {
	case 0x0009,
		0x0020,
		0x00A0,
		0x1680,
		0x2000,
		0x2001,
		0x2002,
		0x2003,
		0x2004,
		0x2005,
		0x2006,
		0x2007,
		0x2008,
		0x2009,
		0x200A,
		0x202F,
		0x205F,
		0x3000:
		return true
	default:
		return false
	}
}

func isNewline(ch rune) bool {
	switch ch {
	case 0x000D, // CR
		0x000A, // LF
		0x0085, // NEL
		0x000C, // FF
		0x2028, // FF
		0x2029: // PS
		return true
	default:
		return false
	}
}

func isIdentifierChar(ch rune) bool {
	switch ch {
	case '\\', '/', '(', ')', '{', '}', '<', '>', ';', '[', ']', '=', ',', '"':
		return false
	}

	if isWhitespace(ch) || ch < 0x20 || ch > 0x10FFFF {
		return false
	}

	return true
}

func isSign(ch rune) bool {
	return ch == '+' || ch == '-'
}

// isLetter returns true if the rune is a letter.
// TODO: use isUnicodeChar instead
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var _EOF = rune(0)

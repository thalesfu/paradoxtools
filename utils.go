package paradoxtools

import (
	"unicode"
)

func IsAlphanumeric(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

func IsWhitespace(b byte) bool {
	r := rune(b)
	return unicode.IsSpace(r)
}

func IsNewlineStart(b byte) bool {
	return b == '\r'
}

func IsNewlineEnd(b byte) bool {
	return b == '\n'
}

func IsLeftParentheses(b byte) bool {
	return b == '{'
}

func IsRightParentheses(b byte) bool {
	return b == '}'
}

func IsEqual(b byte) bool {
	return b == '='
}

func IsQuote(b byte) bool {
	return b == '"'
}

func IsSharp(b byte) bool {
	return b == '#'
}

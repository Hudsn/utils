package parsetime

import (
	"fmt"
	"slices"
)

type lexer struct {
	input       []rune
	currentChar rune
	currentIdx  int
	nextIdx     int
}

const nullchar = rune(0)

func newLexer(input string) *lexer {
	l := &lexer{
		input:      []rune(input),
		currentIdx: 0,
		nextIdx:    0,
	}
	l.next()
	return l
}

func (l *lexer) next() {
	if l.nextIdx >= len(l.input) {
		l.currentChar = nullchar
	} else {
		l.currentChar = l.input[l.nextIdx]
	}
	l.currentIdx = min(l.nextIdx, len(l.input))
	l.nextIdx = min(l.nextIdx+1, len(l.input))
}

func (l *lexer) peek() rune {
	if l.nextIdx >= len(l.input) {
		return nullchar
	}
	return l.input[l.nextIdx]
}

func (l *lexer) nextToken() token {

	switch l.currentChar {
	case nullchar:
		return token{tokenType: eof, start: l.currentIdx, end: l.nextIdx, value: ""}
	case '%':
		tok := l.handleShortcode()
		l.next()
		return tok
	default:
		return l.handleLiteral()
	}
}

func (l *lexer) handleShortcode() token {

	start := l.currentIdx

	if isDigit(l.peek()) {
		return l.shortcodeHandleNumberModifier() // ends on "last char of shortcode"
	}
	if !slices.Contains(directiveCharList, l.peek()) {
		return token{
			tokenType: illegal,
			start:     start,
			end:       l.nextIdx,
			value:     fmt.Sprintf("invalid shortcode: %s", string(l.input[start:l.nextIdx+1])),
		}
	}

	l.next() // on directive character

	return token{
		tokenType: directive,
		start:     start,
		end:       l.nextIdx,
		value:     string(l.input[start:l.nextIdx]),
	}
}

func (l *lexer) shortcodeHandleNumberModifier() token {
	start := l.currentIdx
	l.next()               // now on char after %
	if isDigit(l.peek()) { // 2 digits in a row are not allowed
		return token{
			tokenType: illegal,
			start:     start,
			end:       l.nextIdx,
			value:     "numeric modifiers cannot be more than a single digit",
		}
	}
	if l.peek() != 'N' { // the only time we should see digits is fractional seconds, which should be in the format %{digt}N
		return token{
			tokenType: illegal,
			start:     start,
			end:       l.nextIdx,
			value:     "numeric modifiers must be followed by the shortcode for fractional seconds (ex: '%9N')",
		}
	}
	l.next() // now on "N"

	return token{
		tokenType: directive,
		start:     start,
		end:       l.nextIdx,
		value:     string(l.input[start:l.nextIdx]),
	}

}

func (l *lexer) handleLiteral() token {
	start := l.currentIdx
	value := []rune{}
	shouldContineLoop := true
	for l.currentChar != nullchar && shouldContineLoop {
		toAdd := l.currentChar
		switch l.currentChar {
		case '%':
			if l.peek() == '%' {
				toAdd = '%'
			} else {
				shouldContineLoop = false
				continue
			}
		case '\\':
			var ok bool
			toAdd, ok = l.handleEscapeRune()
			if !ok {
				return token{
					tokenType: illegal,
					start:     start,
					end:       l.nextIdx,
					value:     fmt.Sprintf("unrecognized escape sequence: %s", string(l.input[l.currentIdx:l.nextIdx+1])),
				}
			}
		default:
		}
		value = append(value, toAdd)
		l.next()
	}

	return token{
		tokenType: literal,
		start:     start,
		end:       l.nextIdx,
		value:     string(value),
	}
}
func (l *lexer) handleEscapeRune() (rune, bool) {
	var ret rune = l.currentChar
	switch l.peek() {
	case '\\':
		ret = '\\'
	case 't':
		ret = '\t'
	case 'n':
		ret = '\n'
	case 'r':
		ret = '\r'
	default:
		return nullchar, false
	}
	l.next() // to escaped char
	return ret, true
}

func isLegalChar(char rune) bool {
	return slices.Contains([]rune{'\r', '\n'}, char)
}

// helpers
func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func (l *lexer) stringFromToken(t token) string {
	if t.tokenType == eof {
		return ""
	}
	return string(l.input[t.start:t.end])
}

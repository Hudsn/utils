package parsetime

import (
	"testing"
)

func TestLexLiteral(t *testing.T) {
	input := `justa\\random\nmultilinestring`
	cases := []lexTestCase{
		{
			tokenType:  literal,
			value:      "justa\\random\nmultilinestring",
			start:      0,
			end:        30,
			rangeValue: input,
		},
		{
			tokenType:  eof,
			value:      "",
			start:      30,
			end:        30,
			rangeValue: "",
		},
	}
	checkLexTestCases(t, input, cases)
}

func TestLexFractionalSec(t *testing.T) {
	input := "%4N"
	cases := []lexTestCase{
		{
			tokenType:  directive,
			value:      "%4N",
			start:      0,
			end:        3,
			rangeValue: "%4N",
		},
		{
			tokenType:  eof,
			value:      "",
			start:      3,
			end:        3,
			rangeValue: "",
		},
		{
			tokenType:  eof,
			value:      "",
			start:      3,
			end:        3,
			rangeValue: "",
		},
	}
	checkLexTestCases(t, input, cases)
}

type lexTestCase struct {
	tokenType  tokenType
	value      string
	start      int
	end        int
	rangeValue string // the literal value captured by the start and end markers; for ex in a string "asdf", it would include the quotes as well even though the value is just asdf
}

func checkLexTestCases(t *testing.T, input string, cases []lexTestCase) {
	lexer := newLexer(input)
	for idx, tt := range cases {
		tok := lexer.nextToken()
		if tok.tokenType != tt.tokenType {
			t.Errorf("case %d: wrong token type: want=%+v, got=%+v", idx+1, tt.tokenType, tok.tokenType)
		}
		if tok.value != tt.value {
			t.Errorf("case %d: wrong token value: want=%s, got=%s", idx+1, tt.value, tok.value)
		}
		if tok.start != tt.start {
			t.Errorf("case %d: wrong token start index: want=%d, got=%d", idx+1, tt.start, tok.start)
		}
		if tok.end != tt.end {
			t.Errorf("case %d: wrong token end index: want=%d, got=%d", idx+1, tt.end, tok.end)
		}
		if tt.rangeValue != lexer.stringFromToken(tok) {
			t.Errorf("case %d: wrong token literal derived from range: want=%s, got=%s", idx+1, tt.rangeValue, lexer.stringFromToken(tok))
		}

	}
}

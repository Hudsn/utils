package parsetime

type tokenType int

const (
	_ tokenType = iota
	tok_illegal
	tok_eof
	tok_literal
	tok_directive
)

type token struct {
	tokenType tokenType
	start     int
	end       int
	value     string
}

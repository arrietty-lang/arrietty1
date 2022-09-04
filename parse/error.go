package parse

import (
	"fmt"
	"github.com/x0y14/arrietty/tokenize"
)

type UnexpectedTokenErr struct {
	Token    *tokenize.Token
	Expected string
}

func (e *UnexpectedTokenErr) Error() string {
	return fmt.Sprintf("[%d:%d(%d)] expect %s, but found: %s", e.Token.Pos.LineNo, e.Token.Pos.Lat, e.Token.Pos.Wat, e.Expected, e.Token.Kind.String())
}

func NewUnexpectedTokenErr(expected string, actual *tokenize.Token) *UnexpectedTokenErr {
	return &UnexpectedTokenErr{
		Token:    actual,
		Expected: expected,
	}
}

package tokenize

import "fmt"

type UnexpectedCharacterErr struct {
	Character rune
	Pos       *Position
}

func (e *UnexpectedCharacterErr) Error() string {
	return fmt.Sprintf("[%d:%d(%d)] unexpected: %s", e.Pos.LineNo, e.Pos.Lat, e.Pos.Wat, string(e.Character))
}

func NewUnexpectedCharacterErr(c rune, pos *Position) *UnexpectedCharacterErr {
	return &UnexpectedCharacterErr{
		Character: c,
		Pos:       pos,
	}
}

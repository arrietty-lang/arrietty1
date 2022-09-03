package tokenize

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func GenPosForTest(str string) *Position {
	_wat := utf8.RuneCountInString(str)
	_lat := 0
	_ln := 1
	for _, r := range []rune(str) {
		if r == '\n' {
			_ln++
			_lat = 0
		} else {
			_lat++
		}
	}
	return NewPosition(_ln, _lat, _wat)
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectToken *Token
		expectErr   error
	}{
		{
			"(add int int)",
			"1 + 1",
			&Token{
				Kind: Int,
				Pos:  GenPosForTest(""),
				S:    "",
				F:    0,
				I:    1,
				Next: &Token{
					Kind: White,
					Pos:  GenPosForTest("1"),
					S:    " ",
					F:    0,
					I:    0,
					Next: &Token{
						Kind: Add,
						Pos:  GenPosForTest("1 "),
						S:    "+",
						F:    0,
						I:    0,
						Next: &Token{
							Kind: White,
							Pos:  GenPosForTest("1 +"),
							S:    " ",
							F:    0,
							I:    0,
							Next: &Token{
								Kind: Int,
								Pos:  GenPosForTest("1 + "),
								S:    "",
								F:    0,
								I:    1,
								Next: &Token{
									Kind: Eof,
									Pos:  GenPosForTest("1 + 1"),
									S:    "",
									F:    0,
									I:    0,
									Next: nil,
								},
							},
						},
					},
				},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := Tokenize(tt.in)
			assert.Equal(t, tt.expectToken, tok)
			assert.Equal(t, tt.expectErr, err)

		})
	}
}

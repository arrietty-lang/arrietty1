package interpret

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *Object
	}{
		{
			"1",
			"main() { return 1; }",
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Int,
					Str:      "",
					NumFloat: 0,
					NumInt:   1,
					Items:    nil,
					KVS:      nil,
				}},
		},
		{
			"re",
			"retX(x) { return x; } main() { return retX(2); }",
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Int,
					Str:      "",
					NumFloat: 0,
					NumInt:   2,
					Items:    nil,
					KVS:      nil,
				}},
		},
		{
			"array",
			"list() { return [10, 20, 30, 40]; } main() { return list()[2]; }",
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Int,
					Str:      "",
					NumFloat: 0,
					NumInt:   30,
					Items:    nil,
					KVS:      nil,
				}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			nod := parse.Parse(tok)
			result, err := Interpret(nod)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, result)
		})
	}
}

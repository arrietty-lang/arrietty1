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
		{
			"assign ident",
			"main() { i = 4004; return i; }",
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Int,
					Str:      "",
					NumFloat: 0,
					NumInt:   4004,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"assign dict",
			`main() { d = {"k1": 300}; d["k1"] = "v1"; return d["k1"]; }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     String,
					Str:      "v1",
					NumFloat: 0,
					NumInt:   0,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"assign array",
			`main() { f = 0.1; li = [0, f, "hello"]; return li[1]; }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Float,
					Str:      "",
					NumFloat: 0.1,
					NumInt:   0,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(6000)]}]]; return li[3][0]["k2"][0]; }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Int,
					Str:      "",
					NumFloat: 0,
					NumInt:   6000,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"add string",
			`sayHello(name) { return "hello, " + name; } main() { return sayHello("john"); }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     String,
					Str:      "hello, john",
					NumFloat: 0,
					NumInt:   0,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"add assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(6000+300)]}]]; return li[3][0]["k2"][0] + 20.0; }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Float,
					Str:      "",
					NumFloat: 6320,
					NumInt:   0,
					Items:    nil,
					KVS:      nil,
				},
			},
		},
		{
			"add-sub-mul-div-mod assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(9%2+1-(3*3/(4-3)))]}]]; return li[3][0]["k2"][0] + 20.0; }`,
			&Object{
				Kind: ObjLiteral,
				Literal: &Literal{
					Kind:     Float,
					Str:      "",
					NumFloat: 13,
					NumInt:   0,
					Items:    nil,
					KVS:      nil,
				},
			},
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

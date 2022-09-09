package interpret

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/analyze"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"testing"
)

func TestInterpret(t *testing.T) {
	tests := []struct {
		name      string
		in        string
		expectObj *Object
		expectErr error
	}{
		{
			"int",
			"int main() {return 1;}",
			NewIntObject(1),
			nil,
		},
		{
			"float",
			"float main() {return 1.0;}",
			NewFloatObject(1),
			nil,
		},
		{
			"string",
			"string main() {return \"shu-ka sai-to\";}",
			NewStringObject("shu-ka sai-to"),
			nil,
		},
		{
			"true",
			"bool main() {return true;}",
			NewTrueObject(),
			nil,
		},
		{
			"false",
			"bool main() {return false;}",
			NewFalseObject(),
			nil,
		},
		{
			"null",
			"void main() {return null;}",
			NewNullObject(),
			nil,
		},
		{
			"dict[string]int",
			`dict[string]int main() {return { "k0": 0, "k1": 1 };}`,
			NewDictObject(map[string]*Object{
				"k0": NewIntObject(0),
				"k1": NewIntObject(1),
			}),
			nil,
		},
		{
			"dict[string]dict[string]int",
			`dict[string]dict[string]int main() {return { "k0": {"k1": 0} };}`,
			NewDictObject(map[string]*Object{
				"k0": NewDictObject(map[string]*Object{
					"k1": NewIntObject(0),
				}),
			}),
			nil,
		},
		{
			"dict[string]dict[string]int",
			`dict[string]dict[string]int main() {return { "k0": {"k1": 0} };}`,
			NewDictObject(map[string]*Object{
				"k0": NewDictObject(map[string]*Object{
					"k1": NewIntObject(0),
				}),
			}),
			nil,
		},
		{
			"[]int",
			"[]int main() {return [0, 1]; }",
			NewListObject([]*Object{
				NewIntObject(0),
				NewIntObject(1),
			}),
			nil,
		},
		{
			"[][]int",
			"[][]int main() {return [[0, 1]]; }",
			NewListObject([]*Object{
				NewListObject([]*Object{
					NewIntObject(0),
					NewIntObject(1),
				}),
			}),
			nil,
		},
		{
			"inline assign",
			`int main() {var x int = 0; return x;}`,
			NewIntObject(0),
			nil,
		},
		{
			"var then decl",
			`int main() {var x int; x = 0; return x;}`,
			NewIntObject(0),
			nil,
		},
		{
			"short var-decl",
			`int main() {x := 0; return x;}`,
			NewIntObject(0),
			nil,
		},
		{
			"dict[key] = value",
			`int main() { var d dict[string]int = {}; d["k"] = 0; return d["k"]; }`,
			NewIntObject(0),
			nil,
		},
		{
			"fail list[n] = value",
			`int main() { var l []int = []; l[0] = 30; return l[0]; }`,
			nil,
			fmt.Errorf("index out of range"),
		},
		{
			"list[n] = value",
			`int main() { var l []int = [0, 1, 2]; l[0] = 30; return l[0]; }`,
			NewIntObject(30),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				analyze.CleanUp()
			})
			tokens, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatalf("failed to tokenize: %v", err)
			}

			nodes, err := parse.Parse(tokens)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}

			tops, err := analyze.Analyze(nodes)
			if err != nil {
				t.Fatalf("failed to analyze: %v", err)
			}

			obj, err := Interpret(tops)

			assert.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expectObj, obj)

		})
	}
}

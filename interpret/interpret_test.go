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
			NewReturnValue(NewIntObject(1)),
			nil,
		},
		{
			"float",
			"float main() {return 1.0;}",
			NewReturnValue(NewFloatObject(1)),
			nil,
		},
		{
			"string",
			"string main() {return \"shu-ka sai-to\";}",
			NewReturnValue(NewStringObject("shu-ka sai-to")),
			nil,
		},
		{
			"true",
			"bool main() {return true;}",
			NewReturnValue(NewTrueObject()),
			nil,
		},
		{
			"false",
			"bool main() {return false;}",
			NewReturnValue(NewFalseObject()),
			nil,
		},
		{
			"null",
			"void main() {return null;}",
			NewReturnValue(NewNullObject()),
			nil,
		},
		{
			"dict[string]int",
			`dict[string]int main() {return { "k0": 0, "k1": 1 };}`,
			NewReturnValue(NewDictObject(map[string]*Object{
				"k0": NewIntObject(0),
				"k1": NewIntObject(1),
			})),
			nil,
		},
		{
			"dict[string]dict[string]int",
			`dict[string]dict[string]int main() {return { "k0": {"k1": 0} };}`,
			NewReturnValue(NewDictObject(map[string]*Object{
				"k0": NewDictObject(map[string]*Object{
					"k1": NewIntObject(0),
				}),
			})),
			nil,
		},
		{
			"dict[string]dict[string]int",
			`dict[string]dict[string]int main() {return { "k0": {"k1": 0} };}`,
			NewReturnValue(NewDictObject(map[string]*Object{
				"k0": NewDictObject(map[string]*Object{
					"k1": NewIntObject(0),
				}),
			})),
			nil,
		},
		{
			"[]int",
			"[]int main() {return [0, 1]; }",
			NewReturnValue(NewListObject([]*Object{
				NewIntObject(0),
				NewIntObject(1),
			})),
			nil,
		},
		{
			"[][]int",
			"[][]int main() {return [[0, 1]]; }",
			NewReturnValue(NewListObject([]*Object{
				NewListObject([]*Object{
					NewIntObject(0),
					NewIntObject(1),
				}),
			})),
			nil,
		},
		{
			"inline assign",
			`int main() {var x int = 0; return x;}`,
			NewReturnValue(NewIntObject(0)),
			nil,
		},
		{
			"var then decl",
			`int main() {var x int; x = 0; return x;}`,
			NewReturnValue(NewIntObject(0)),
			nil,
		},
		{
			"short var-decl",
			`int main() {x := 0; return x;}`,
			NewReturnValue(NewIntObject(0)),
			nil,
		},
		{
			"dict[key] = value",
			`int main() { var d dict[string]int = {}; d["k"] = 0; return d["k"]; }`,
			NewReturnValue(NewIntObject(0)),
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
			NewReturnValue(NewIntObject(30)),
			nil,
		},
		{
			"int + int",
			`int main() { return 1+2; }`,
			NewReturnValue(NewIntObject(3)),
			nil,
		},
		{
			"int + int + int",
			`int main() { return 1+2+3; }`,
			NewReturnValue(NewIntObject(6)),
			nil,
		},
		{
			"float + int",
			`float main() { return 1.0+2; }`,
			NewReturnValue(NewFloatObject(3)),
			nil,
		},
		{
			"float + int + int",
			`float main() { return 1.0+2+3; }`,
			NewReturnValue(NewFloatObject(6)),
			nil,
		},
		{
			"for",
			`int main() { var sum int = 0; for(i:=0; i<3; i=i+1){ sum = sum + i; } return sum; }`,
			NewReturnValue(NewIntObject(3)),
			nil,
		},
		{
			"while",
			`int main() { var sum int = 30; while(sum > 0) { sum = sum - 1; } return sum; }`,
			NewReturnValue(NewIntObject(0)),
			nil,
		},
		{
			"if",
			`int main() { name := "john"; if (name == "john") { return 10; } return 1000;}`,
			NewReturnValue(NewIntObject(10)),
			nil,
		},
		{
			"if",
			`int main() { name := "tom"; if (name == "john") { return 10; } return 1000; }`,
			NewReturnValue(NewIntObject(1000)),
			nil,
		},
		{
			"if else",
			`int main() { name := "john"; if (name == "john") { return 10; } else { return 1000; } }`,
			NewReturnValue(NewIntObject(10)),
			nil,
		},
		{
			"if else",
			`int main() { name := "tom"; if (name == "john") { return 10; } else { return 1000; } }`,
			NewReturnValue(NewIntObject(1000)),
			nil,
		},
		{
			"call",
			`string hello(name string) { return "hello, " + name; } void main() { hello("john"); }`,
			NewReturnValue(NewStringObject("hello, john")),
			nil,
		},
		//{
		//	"tarai",
		//	`int tarai(x int, y int, z int) {
		//			if (x > y) {
		//				return tarai(tarai(x-1, y, z), tarai(y-1, z, x), tarai(z-1, x, y));
		//			} else {
		//				return y;
		//			}
		//		}
		//		int main() {
		//			return tarai(14, 7, 0);
		//		}`,
		//	NewReturnValue(NewIntObject(14)),
		//	nil,
		//},// 4分半かかるから使うとgithub action credit枯渇する
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

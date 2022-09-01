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
			NewInt(1).AsRet(),
		},
		{
			"re",
			"retX(x) { return x; } main() { return retX(2); }",
			NewInt(2).AsRet(),
		},
		{
			"array",
			"list() { return [10, 20, 30, 40]; } main() { return list()[2]; }",
			NewInt(30).AsRet(),
		},
		{
			"assign ident",
			"main() { i = 4004; return i; }",
			NewInt(4004).AsRet(),
		},
		{
			"assign dict",
			`main() { d = {"k1": 300}; d["k1"] = "v1"; return d["k1"]; }`,
			NewString("v1").AsRet(),
		},
		{
			"assign array",
			`main() { f = 0.1; li = [0, f, "hello"]; return li[1]; }`,
			NewFloat(0.1).AsRet(),
		},
		{
			"assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(6000)]}]]; return li[3][0]["k2"][0]; }`,
			NewInt(6000).AsRet(),
		},
		{
			"add string",
			`sayHello(name) { return "hello, " + name; } main() { return sayHello("john"); }`,
			NewString("hello, john").AsRet(),
		},
		{
			"add assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(6000+300)]}]]; return li[3][0]["k2"][0] + 20.0; }`,
			NewFloat(6320).AsRet(),
		},
		{
			"add-sub-mul-div-mod assign array array",
			`retX(x){return x;} main() { li = [0, 0, "hello", [{"k2": [retX(9%2+1-(3*3/(4-3)))]}]]; return li[3][0]["k2"][0] + 20.0; }`,
			NewFloat(13).AsRet(),
		},
		// eq
		{
			"same string",
			`main() { s = "hello"; return "hello" == s; }`,
			NewTrue().AsRet(),
		},
		{
			"not same string",
			`main() { s = "hello"; return "hello11" == s; }`,
			NewFalse().AsRet(),
		},
		{
			"same int",
			`main() {i = 1; return 1 == i;}`,
			NewTrue().AsRet(),
		},
		{
			"not same int",
			`main() {i = 1; return 99 == i;}`,
			NewFalse().AsRet(),
		},
		{
			"same float",
			`main() {f = 1.0; return 1.0 == f; }`,
			NewTrue().AsRet(),
		},
		{
			"not same float",
			`main() {f = 1.0; return 1.1 == f; }`,
			NewFalse().AsRet(),
		},
		// ne
		{
			"same string",
			`main() { s = "hello"; return "hello" != s; }`,
			NewFalse().AsRet(),
		},
		{
			"not same string",
			`main() { s = "hello"; return "hello11" != s; }`,
			NewTrue().AsRet(),
		},
		{
			"same int",
			`main() {i = 1; return 1 != i;}`,
			NewFalse().AsRet(),
		},
		{
			"not same int",
			`main() {i = 1; return 99 != i;}`,
			NewTrue().AsRet(),
		},
		{
			"same float",
			`main() {f = 1.0; return 1.0 != f; }`,
			NewFalse().AsRet(),
		},
		{
			"not same float",
			`main() {f = 1.0; return 1.1 != f; }`,
			NewTrue().AsRet(),
		},
		// lt
		{
			"lt int",
			`main() { t = 1 < 5; f = 5 < 1; return t != f; }`,
			NewTrue().AsRet(),
		},
		{
			"lt float",
			`main() { t = 1.1 < 5.1; f = 5.1 < 1.1; return t != f; }`,
			NewTrue().AsRet(),
		},
		// le
		{
			"le int",
			`main() { t = 5 <= 5; f = 5 <= 1; return t != f; }`,
			NewTrue().AsRet(),
		},
		{
			"le float",
			`main() { t = 5.1 <= 5.1; f = 5.1 <= 1.1; return t != f; }`,
			NewTrue().AsRet(),
		},
		// gt
		{
			"gt int",
			`main() { f = 1 > 5; t = 5 > 1; return f != t; }`,
			NewTrue().AsRet(),
		},
		{
			"gt float",
			`main() { f = 1.1 > 5.1; t = 5.1 > 1.1; return f != t; }`,
			NewTrue().AsRet(),
		},
		{
			"ge int",
			`main() { t = 5 >= 5; t2 = 5 >= 1; return t == t2; }`,
			NewTrue().AsRet(),
		},
		{
			"ge float",
			`main() { t = 5.1 >= 5.1; t2 = 5.1 >= 1.1; return t == t2; }`,
			NewTrue().AsRet(),
		},
		// not
		{
			"not",
			`main() {return !true;}`,
			NewFalse().AsRet(),
		},
		{
			"not",
			`main() {return !false;}`,
			NewTrue().AsRet(),
		},
		// or
		{
			"or",
			`main() {return true || true;}`,
			NewTrue().AsRet(),
		},
		{
			"or",
			`main() {return true || false;}`,
			NewTrue().AsRet(),
		},
		{
			"or",
			`main() {return false || false;}`,
			NewFalse().AsRet(),
		},
		// and
		{
			"and",
			`main() {return true && true;}`,
			NewTrue().AsRet(),
		},
		{
			"and",
			`main() {return true && false;}`,
			NewFalse().AsRet(),
		},
		{
			"and",
			`main() {return false && false;}`,
			NewFalse().AsRet(),
		},
		// not and or
		{
			"not and or",
			`main() { f = !false && false; t = !(!true || true); return !(f && t); }`,
			NewTrue().AsRet(),
		},
		{
			"if",
			`main() { if (30 < 40) { return 30; } return 40; }`,
			NewInt(30).AsRet(),
		},
		{
			"if",
			`main() { if (30 > 40) { return 30; } return 40; }`,
			NewInt(40).AsRet(),
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

package interpret

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/analyze"
	"github.com/x0y14/arrietty/apm"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"log"
	"path/filepath"
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
		{
			"int + int",
			`int main() { return 1+2; }`,
			NewIntObject(3),
			nil,
		},
		{
			"int + int + int",
			`int main() { return 1+2+3; }`,
			NewIntObject(6),
			nil,
		},
		{
			"float + int",
			`float main() { return 1.0+2; }`,
			NewFloatObject(3),
			nil,
		},
		{
			"float + int + int",
			`float main() { return 1.0+2+3; }`,
			NewFloatObject(6),
			nil,
		},
		{
			"for",
			`int main() { var sum int = 0; for(i:=0; i<3; i=i+1){ sum = sum + i; } return sum; }`,
			NewIntObject(3),
			nil,
		},
		{
			"while",
			`int main() { var sum int = 30; while(sum > 0) { sum = sum - 1; } return sum; }`,
			NewIntObject(0),
			nil,
		},
		{
			"if",
			`int main() { name := "john"; if (name == "john") { return 10; } return 1000;}`,
			NewIntObject(10),
			nil,
		},
		{
			"if",
			`int main() { name := "tom"; if (name == "john") { return 10; } return 1000; }`,
			NewIntObject(1000),
			nil,
		},
		{
			"if else",
			`int main() { name := "john"; if (name == "john") { return 10; } else { return 1000; } }`,
			NewIntObject(10),
			nil,
		},
		{
			"if else",
			`int main() { name := "tom"; if (name == "john") { return 10; } else { return 1000; } }`,
			NewIntObject(1000),
			nil,
		},
		{
			"call",
			`string hello(name string) { return "hello, " + name; } string main() { return hello("john"); }`,
			NewStringObject("hello, john"),
			nil,
		},
		{
			"not true",
			`bool main() { return !true; }`,
			NewFalseObject(),
			nil,
		},
		{
			"not false",
			`bool main() { return !false; }`,
			NewTrueObject(),
			nil,
		},
		{
			"not true",
			`bool main() { return !(1==1); }`,
			NewFalseObject(),
			nil,
		},
		{
			"not true && true",
			`bool main() { return !(true&&true); }`,
			NewFalseObject(),
			nil,
		},
		{
			"not true && false",
			`bool main() { return !(true&&false); }`,
			NewTrueObject(),
			nil,
		},
		{
			"not true or true",
			`bool main() { return !(true||true); }`,
			NewFalseObject(),
			nil,
		},
		{
			"not true or false",
			`bool main() { return !(true||false); }`,
			NewFalseObject(),
			nil,
		},
		{
			"built-in str_len",
			`int main() { return 1+strlen("a"); }`,
			NewIntObject(2),
			nil,
		},
		{
			"built-in len",
			`int main() {
					var sum int = 0;
					sum = sum + len( [0,1,2,3,4] ); # 5
					sum = sum + len( ["z","o","s"] ); # 8
					sum = sum + len( [[]] ); # 9
					sum = sum + len( [{"k1":"v1"}, {"k2":"v2"}, {}] ); # 12
					sum = sum + len( [true, false] ); #14
					sum = sum + len( [null, null] ); # 16
					return sum;
				}`,
			NewIntObject(16),
			nil,
		},
		{
			"built-in append",
			`[]any main() {
					var nums []int = [];
					append(nums, 1);
					append(nums, 2);
					append(nums, "s");
					return nums;
				}`,
			NewListObject([]*Object{
				NewIntObject(1),
				NewIntObject(2),
				NewStringObject("s"),
			}),
			nil,
		},
		{
			"built-in print",
			"void main() {" +
				"	var s []string = [];" +
				"	append(s, \"hello\");" +
				"	append(s, \"world\");" +
				"	append(s, \"!\");" +
				"	for (i:=0; i<len(s); i=i+1) {" +
				"		print(s[i] + \"\n\");" +
				"	}" +
				"}",
			nil,
			nil,
		},
		{
			"itos",
			`string main() { return itos(1); }`,
			NewStringObject("1"),
			nil,
		},
		{
			"1",
			`bool main() {var i int = 4; return i%3 == 0;}`,
			NewFalseObject(),
			nil,
		},
		{
			"1",
			`bool main() {var i int = 4; if (i%5 == 0) { return false; } else { return true; } }`,
			NewTrueObject(),
			nil,
		},
		{
			"1",
			`bool main() {var i int = 4; if (i%3 == 0) { return false; } else if (i%5 == 0) { return false; } else { return true;}}`,
			NewTrueObject(),
			nil,
		},
		{
			"split",
			`[]string main() {
					s := split("hello", "");
					return s;
				}`,
			NewListObject([]*Object{
				NewStringObject("h"),
				NewStringObject("e"),
				NewStringObject("l"),
				NewStringObject("l"),
				NewStringObject("o"),
			}),
			nil,
		},
		{
			"split",
			`[]string main() {
					s := split("hello, world", ",");
					return s;
				}`,
			NewListObject([]*Object{
				NewStringObject("hello"),
				NewStringObject(" world"),
			}),
			nil,
		},
		{
			"keys",
			`[]string main() {
					var d dict[string]int;
					d = {
						"k0": 0,
						"k1": 1,
						"k2": 2
					};
					return keys(d);
				}`,
			NewListObject([]*Object{
				NewStringObject("k0"),
				NewStringObject("k1"),
				NewStringObject("k2"),
			}),
			nil,
		},
		{
			"stoi",
			`int main() { return stoi("1000"); }`,
			NewIntObject(1000),
			nil,
		},
		{
			"as_string",
			`string main() { var s any = "hello"; return as_string(s); }`,
			NewStringObject("hello"),
			nil,
		},
		{
			"as_int",
			`int main() { var i any = 4440; return as_int(i); }`,
			NewIntObject(4440),
			nil,
		},
		{
			"as_float",
			`float main() { var f any = 100.343434; return as_float(f); }`,
			NewFloatObject(100.343434),
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
		//	NewIntObject(14),
		//	nil,
		//},// 4分半かかるから使うとgithub action credit枯渇する
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				analyze.ResetSymbols()
			})
			tokens, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatalf("failed to tokenize: %v", err)
			}

			nodes, err := parse.Parse(tokens)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}

			err = analyze.PkgAnalyze("test_pkg", [][]*parse.Node{nodes})
			//tops, err := analyze.Analyze(nodes)
			//if err != nil {
			//	t.Fatalf("failed to analyze: %v", err)
			//}

			packages := analyze.GetAnalyzedPackages()

			Setup()
			err = LoadScript("test_pkg", packages["test_pkg"])
			if err != nil {
				t.Fatalf("failed to load semantics node tree: %v", err)
			}
			obj, err := Interpret("test_pkg", "main")

			assert.Equal(t, tt.expectErr, err)
			assert.Equal(t, tt.expectObj, obj)

		})
	}
}

func TestAttachPkg(t *testing.T) {
	// 念のため実行環境をお掃除
	t.Cleanup(func() {
		analyze.ResetSymbols()
	})
	Setup()

	entryArrAbs, err := filepath.Abs("../examples/test/use_sample_pkg.arr")
	if err != nil {
		log.Fatalf("failed to absoluting: %v", err)
	}
	pkgDirOfEntryArr := filepath.Dir(entryArrAbs)
	entryPkgInfo, err := apm.GetCurrentPackageInfo(pkgDirOfEntryArr)
	if err != nil {
		log.Fatalf("failed to read pkg.json: %v", err)
	}
	entryPkgArrs, err := apm.GetArrFilePathsInCurrent(pkgDirOfEntryArr)
	if err != nil {
		log.Fatalf("failed to get .arr files: %v", err)
	}

	tokens, err := tokenize.FromPaths(entryPkgArrs)
	if err != nil {
		log.Fatalf("failed to tokenize: %v", err)
	}

	syntaxTrees, err := parse.FromTokens(tokens)
	if err != nil {
		log.Fatalf("failed to parse: %v", err)
	}

	err = analyze.PkgAnalyze(entryPkgInfo.Name, syntaxTrees)
	if err != nil {
		log.Fatalf("failed to analyze: %v", err)
	}
	semanticsTrees := analyze.GetAnalyzedPackages()

	Setup()
	for pkgName, semTree := range semanticsTrees {
		err = LoadScript(pkgName, semTree)
		if err != nil {
			log.Fatalf("failed to load semanticsTree: %v", err)
		}
	}

	returnValue, err := Interpret(entryPkgInfo.Name, "main")
	if err != nil {
		log.Fatalf("failed to run function: %s.main: %v", entryPkgInfo.Name, err)
	}

	assert.Equal(t, NewIntObject(0), returnValue)
}

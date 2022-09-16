package parse

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/tokenize"
	"testing"
	"unicode/utf8"
)

func GenPosForTest(str string) *tokenize.Position {
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
	return tokenize.NewPosition(_ln, _lat, _wat)
}

func TestParseNode(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expectNodes []*Node
		expectErr   error
	}{
		{
			"main, direct return(int)",
			"void main() { return 1; }",
			[]*Node{
				NewNodeFunctionDefine(
					GenPosForTest(""),
					NewNodeWithChildren(GenPosForTest(""), Void, nil),
					NewNodeIdent(GenPosForTest("void "), "main"),
					nil,
					NewNodeWithChildren(
						GenPosForTest("void main() "),
						Block,
						[]*Node{
							NewNodeReturn(
								GenPosForTest("void main() { "),
								NewNodeImmediate(
									GenPosForTest("void main() { return "),
									&tokenize.Token{Kind: tokenize.Int, I: 1})),
						},
					),
				),
			},
			nil,
		},
		{
			"invalid, only type",
			"void",
			nil,
			NewUnexpectedTokenErr("Ident", &tokenize.Token{
				Kind: tokenize.Eof,
				Pos:  GenPosForTest("void"),
				S:    "",
				F:    0,
				I:    0,
				Next: nil,
			}),
		},
		{
			"void func def, no return",
			`void sayHello(name string) { v_print("hello, " + name); }`,
			[]*Node{
				NewNodeFunctionDefine(
					GenPosForTest(""),
					NewNodeWithChildren(
						GenPosForTest(""),
						Void,
						nil),
					NewNodeIdent(
						GenPosForTest("void "),
						"sayHello"),
					NewNodeWithChildren(
						GenPosForTest("void sayHello("),
						Params,
						[]*Node{
							NewNode(
								GenPosForTest("void sayHello("),
								Param,
								NewNodeIdent(
									GenPosForTest("void sayHello("),
									"name"),
								NewNodeWithChildren(
									GenPosForTest("void sayHello(name "),
									String,
									nil),
							),
						}),
					NewNodeWithChildren(
						GenPosForTest("void sayHello(name string) "),
						Block,
						[]*Node{
							NewNodeCall(
								GenPosForTest("void sayHello(name string) { "),
								NewNodeIdent(
									GenPosForTest("void sayHello(name string) { "),
									"v_print",
								),
								NewNodeWithChildren(
									GenPosForTest(`void sayHello(name string) { v_print("hello, " `),
									Args,
									[]*Node{
										NewNode(
											GenPosForTest(`void sayHello(name string) { v_print("hello, " `),
											Add,
											NewNodeImmediate(
												GenPosForTest("void sayHello(name string) { v_print("),
												&tokenize.Token{Kind: tokenize.String, S: "hello, "},
											),
											NewNodeIdent(
												GenPosForTest(`void sayHello(name string) { v_print("hello, " + `),
												"name",
											),
										),
									},
								),
							),
						},
					),
				),
			},
			nil,
		}, {
			"variable decl",
			`void main() { var a int; a = 1; var b int = 2; c := 3; }`,
			[]*Node{
				NewNodeFunctionDefine(
					GenPosForTest(""),
					NewNodeWithChildren(
						GenPosForTest(""),
						Void,
						nil,
					),
					NewNodeIdent(
						GenPosForTest("void "),
						"main",
					),
					nil,
					NewNodeWithChildren(
						GenPosForTest("void main() "),
						Block,
						[]*Node{
							NewNode(
								GenPosForTest("void main() { "),
								VarDecl,
								NewNodeIdent(
									GenPosForTest("void main() { var "),
									"a"),
								NewNodeWithChildren(
									GenPosForTest("void main() { var a "),
									Int,
									nil),
							),
							NewNode(
								GenPosForTest("void main() { var a int; a "),
								Assign,
								NewNodeIdent(
									GenPosForTest("void main() { var a int; "),
									"a"),
								NewNodeImmediate(
									GenPosForTest("void main() { var a int; a = "),
									&tokenize.Token{Kind: tokenize.Int, I: 1}),
							),
							NewNode(
								GenPosForTest("void main() { var a int; a = 1; var b int "),
								Assign,
								NewNode(
									GenPosForTest("void main() { var a int; a = 1; "),
									VarDecl,
									NewNodeIdent(
										GenPosForTest("void main() { var a int; a = 1; var "),
										"b"),
									NewNodeWithChildren(
										GenPosForTest("void main() { var a int; a = 1; var b "),
										Int,
										nil),
								),
								NewNodeImmediate(
									GenPosForTest("void main() { var a int; a = 1; var b int = "),
									&tokenize.Token{Kind: tokenize.Int, I: 2}),
							),
							NewNode(
								GenPosForTest("void main() { var a int; a = 1; var b int = 2; c "),
								ShortVarDecl,
								NewNodeIdent(
									GenPosForTest("void main() { var a int; a = 1; var b int = 2; "),
									"c",
								),
								NewNodeImmediate(
									GenPosForTest("void main() { var a int; a = 1; var b int = 2; c := "),
									&tokenize.Token{Kind: tokenize.Int, I: 3},
								),
							),
						}),
				),
			},
			nil,
		}, {
			"variable decl with comment",
			`#this is comment
void main() { var a int; a = 1; var b int = 2; c := 3; }`,
			[]*Node{
				NewNodeComment(
					GenPosForTest(""),
					"this is comment",
				),
				NewNodeFunctionDefine(
					GenPosForTest("#this is comment\n"),
					NewNodeWithChildren(
						GenPosForTest("#this is comment\n"),
						Void,
						nil,
					),
					NewNodeIdent(
						GenPosForTest("#this is comment\nvoid "),
						"main",
					),
					nil,
					NewNodeWithChildren(
						GenPosForTest("#this is comment\nvoid main() "),
						Block,
						[]*Node{
							NewNode(
								GenPosForTest("#this is comment\nvoid main() { "),
								VarDecl,
								NewNodeIdent(
									GenPosForTest("#this is comment\nvoid main() { var "),
									"a"),
								NewNodeWithChildren(
									GenPosForTest("#this is comment\nvoid main() { var a "),
									Int,
									nil),
							),
							NewNode(
								GenPosForTest("#this is comment\nvoid main() { var a int; a "),
								Assign,
								NewNodeIdent(
									GenPosForTest("#this is comment\nvoid main() { var a int; "),
									"a"),
								NewNodeImmediate(
									GenPosForTest("#this is comment\nvoid main() { var a int; a = "),
									&tokenize.Token{Kind: tokenize.Int, I: 1}),
							),
							NewNode(
								GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b int "),
								Assign,
								NewNode(
									GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; "),
									VarDecl,
									NewNodeIdent(
										GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var "),
										"b"),
									NewNodeWithChildren(
										GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b "),
										Int,
										nil),
								),
								NewNodeImmediate(
									GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b int = "),
									&tokenize.Token{Kind: tokenize.Int, I: 2}),
							),
							NewNode(
								GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b int = 2; c "),
								ShortVarDecl,
								NewNodeIdent(
									GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b int = 2; "),
									"c",
								),
								NewNodeImmediate(
									GenPosForTest("#this is comment\nvoid main() { var a int; a = 1; var b int = 2; c := "),
									&tokenize.Token{Kind: tokenize.Int, I: 3},
								),
							),
						}),
				),
			},
			nil,
		},
		{
			"main, direct return",
			"void main() { return; }",
			[]*Node{
				NewNodeFunctionDefine(
					GenPosForTest(""),
					NewNodeWithChildren(GenPosForTest(""), Void, nil),
					NewNodeIdent(GenPosForTest("void "), "main"),
					nil,
					NewNodeWithChildren(
						GenPosForTest("void main() "),
						Block,
						[]*Node{
							NewNodeReturn(
								GenPosForTest("void main() { "),
								nil),
						},
					),
				),
			},
			nil,
		},
		{
			"list",
			"void main() { return [-2]; }",
			[]*Node{
				NewNodeFunctionDefine(
					GenPosForTest(""),
					NewNodeWithChildren(GenPosForTest(""), Void, nil),
					NewNodeIdent(GenPosForTest("void "), "main"),
					nil,
					NewNodeWithChildren(
						GenPosForTest("void main() "),
						Block,
						[]*Node{
							NewNodeReturn(
								GenPosForTest("void main() { "),
								NewNodeWithChildren(
									GenPosForTest("void main() { return "),
									List,
									[]*Node{
										NewNode(
											GenPosForTest("void main() { return ["),
											Minus,
											NewNodeImmediate(
												GenPosForTest("void main() { return [-"),
												&tokenize.Token{Kind: tokenize.Int, I: 2}),
											nil,
										),
									})),
						},
					),
				),
			},
			nil,
		},
		{
			"import, main, direct return",
			"import \"pkg\"; void main() { return; }",
			[]*Node{
				NewNodeImport(
					GenPosForTest(""),
					"pkg"),
				NewNodeFunctionDefine(
					GenPosForTest(`import "pkg"; `),
					NewNodeWithChildren(GenPosForTest(`import "pkg"; `), Void, nil),
					NewNodeIdent(GenPosForTest(`import "pkg"; void `), "main"),
					nil,
					NewNodeWithChildren(
						GenPosForTest(`import "pkg"; void main() `),
						Block,
						[]*Node{
							NewNodeReturn(
								GenPosForTest(`import "pkg"; void main() { `),
								nil),
						},
					),
				),
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			nodes, err := Parse(tok)
			if !assert.Equal(t, tt.expectErr, err) {
				t.Fatalf("%v", err)
			}
			assert.Equal(t, tt.expectNodes, nodes)
		})
	}
}

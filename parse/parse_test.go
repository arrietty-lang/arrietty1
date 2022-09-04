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
			"",
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
			"",
			"void",
			nil,
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

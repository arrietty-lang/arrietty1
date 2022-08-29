package parse

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/tokenize"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect []*Node
	}{
		{
			"add",
			"main() { 1 + 1; }",
			[]*Node{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			node := Parse(tok)
			assert.Equal(t, tt.expect, node)
		})
	}
}
package interpret

import (
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
			"1",
			"int main() {return 1;}",
			NewIntObject(1),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

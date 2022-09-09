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

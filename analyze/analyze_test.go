package analyze

import (
	"github.com/stretchr/testify/assert"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name       string
		in         string
		expectTops []*TopLevel
		expectErr  error
	}{
		{
			"comment only",
			"# hello",
			[]*TopLevel{
				{Kind: TPComment, Comment: NewComment(" hello")},
			},
			nil,
		},
		{
			"main",
			"int main() { v := 1; v = 2; var v2 float = 2; for( i:=0; i < 5; i=i+1) { if(i==0){return 11;} return 1;}}",
			[]*TopLevel{
				{Kind: TPComment, Comment: NewComment(" hello")},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := tokenize.Tokenize(tt.in)
			if err != nil {
				t.Fatalf("failed to tokenize: %v", err)
			}

			nodes, err := parse.Parse(tok)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
			}

			tops, err := Analyze(nodes)
			if !assert.Equal(t, tt.expectErr, err) {
				t.Fatalf("failed to analyze: %v", err)
			}
			assert.Equal(t, tt.expectTops, tops)
		})
	}
}

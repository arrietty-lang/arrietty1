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
			`void f() { var x dict[string][][]dict[string]int = {};  x["k"] = [[{"v":1}]];  x["k"][0][0]["v"]=1; }`,
			[]*TopLevel{
				{Kind: TPComment, Comment: NewComment(" hello")},
			},
			nil,
		},
		{
			"",
			`int main() {
    for (i:=0; i<100; i=i+1) {
        if (i%15 == 0) {
            print( itos(i) + " fizzbuzz" + "\n" );
        } else if (i%3 == 0) {
            print( itos(i) + " fizz" + "\n" );
        } else if (i%5 == 0) {
            print( itos(i) + " buzz" + "\n" );
        } else {
            print( itos(i) + "\n" );
        }
    }
    return 0;
}`,
			nil,
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

			tops, err := Analyze("placeholder", nodes)
			if !assert.Equal(t, tt.expectErr, err) {
				t.Fatalf("failed to analyze: %v", err)
			}
			assert.Equal(t, tt.expectTops, tops)
		})
	}
}

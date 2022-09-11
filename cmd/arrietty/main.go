package main

import (
	"github.com/x0y14/arrietty/analyze"
	"github.com/x0y14/arrietty/interpret"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		log.Fatalf("too many args")
	}

	tokens, err := tokenize.Tokenize(args[0])
	if err != nil {
		log.Fatalf("failed to tokenize: %v", err)
	}

	nodes, err := parse.Parse(tokens)
	if err != nil {
		log.Fatalf("failed to parse: %v", err)
	}

	tops, err := analyze.Analyze(nodes)
	if err != nil {
		log.Fatalf("failed to analyze: %v", err)
	}

	r, err := interpret.Interpret(tops)
	if err != nil {
		log.Fatalf("failed to interpret: %v", err)
	}

	if r.Kind == interpret.OInt {
		os.Exit(r.I)
	}
}

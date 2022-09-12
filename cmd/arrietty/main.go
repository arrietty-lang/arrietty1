package main

import (
	"github.com/x0y14/arrietty/analyze"
	"github.com/x0y14/arrietty/interpret"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		log.Fatalf("too many args")
	}

	filepath := args[0]
	if !strings.HasSuffix(filepath, ".arr") {
		log.Fatalf("Please specify the path of the file ending with .arr")
	}

	src, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	tokens, err := tokenize.Tokenize(string(src))
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

	if r != nil && r.Kind == interpret.OInt {
		os.Exit(r.I)
	}
}

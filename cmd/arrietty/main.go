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

func readFile(filepath string) string {
	BUFSIZE := 4 * 1024

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]byte, 0, BUFSIZE)

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		data = append(data, buf...)
	}

	return strings.Replace(string(data), "\x00", "", -1)
}

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		log.Fatalf("too many args")
	}

	filepath := args[0]
	if !strings.HasSuffix(filepath, ".arr") {
		log.Fatalf("Please specify the path of the file ending with .arr")
	}

	src := readFile(filepath)

	tokens, err := tokenize.Tokenize(src)
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

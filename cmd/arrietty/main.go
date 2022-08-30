package main

import (
	"github.com/x0y14/arrietty/interpret"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"log"
	"os"
	"strings"
)

const (
	entrypoint = "main"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("need filepath")
	}

	for _, path := range args {
		if !strings.HasSuffix(path, ".arr") {
			log.Fatalf("pls set .arr file")
		}
		b, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read file: %v", path)
		}

		tok, err := tokenize.Tokenize(string(b))
		if err != nil {
			log.Fatalf("failed to tokenize: %v", err)
		}
		nod := parse.Parse(tok)
		err = interpret.Import(nod)
		if err != nil {
			log.Fatalf("failed to import node: %v", err)
		}
	}

	r, err := interpret.Run(entrypoint, nil)
	if err != nil {
		log.Fatalf("failed to run %s: %v", entrypoint, err)
	}
	//fmt.Printf("%v", r)
	os.Exit(r.NumInt)
}

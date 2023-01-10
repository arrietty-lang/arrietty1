package main

import (
	"github.com/x0y14/arrietty/analyze"
	"github.com/x0y14/arrietty/apm"
	"github.com/x0y14/arrietty/interpret"
	"github.com/x0y14/arrietty/parse"
	"github.com/x0y14/arrietty/tokenize"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatalf("too many args")
	}

	// args[0]は動作指定
	// args[1]はパラメータ

	switch args[0] {
	case "run":
		// プログラム実行

		entryArr := args[1]
		if !strings.HasSuffix(entryArr, ".arr") {
			log.Fatalf("Please specify the path of the file ending with .arr")
		}
		entryArrAbs, err := filepath.Abs(entryArr)
		if err != nil {
			log.Fatalf("failed to absoluting: %v", err)
		}
		pkgDirOfEntryArr := filepath.Dir(entryArrAbs)
		entryPkgInfo, err := apm.GetCurrentPackageInfo(pkgDirOfEntryArr)
		if err != nil {
			log.Fatalf("failed to read pkg.json: %v", err)
		}
		entryPkgArrs, err := apm.GetArrFilePathsInCurrent(pkgDirOfEntryArr)
		if err != nil {
			log.Fatalf("failed to get .arr files: %v", err)
		}

		// todo : check Is entryPackage's dependencies installed?
		// todo : pkg installer
		// todo : & Is same version?
		// todo : pkg updater

		tokens, err := tokenize.FromPaths(entryPkgArrs)
		if err != nil {
			log.Fatalf("failed to tokenize: %v", err)
		}

		syntaxTrees, err := parse.FromTokens(tokens)
		if err != nil {
			log.Fatalf("failed to parse: %v", err)
		}

		err = analyze.PkgAnalyze(entryPkgInfo.Name, syntaxTrees)
		if err != nil {
			log.Fatalf("failed t analyze: %v", err)
		}
		semanticsTrees := analyze.GetAnalyzedPackages()

		interpret.Setup()
		for pkgName, semTree := range semanticsTrees {
			err = interpret.LoadScript(pkgName, semTree)
			if err != nil {
				log.Fatalf("failed to load semanticsTree: %v", err)
			}
		}

		returnValue, err := interpret.Interpret(entryPkgInfo.Name, "main")
		if err != nil {
			log.Fatalf("failed to run function: %s.main: %v", entryPkgInfo.Name, err)
		}

		if returnValue != nil && returnValue.Kind == interpret.OInt {
			os.Exit(returnValue.I)
		}
	case "init":
		// パッケージ初期化
		root, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get working directory: %v", err)
		}
		err = apm.InitPackage(root, args[1])
		if err != nil {
			log.Fatalf("failed to initialize package: %v", err)
		}
	}
}

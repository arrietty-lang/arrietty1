package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

var currentPkg *PkgSymbols
var builtinPkg *PkgSymbols
var currentFunc *FunctionSymbol
var symbolTable *SymbolTable

var packages map[string]map[string]*TopLevel

func ResetSymbols() {
	currentPkg = nil
	currentFunc = nil
	symbolTable = &SymbolTable{}
	symbolTable.Packages = map[string]*PkgSymbols{}
	packages = map[string]map[string]*TopLevel{}
	attachBuiltin() // 付け直し
}

func init() {
	ResetSymbols()
}

func Analyze(syntaxTree []*parse.Node) (map[string]*TopLevel, error) {
	scripts := map[string]*TopLevel{}
	for _, n := range syntaxTree {
		top, err := NewToplevel(n)
		if err != nil {
			return nil, err
		}
		if top.Kind == TPFuncDef {
			scripts[top.FuncDef.Ident] = top
		}
	}

	return scripts, nil
}

func PkgAnalyze(pkgName string, syntaxTrees [][]*parse.Node) error {
	pkg, err := symbolTable.DeclarePkg(pkgName)
	if err != nil {
		return fmt.Errorf("failed to declare pkg: %v", err)
	}
	currentPkg = pkg
	packages[pkgName] = map[string]*TopLevel{}

	for _, tree := range syntaxTrees {
		functionsInFile, err := Analyze(tree)
		if err != nil {
			return err
		}
		for fnName, toplevel := range functionsInFile {
			packages[pkgName][fnName] = toplevel
		}
	}

	return nil
}

func GetAnalyzedPackages() map[string]map[string]*TopLevel {
	return packages
}

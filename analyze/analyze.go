package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

var currentPkg *PkgSymbols
var builtinPkg *PkgSymbols
var currentFunc *FunctionSymbol
var symbolTable *SymbolTable

func ResetSymbols() {
	currentPkg = nil
	currentFunc = nil
	symbolTable = &SymbolTable{}
	symbolTable.Packages = map[string]*PkgSymbols{}
	attachBuiltin() // 付け直し
}

func init() {
	currentPkg = nil
	currentFunc = nil
	symbolTable = &SymbolTable{}
	symbolTable.Packages = map[string]*PkgSymbols{}
	attachBuiltin()
}

func Analyze(pkgName string, nodes []*parse.Node) (map[string]*TopLevel, error) {
	pkg, err := symbolTable.DeclarePkg(pkgName)
	if err != nil {
		return nil, fmt.Errorf("failed to declare pkg: %v", err)
	}
	currentPkg = pkg

	scripts := map[string]*TopLevel{}
	for _, n := range nodes {
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

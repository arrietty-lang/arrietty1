package analyze

import "fmt"

// SymbolTable key is pkg name
type SymbolTable struct {
	// key is package name
	Packages map[string]*PkgSymbols
}

func (s *SymbolTable) IsDefinedPkg(pkgName string) (*PkgSymbols, bool) {
	v, ok := s.Packages[pkgName]
	return v, ok
}

func (s *SymbolTable) DeclarePkg(pkgName string) (*PkgSymbols, error) {
	_, ok := s.IsDefinedPkg(pkgName)
	if ok {
		return nil, fmt.Errorf("defined pkg: %s", pkgName)
	}
	s.Packages[pkgName] = &PkgSymbols{
		Functions: map[string]*FunctionSymbol{},
		Variables: map[string]*VariableSymbol{},
	}
	return s.Packages[pkgName], nil
}

type PkgSymbols struct {
	// key is function name
	Functions map[string]*FunctionSymbol
	// key is variable name
	Variables map[string]*VariableSymbol
}

func (p *PkgSymbols) IsDefinedFunc(functionName string) (*FunctionSymbol, bool) {
	v, ok := p.Functions[functionName]
	return v, ok
}

func (p *PkgSymbols) DeclareFunc(functionName string) (*FunctionSymbol, error) {
	if _, ok := p.IsDefinedFunc(functionName); ok {
		return nil, fmt.Errorf("defined func: %s", functionName)
	}
	if _, ok := p.IsDefinedVar(functionName); ok {
		return nil, fmt.Errorf("%s is defined as variable", functionName)
	}
	if _, ok := symbolTable.Packages["builtin"].IsDefinedFunc(functionName); ok {
		return nil, fmt.Errorf("%s is defined in builtin", functionName)
	}

	p.Functions[functionName] = &FunctionSymbol{
		Public:         false,
		Ident:          "",
		ReturnType:     nil,
		Params:         []*VariableSymbol{},
		LocalVariables: map[string]*VariableSymbol{},
	}
	return p.Functions[functionName], nil
}

func (p *PkgSymbols) IsDefinedVar(variableName string) (*VariableSymbol, bool) {
	v, ok := p.Variables[variableName]
	return v, ok
}

func (p *PkgSymbols) DeclareVar(variableName string) (*VariableSymbol, error) {
	if _, ok := p.IsDefinedVar(variableName); ok {
		return nil, fmt.Errorf("defined var: %s", variableName)
	}
	if _, ok := p.IsDefinedFunc(variableName); ok {
		return nil, fmt.Errorf("%s is defined as function", variableName)
	}
	if _, ok := symbolTable.Packages["builtin"].IsDefinedFunc(variableName); ok {
		return nil, fmt.Errorf("%s is defined in builtin", variableName)
	}

	p.Variables[variableName] = &VariableSymbol{}
	return p.Variables[variableName], nil
}

type FunctionSymbol struct {
	Public         bool
	Ident          string
	ReturnType     *DataType
	Params         []*VariableSymbol
	LocalVariables map[string]*VariableSymbol
}

func (f *FunctionSymbol) IsDefinedLocalVar(localVariableName string) (*VariableSymbol, bool) {
	v, ok := f.LocalVariables[localVariableName]
	return v, ok
}

func (f *FunctionSymbol) DeclareLocalVar(localVariableName string) (*VariableSymbol, error) {
	if _, ok := f.LocalVariables[localVariableName]; ok {
		return nil, fmt.Errorf("defined local var: %s", localVariableName)
	}
	f.LocalVariables[localVariableName] = &VariableSymbol{}
	return f.LocalVariables[localVariableName], nil
}

type VariableSymbol struct {
	Public   bool
	Ident    string
	DataType *DataType
}

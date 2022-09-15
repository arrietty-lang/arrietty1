package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

type RuntimeMemory struct {
	Packages map[string]*Memory
}

type Memory struct {
	Variables    map[string]*Object
	RawFunctions map[string]*analyze.TopLevel
}

func NewMemory(vars map[string]*Object, fs map[string]*analyze.TopLevel) *Memory {
	if vars == nil {
		vars = map[string]*Object{}
	}
	if fs == nil {
		fs = map[string]*analyze.TopLevel{}
	}

	return &Memory{
		Variables:    vars,
		RawFunctions: fs,
	}
}

func (m *Memory) GetFunc(ident string) (*analyze.FuncDef, error) {
	f, ok := m.RawFunctions[ident]
	if !ok {
		return nil, fmt.Errorf("function %s is not defined", ident)
	}
	return f.FuncDef, nil
}

func (m *Memory) DeclareVar(ident string) error {
	_, ok := m.Variables[ident]
	if ok {
		return fmt.Errorf("%s is already declared", ident)
	}
	m.Variables[ident] = nil
	return nil
}

func (m *Memory) AssignVar(ident string, obj *Object, allowOverwrite bool) error {
	v, ok := m.Variables[ident]
	// 宣言されていない
	if !ok {
		return fmt.Errorf("%s is not declared", ident)
	}

	// 宣言のみされているので書き込む
	if v == nil {
		m.Variables[ident] = obj
		return nil
	}

	// 値が入っているが許可されているので上書きする
	if v != nil && allowOverwrite {
		m.Variables[ident] = obj
		return nil
	}

	// 宣言されていて、値が代入されている
	return fmt.Errorf("can't assign obj to %s because allowOverwrite flag is false", ident)
}

func (m *Memory) GetVar(ident string) (*Object, error) {
	v, ok := m.Variables[ident]
	if !ok {
		return nil, fmt.Errorf("%s is not declared", ident)
	}
	return v, nil
}

func (m *Memory) IsVarDeclared(ident string) bool {
	_, ok := m.Variables[ident]
	return ok
}

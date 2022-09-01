package ast

type Function struct {
	RetVal ValueType
	Name   string
	Params []*FuncParam
	Stmts  []*Stmt
}

package ast

type FuncDef struct {
	RetType ValueType
	Name    string
	Params  []*FuncParam
}

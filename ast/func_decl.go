package ast

// impl extern?

type FuncDecl struct {
	RetType ValueType
	Name    string
	Params  []*FuncParam
}

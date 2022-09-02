package ast

type FuncParam struct {
	Name string
	Type ValueType
	Val  *Value
}

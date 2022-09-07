package analyze

import "github.com/x0y14/arrietty/parse"

type FuncParam struct {
	Ident string
	Type  *DataType
}

func NewFuncParam(node *parse.Node) (*FuncParam, error) {
	name := node.Lhs.S
	typ, err := NewDataType(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &FuncParam{
		Ident: name,
		Type:  typ,
	}, nil
}

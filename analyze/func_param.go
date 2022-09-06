package analyze

import "github.com/x0y14/arrietty/parse"

type FuncParam struct {
	Ident string
	Type  *ValueType
}

func NewFuncParam(node *parse.Node) (*FuncParam, error) {
	name := node.Lhs.S
	typ, err := NewValueType(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &FuncParam{
		Ident: name,
		Type:  typ,
	}, nil
}

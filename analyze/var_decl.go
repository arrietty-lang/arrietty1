package analyze

import "github.com/x0y14/arrietty/parse"

type VarDecl struct {
	Ident string
	Type  *DataType
}

func NewVarDecl(node *parse.Node) (*VarDecl, error) {
	identNode := node.Lhs
	ident := identNode.S

	// 型解析
	typ, err := NewDataTypeFromNode(node.Rhs)
	if err != nil {
		return nil, err
	}

	// シンボルテーブルに保存
	localVar, err := currentFunc.DeclareLocalVar(ident)
	localVar.Ident = ident
	localVar.DataType = typ

	return &VarDecl{
		Ident: ident,
		Type:  typ,
	}, nil
}

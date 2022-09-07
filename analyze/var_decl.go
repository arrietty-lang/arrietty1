package analyze

import "github.com/x0y14/arrietty/parse"

type VarDecl struct {
	Ident string
	Type  *DataType
}

func NewVarDecl(node *parse.Node) (*VarDecl, error) {

	identNode := node.Lhs
	ident := identNode.S

	if !isDefinableIdent(currentFunction, ident) {
		// 関数内あるいは、ルートの関数名としてすでに定義されてる
		return nil, NewAlreadyDefinedErr(currentFunction, ident)
	}

	typ, err := NewDataType(node.Rhs)
	if err != nil {
		return nil, err
	}

	// シンボルテーブルに保存
	err = defineVar(currentFunction, ident, typ)
	if err != nil {
		return nil, err
	}

	return &VarDecl{
		Ident: ident,
		Type:  typ,
	}, nil
}

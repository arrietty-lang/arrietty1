package analyze

import "github.com/x0y14/arrietty/parse"

type TopLevel struct { // toplevel
	Kind    ToplevelKind
	FuncDef *FuncDef
	Comment *Comment
}

func NewTopLevelFuncDef(node *parse.Node) (*TopLevel, error) {
	returnTypeNode := node.Children[0]
	nameNode := node.Children[1]
	paramsNode := node.Children[2]
	bodyNode := node.Children[3]

	if isFuncDefined(nameNode.S) {
		return nil, NewAlreadyDefinedErr("root", nameNode.S)
	}

	// scope
	decls[nameNode.S] = map[string]*ValueType{}

	def, err := NewFuncDef(returnTypeNode, nameNode, paramsNode, bodyNode)
	if err != nil {
		return nil, err
	}

	return &TopLevel{
		Kind:    TPFuncDef,
		FuncDef: def,
		Comment: nil,
	}, nil
}

func NewTopLevelComment(node *parse.Node) (*TopLevel, error) {
	return &TopLevel{
		Kind:    TPComment,
		FuncDef: nil,
		Comment: NewComment(node.S),
	}, nil
}

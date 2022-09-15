package analyze

import "github.com/x0y14/arrietty/parse"

type TopLevel struct {
	Kind    ToplevelKind
	FuncDef *FuncDef
	Comment *Comment
	Import  *Import
}

func NewToplevel(node *parse.Node) (*TopLevel, error) {
	switch node.Kind {
	case parse.FuncDef:
		return newTopLevelFuncDef(node)
	case parse.Comment:
		return newTopLevelComment(node)
	case parse.Import:
		return newTopLevelImport(node)
	}

	return nil, NewUnexpectNodeErr(node)
}

func newTopLevelFuncDef(node *parse.Node) (*TopLevel, error) {
	returnTypeNode := node.Children[0]
	nameNode := node.Children[1]
	paramsNode := node.Children[2]
	bodyNode := node.Children[3]

	_, yes := currentPkg.IsDefinedFunc(nameNode.S)
	if yes {
		return nil, NewAlreadyDefinedErr("file-toplevel", nameNode.S)
	}

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

func newTopLevelComment(node *parse.Node) (*TopLevel, error) {
	return &TopLevel{
		Kind:    TPComment,
		FuncDef: nil,
		Comment: NewComment(node.S),
	}, nil
}

func newTopLevelImport(node *parse.Node) (*TopLevel, error) {

	return &TopLevel{
		Kind:    TPImport,
		FuncDef: nil,
		Comment: nil,
		Import:  NewImport(node.S),
	}, nil
}

package analyze

import "github.com/x0y14/arrietty/parse"

type TopLevel struct { // toplevel
	Kind    ToplevelKind
	FuncDef *FuncDef
	Comment *Comment
}

func NewComment(node *parse.Node) *TopLevel {
	return &TopLevel{
		Kind:    TPComment,
		FuncDef: nil,
		Comment: &Comment{Value: node.S},
	}
}

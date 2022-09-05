package analyze

type SemanticNode struct { // toplevel
	Kind    SemanticNodeKind
	FuncDef *FuncDef
	Comment *Comment
}

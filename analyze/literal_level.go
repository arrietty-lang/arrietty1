package analyze

type LiteralLevel struct {
	Kind LiteralLevelKind

	HighPriority *ExprLevel

	Ident    string
	CallArgs []*ExprLevel

	Atom *Atom
	Dict *DictLevel
	List *ListLevel
}

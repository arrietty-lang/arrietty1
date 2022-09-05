package analyze

type ExprLevel struct {
	Kind        ExprLevelKind
	AssignLevel *AssignLevel
}

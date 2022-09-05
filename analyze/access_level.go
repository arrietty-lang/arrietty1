package analyze

type AccessLevel struct {
	Kind         AccessLevelKind
	LiteralLevel *LiteralLevel
	Index        *ExprLevel
}

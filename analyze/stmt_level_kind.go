package analyze

type StmtLevelKind int

const (
	_ StmtLevelKind = iota
	STExprLevel
	STReturn
	STIfElse
	STWhile
	STFor
	STComment
)

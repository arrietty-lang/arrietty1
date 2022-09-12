package analyze

type StmtLevelKind int

const (
	_ StmtLevelKind = iota
	STExprLevel
	STReturn
	STIfElse
	STWhile
	STFor
	STBlock
	STComment
)

var sKind = [...]string{
	STExprLevel: "STExprLevel",
	STReturn:    "STReturn",
	STIfElse:    "STIfElse",
	STWhile:     "STWhile",
	STFor:       "STFor",
	STBlock:     "STBlock",
	STComment:   "STComment",
}

func (s StmtLevelKind) String() string {
	return sKind[s]
}

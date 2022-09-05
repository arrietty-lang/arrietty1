package analyze

type LiteralLevelKind int

const (
	_ LiteralLevelKind = iota
	LParentheses
	LIdent
	LCall
	LAtom
)

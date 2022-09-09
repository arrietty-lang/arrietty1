package analyze

type LiteralLevelKind int

const (
	_ LiteralLevelKind = iota
	LParentheses
	LIdent
	LCall
	LAtom
	LList
	LDict
)

var lKind = [...]string{
	LParentheses: "LParentheses",
	LIdent:       "LIdent",
	LCall:        "LCall",
	LAtom:        "LAtom",
	LList:        "LList",
	LDict:        "LDict",
}

func (l LiteralLevelKind) String() string {
	return lKind[l]
}

package analyze

type ToplevelKind int

const (
	_ ToplevelKind = iota
	TPFuncDef
	TPComment
	TPImport
)

var toplevelKinds = [...]string{
	TPFuncDef: "FuncDef",
	TPComment: "Comment",
	TPImport:  "Import",
}

func (t ToplevelKind) String() string {
	return toplevelKinds[t]
}

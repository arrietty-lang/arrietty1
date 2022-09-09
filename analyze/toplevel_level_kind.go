package analyze

type ToplevelKind int

const (
	_ ToplevelKind = iota
	TPFuncDef
	TPComment
)

var toplevelKinds = [...]string{
	TPFuncDef: "FuncDef",
	TPComment: "Comment",
}

func (t ToplevelKind) String() string {
	return toplevelKinds[t]
}

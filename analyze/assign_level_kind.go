package analyze

type AssignLevelKind int

const (
	_ AssignLevelKind = iota
	AAndOrLevel
	AVarDecl // var x type
	AAssign  // to ident, to dict key, to list index
)

var aKind = [...]string{
	AAndOrLevel: "AAndOrLevel",
	AVarDecl:    "AVarDecl",
	AAssign:     "AAssign",
}

func (a AssignLevelKind) String() string {
	return aKind[a]
}

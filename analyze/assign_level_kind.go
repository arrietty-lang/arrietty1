package analyze

type AssignLevelKind int

const (
	_ AssignLevelKind = iota
	AAndOrLevel
	AVarDecl // var x type
	AAssign  // to ident, to dict key, to list index
)

package analyze

type AssignmentKind int

const (
	_ AssignmentKind = iota
	ToDefinedIdent
	ToUndefinedIdent
	ToDictKey
	ToListIndex
)

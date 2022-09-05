package analyze

type AssignKind int

const (
	_ AssignKind = iota
	ToIdent
	ToDictKey
	ToListIndex
)

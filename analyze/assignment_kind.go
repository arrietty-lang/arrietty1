package analyze

type AssignmentKind int

const (
	_ AssignmentKind = iota
	ToDefinedIdent
	ToUndefinedIdent
	ToDictKey
	ToListIndex
)

var assKind = [...]string{
	ToDefinedIdent:   "ToDefinedIdent",
	ToUndefinedIdent: "ToUndefinedIdent",
	ToDictKey:        "ToDictKey",
	ToListIndex:      "ToListIndex",
}

func (a AssignmentKind) String() string {
	return assKind[a]
}

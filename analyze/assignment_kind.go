package analyze

type AssignmentKind int

const (
	_ AssignmentKind = iota
	ToUnknown
	ToDefinedIdent
	ToUndefinedIdent
	ToDictKey
	ToListIndex
)

var assKind = [...]string{
	ToUnknown:        "ToUnknown",
	ToDefinedIdent:   "ToDefinedIdent",
	ToUndefinedIdent: "ToUndefinedIdent",
	ToDictKey:        "ToDictKey",
	ToListIndex:      "ToListIndex",
}

func (a AssignmentKind) String() string {
	return assKind[a]
}

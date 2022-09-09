package analyze

type AccessLevelKind int

const (
	_ AccessLevelKind = iota
	ACLiteralLevel
	ACDictIndex
	ACListIndex
	ACUnknownIndex
)

var acKind = [...]string{
	ACLiteralLevel: "ACLiteralLevel",
	ACDictIndex:    "ACDictIndex",
	ACListIndex:    "ACListIndex",
	ACUnknownIndex: "ACUnknownIndex",
}

func (a AccessLevelKind) String() string {
	return acKind[a]
}

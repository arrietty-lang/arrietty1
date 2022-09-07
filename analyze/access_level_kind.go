package analyze

type AccessLevelKind int

const (
	_ AccessLevelKind = iota
	ACLiteralLevel
	ACDictIndex
	ACListIndex
	ACUnknownIndex
)

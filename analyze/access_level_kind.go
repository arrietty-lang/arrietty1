package analyze

type AccessLevelKind int

const (
	_ AccessLevelKind = iota
	ACLiteralLevel
	ACDict
	ACList
)

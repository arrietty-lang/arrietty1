package analyze

type EqualityLevelKind int

const (
	_ EqualityLevelKind = iota
	EQRelationalLevel
	EQEqual
	EQNotEqual
)

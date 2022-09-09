package analyze

type EqualityLevelKind int

const (
	_ EqualityLevelKind = iota
	EQRelationalLevel
	EQEqual
	EQNotEqual
)

var eKind = [...]string{
	EQRelationalLevel: "EQRelationalLevel",
	EQEqual:           "EQEqual",
	EQNotEqual:        "EQNotEqual",
}

func (e EqualityLevelKind) String() string {
	return eKind[e]
}

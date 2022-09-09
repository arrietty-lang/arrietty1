package analyze

type AndOrLevelKind int

const (
	_ AndOrLevelKind = iota
	ANEqualityLevel
	ANAnd
	ANOr
)

var anKind = [...]string{
	ANEqualityLevel: "ANEqualityLevel",
	ANAnd:           "ANAnd",
	ANOr:            "ANOr",
}

func (a AndOrLevelKind) String() string {
	return anKind[a]
}

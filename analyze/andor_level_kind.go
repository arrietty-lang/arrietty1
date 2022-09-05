package analyze

type AndOrLevelKind int

const (
	_ AndOrLevelKind = iota
	ANEqualityLevel
	ANAnd
	ANOr
)

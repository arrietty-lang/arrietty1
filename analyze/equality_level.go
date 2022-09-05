package analyze

type EqualityLevel struct {
	Kind            EqualityLevelKind
	RelationalLevel *RelationalLevel
	LHS             *EqualityLevel
	RHS             *EqualityLevel
}

package analyze

type AndOrLevel struct {
	Kind     AndOrLevelKind
	Equality *EqualityLevel
	LHS      *AndOrLevel
	RHS      *AndOrLevel
}

package analyze

type RelationalLevel struct {
	Kind     RelationalLevelKind
	AddLevel *AddLevel
	LHS      *RelationalLevel
	RHS      *RelationalLevel
}

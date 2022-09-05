package analyze

type AddLevel struct {
	Kind     AndOrLevelKind
	MulLevel *MulLevel
	LHS      *AddLevel
	RHS      *AddLevel
}

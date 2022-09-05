package analyze

type MulLevel struct {
	Kind MulLevelKind

	UnaryLevel *UnaryLevel
	LHS        *MulLevel
	RHS        *MulLevel
}

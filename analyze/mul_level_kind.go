package analyze

type MulLevelKind int

const (
	_ MulLevelKind = iota
	MUUnaryLevel
	MUMul
	MUDiv
	MUMod
)

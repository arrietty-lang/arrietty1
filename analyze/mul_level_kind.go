package analyze

type MulLevelKind int

const (
	_ MulLevelKind = iota
	MUUnaryLevel
	MUMUl
	MUDiv
	MUMod
)

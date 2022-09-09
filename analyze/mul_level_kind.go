package analyze

type MulLevelKind int

const (
	_ MulLevelKind = iota
	MUUnaryLevel
	MUMul
	MUDiv
	MUMod
)

var mKind = [...]string{
	MUUnaryLevel: "MUUnaryLevel",
	MUMul:        "MUMul",
	MUDiv:        "MUDiv",
	MUMod:        "MUMod",
}

func (m MulLevelKind) String() string {
	return mKind[m]
}

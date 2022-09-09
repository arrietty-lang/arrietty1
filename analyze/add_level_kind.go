package analyze

type AddLevelKind int

const (
	_ AddLevelKind = iota
	ADMulLevel
	ADAdd
	ADSub
)

var adKind = [...]string{
	ADMulLevel: "ADMulLevel",
	ADAdd:      "ADAdd",
	ADSub:      "ADSub",
}

func (a AddLevelKind) String() string {
	return adKind[a]
}

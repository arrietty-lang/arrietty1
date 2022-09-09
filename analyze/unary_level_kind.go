package analyze

type UnaryLevelKind int

const (
	_ UnaryLevelKind = iota
	UNPlus
	UNMinus
	UNNot
	UNPrimaryLevel
)

var uKind = [...]string{
	UNPlus:         "UNPlus",
	UNMinus:        "UNMinus",
	UNNot:          "UNNot",
	UNPrimaryLevel: "UNPrimaryLevel",
}

func (u UnaryLevelKind) String() string {
	return uKind[u]
}

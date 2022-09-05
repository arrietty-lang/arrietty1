package analyze

type UnaryLevelKind int

const (
	_ UnaryLevelKind = iota
	UNPlus
	UNMinus
	UNNot
	UNPrimaryLevel
)

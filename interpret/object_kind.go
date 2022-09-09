package interpret

type ObjectKind int

const (
	_ ObjectKind = iota
	OFloat
	OInt
	OString
	OTrue
	OFalse
	ONull

	ODict
	OList
)

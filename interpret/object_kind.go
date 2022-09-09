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

var oKind = [...]string{
	OFloat:  "OFloat",
	OInt:    "OInt",
	OString: "OString",
	OTrue:   "OTrue",
	OFalse:  "OFalse",
	ONull:   "ONull",
	ODict:   "ODict",
	OList:   "OList",
}

func (o ObjectKind) String() string {
	return oKind[o]
}

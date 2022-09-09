package analyze

type AtomKind int

const (
	_ AtomKind = iota
	AFloat
	AInt
	AString
	ATrue
	AFalse
	ANull
)

var atomKind = [...]string{
	AFloat:  "AFloat",
	AInt:    "AInt",
	AString: "AString",
	ATrue:   "ATrue",
	AFalse:  "AFalse",
	ANull:   "ANull",
}

func (a AtomKind) String() string {
	return atomKind[a]
}

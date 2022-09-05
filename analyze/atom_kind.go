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

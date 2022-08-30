package interpret

type LiteralKind int

const (
	_ LiteralKind = iota
	Ident
	Float
	Int
	String
	Raw
	Array
	Dict
	True
	False
	Null
)

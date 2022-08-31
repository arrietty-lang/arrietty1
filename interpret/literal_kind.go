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

var literalKinds = [...]string{
	Ident:  "Ident",
	Float:  "Float",
	Int:    "Int",
	String: "String",
	Raw:    "Raw",
	Array:  "Array",
	Dict:   "Dict",
	True:   "True",
	False:  "False",
	Null:   "Null",
}

func (l LiteralKind) String() string {
	return literalKinds[l]
}

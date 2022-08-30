package interpret

type Literal struct {
	Kind     LiteralKind
	Str      string
	NumFloat float64
	NumInt   int

	Items []*Object
	KVS   map[string]*Object
}

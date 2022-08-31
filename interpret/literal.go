package interpret

type Literal struct {
	Kind     LiteralKind
	Str      string
	NumFloat float64
	NumInt   int

	Items []*Object
	KVS   map[string]*Object
}

func (l *Literal) IsFloat() bool {
	return l.Kind == Float
}

func (l *Literal) IsInt() bool {
	return l.Kind == Int
}

func (l *Literal) IsString() bool {
	return l.Kind == String
}

func (l *Literal) IsRaw() bool {
	return l.Kind == String
}

func (l *Literal) IsArray() bool {
	return l.Kind == Array
}

func (l *Literal) IsDict() bool {
	return l.Kind == Dict
}

func (l *Literal) IsBool() bool {
	return l.Kind == True || l.Kind == False
}

func (l *Literal) IsNull() bool {
	return l.Kind == Null
}

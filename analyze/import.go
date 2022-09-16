package analyze

type Import struct {
	Value string
}

func NewImport(c string) *Import {
	return &Import{Value: c}
}

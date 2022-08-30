package interpret

type ObjectKind int

const (
	_ ObjectKind = iota
	ObjLiteral
	ObjFn
)

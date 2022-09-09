package analyze

type ExprLevelKind int

const (
	_ ExprLevelKind = iota
	EXAssignLevel
)

var exKind = [...]string{
	EXAssignLevel: "EXAssignLevel",
}

func (e ExprLevelKind) String() string {
	return exKind[e]
}

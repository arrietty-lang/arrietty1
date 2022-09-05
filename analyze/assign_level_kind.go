package analyze

type AssignLevelKind int

const (
	_ AssignLevelKind = iota
	AAndOrLevel
	AVarDecl
	AAssign
)

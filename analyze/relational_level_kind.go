package analyze

type RelationalLevelKind int

const (
	_ RelationalLevelKind = iota
	REAddLevel
	RELt
	RELe
	REGt
	REGe
)

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

var rKind = [...]string{
	REAddLevel: "REAddLevel",
	RELt:       "RELt",
	RELe:       "RELe",
	REGt:       "REGt",
	REGe:       "REGe",
}

func (r RelationalLevelKind) String() string {
	return rKind[r]
}

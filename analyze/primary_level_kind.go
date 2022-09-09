package analyze

type PrimaryLevelKind int

const (
	_ PrimaryLevelKind = iota
	PRAccessLevel
)

var pKind = [...]string{
	PRAccessLevel: "PRAccessLevel",
}

func (p PrimaryLevelKind) String() string {
	return pKind[p]
}

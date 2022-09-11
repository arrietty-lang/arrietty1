package analyze

type TypeKind int

const (
	_ TypeKind = iota
	TFloat
	TInt
	TString
	TBool
	TVoid
	TDict
	TList
	TAny
)

package analyze

type Assign struct {
	Kind AssignKind

	VarDecl *VarDecl // dest
	// dict
	// list

	Value *AndOrLevel // src
}

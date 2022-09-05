package analyze

type AssignLevel struct {
	Kind       AssignLevelKind
	VarDecl    *VarDecl
	Assign     *Assign
	AndOrLevel *AndOrLevel
}

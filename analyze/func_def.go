package analyze

type FuncDef struct {
	ReturnType Types
	Ident      string
	Params     []*FuncParam
	Body       []*StmtLevel
}

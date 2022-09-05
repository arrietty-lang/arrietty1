package analyze

type While struct {
	Cond       *ExprLevel
	WhileBlock []*StmtLevel
}

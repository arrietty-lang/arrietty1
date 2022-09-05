package analyze

type IfElse struct {
	Cond      *ExprLevel
	IfBlock   []*StmtLevel
	ElseBlock []*StmtLevel
}

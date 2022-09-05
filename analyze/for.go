package analyze

type For struct {
	Init     *ExprLevel
	Cond     *ExprLevel
	Loop     *ExprLevel
	ForBlock []*StmtLevel
}

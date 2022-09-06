package analyze

type StmtLevel struct {
	Kind      StmtLevelKind
	ExprLevel *ExprLevel
	Return    *Return
	IfElse    *IfElse
	While     *While
	For       *For
	Comment   *Comment
}

func NewStmtLevelComment(c string) (*StmtLevel, error) {
	return &StmtLevel{Kind: STComment, Comment: NewComment(c)}, nil
}

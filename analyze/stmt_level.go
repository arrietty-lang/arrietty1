package analyze

import "github.com/x0y14/arrietty/parse"

type StmtLevel struct {
	Kind      StmtLevelKind
	ExprLevel *ExprLevel
	Return    *Return
	IfElse    *IfElse
	While     *While
	For       *For
	Comment   *Comment
}

func NewStmtLevel(node *parse.Node) (*StmtLevel, error) {
	switch node.Kind {
	case parse.Comment:
		return newStmtLevelComment(node)
	case parse.Return:
		return newStmtLevelReturn(node)
	case parse.If, parse.IfElse:
		return newStmtLevelIfElse(node)
	case parse.While:
		return newStmtLevelWhile(node)
	case parse.For:
		return newStmtLevelFor(node)
	}

	e, err := NewExprLevel(node)
	if err != nil {
		return nil, err
	}
	return &StmtLevel{Kind: STExprLevel, ExprLevel: e}, nil
}

func newStmtLevelComment(node *parse.Node) (*StmtLevel, error) {
	return &StmtLevel{Kind: STComment, Comment: NewComment(node.S)}, nil
}

func newStmtLevelReturn(node *parse.Node) (*StmtLevel, error) {
	r, err := NewReturn(node)
	if err != nil {
		return nil, err
	}

	return &StmtLevel{Kind: STReturn, Return: r}, nil
}
func newStmtLevelIfElse(node *parse.Node) (*StmtLevel, error) {
	i, err := NewIfElse(node)
	if err != nil {
		return nil, err
	}

	return &StmtLevel{Kind: STIfElse, IfElse: i}, nil
}
func newStmtLevelWhile(node *parse.Node) (*StmtLevel, error) {
	w, err := NewWhile(node)
	if err != nil {
		return nil, err
	}

	return &StmtLevel{Kind: STWhile, While: w}, nil
}
func newStmtLevelFor(node *parse.Node) (*StmtLevel, error) {
	f, err := NewFor(node)
	if err != nil {
		return nil, err
	}

	return &StmtLevel{Kind: STFor, For: f}, nil
}

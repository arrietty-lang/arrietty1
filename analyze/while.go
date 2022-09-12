package analyze

import "github.com/x0y14/arrietty/parse"

type While struct {
	Cond       *ExprLevel
	WhileBlock *StmtLevel
}

func NewWhile(node *parse.Node) (*While, error) {
	cond, err := NewExprLevel(node.Cond)
	if err != nil {
		return nil, err
	}

	//var while_ *StmtLevel = nil
	whileBlock := node.Children[0]
	while_, err := NewStmtLevel(whileBlock)
	if err != nil {
		return nil, err
	}

	return &While{
		Cond:       cond,
		WhileBlock: while_,
	}, nil
}

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

	var while_ *StmtLevel = nil
	whileBlock := node.Children[0]
	if whileBlock.Children != nil {
		w, err := newStmtLevelBlock(whileBlock)
		if err != nil {
			return nil, err
		}
		while_ = w
	}

	return &While{
		Cond:       cond,
		WhileBlock: while_,
	}, nil
}

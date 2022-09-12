package analyze

import "github.com/x0y14/arrietty/parse"

type For struct {
	Init     *ExprLevel
	Cond     *ExprLevel
	Loop     *ExprLevel
	ForBlock *StmtLevel
}

func NewFor(node *parse.Node) (*For, error) {

	var init *ExprLevel = nil
	var cond *ExprLevel = nil
	var loop *ExprLevel = nil

	if node.Init != nil {
		i, err := NewExprLevel(node.Init)
		if err != nil {
			return nil, err
		}
		init = i
	}

	if node.Cond != nil {
		c, err := NewExprLevel(node.Cond)
		if err != nil {
			return nil, err
		}
		cond = c
	}

	if node.Loop != nil {
		l, err := NewExprLevel(node.Loop)
		if err != nil {
			return nil, err
		}
		loop = l
	}

	forBlock := node.Children[0]
	for_, err := NewStmtLevel(forBlock)
	if err != nil {
		return nil, err
	}

	return &For{Init: init, Cond: cond, Loop: loop, ForBlock: for_}, nil
}

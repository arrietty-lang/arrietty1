package analyze

import "github.com/x0y14/arrietty/parse"

type For struct {
	Init     *ExprLevel
	Cond     *ExprLevel
	Loop     *ExprLevel
	ForBlock []*StmtLevel
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

	var fors []*StmtLevel = nil
	forBlock := node.Children[0]
	if forBlock.Children != nil {
		for _, forStmt := range forBlock.Children {
			stmt, err := NewStmtLevel(forStmt)
			if err != nil {
				return nil, err
			}
			fors = append(fors, stmt)
		}
	}

	return &For{Init: init, Cond: cond, Loop: loop, ForBlock: fors}, nil
}

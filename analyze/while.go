package analyze

import "github.com/x0y14/arrietty/parse"

type While struct {
	Cond       *ExprLevel
	WhileBlock []*StmtLevel
}

func NewWhile(node *parse.Node) (*While, error) {
	cond, err := NewExprLevel(node.Cond)
	if err != nil {
		return nil, err
	}

	var whiles []*StmtLevel = nil
	whileBlock := node.Children[0]
	if node.Children != nil {
		for _, whileStmt := range whileBlock.Children {
			stmt, err := NewStmtLevel(whileStmt)
			if err != nil {
				return nil, err
			}
			whiles = append(whiles, stmt)
		}
	}

	return &While{
		Cond:       cond,
		WhileBlock: whiles,
	}, nil
}

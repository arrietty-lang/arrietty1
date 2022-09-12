package analyze

import "github.com/x0y14/arrietty/parse"

type Block struct {
	Body []*StmtLevel
}

func NewBlock(node *parse.Node) (*Block, error) {
	var stmts []*StmtLevel = nil
	if node.Children != nil {
		for _, stmtNode := range node.Children {
			stmt, err := NewStmtLevel(stmtNode)
			if err != nil {
				return nil, err
			}
			stmts = append(stmts, stmt)
		}
	}
	return &Block{Body: stmts}, nil
}

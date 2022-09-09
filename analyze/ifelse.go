package analyze

import "github.com/x0y14/arrietty/parse"

type IfElse struct {
	Cond      *ExprLevel
	IfBlock   []*StmtLevel
	ElseBlock []*StmtLevel
}

func NewIfElse(node *parse.Node) (*IfElse, error) {
	cond, err := NewExprLevel(node.Cond)
	if err != nil {
		return nil, err
	}

	var ifs []*StmtLevel = nil
	ifBlock := node.Children[0]
	if ifBlock.Children != nil {
		for _, ifBlockStmt := range ifBlock.Children {
			stmt, err := NewStmtLevel(ifBlockStmt)
			if err != nil {
				return nil, err
			}
			ifs = append(ifs, stmt)
		}
	}
	// elseを解析せずに返す
	if node.Kind == parse.If {
		return &IfElse{
			Cond:      cond,
			IfBlock:   ifs,
			ElseBlock: nil,
		}, nil
	}

	var elses []*StmtLevel = nil
	elseBlock := node.Children[1]
	if elseBlock.Children != nil {
		for _, elseBlockStmt := range elseBlock.Children {
			stmt, err := NewStmtLevel(elseBlockStmt)
			if err != nil {
				return nil, err
			}
			elses = append(elses, stmt)
		}
	}

	return &IfElse{Cond: cond, IfBlock: ifs, ElseBlock: elses}, nil
}

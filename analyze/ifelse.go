package analyze

import "github.com/x0y14/arrietty/parse"

type IfElse struct {
	Cond      *ExprLevel
	IfBlock   *StmtLevel
	ElseBlock *StmtLevel
}

func NewIfElse(node *parse.Node) (*IfElse, error) {
	cond, err := NewExprLevel(node.Cond)
	if err != nil {
		return nil, err
	}

	ifBlock := node.Children[0]
	if_, err := newStmtLevelBlock(ifBlock)
	if err != nil {
		return nil, err
	}
	// elseを解析せずに返す
	if node.Kind == parse.If {
		return &IfElse{
			Cond:      cond,
			IfBlock:   if_,
			ElseBlock: nil,
		}, nil
	}

	var else_ *StmtLevel = nil
	elseBlock := node.Children[1]
	if elseBlock.Children != nil {
		e, err := newStmtLevelBlock(elseBlock)
		if err != nil {
			return nil, err
		}
		else_ = e
	}

	return &IfElse{Cond: cond, IfBlock: if_, ElseBlock: else_}, nil
}

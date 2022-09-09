package analyze

import "github.com/x0y14/arrietty/parse"

type ExprLevel struct {
	Kind        ExprLevelKind
	AssignLevel *AssignLevel
}

func (e *ExprLevel) GetType() (*DataType, error) {
	return e.AssignLevel.GetType()
}

func NewExprLevel(node *parse.Node) (*ExprLevel, error) {
	a, err := NewAssignLevel(node)
	if err != nil {
		return nil, err
	}
	return &ExprLevel{Kind: EXAssignLevel, AssignLevel: a}, nil
}

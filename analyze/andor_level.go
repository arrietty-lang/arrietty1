package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type AndOrLevel struct {
	Kind          AndOrLevelKind
	EqualityLevel *EqualityLevel
	LHS           *AndOrLevel
	RHS           *AndOrLevel
}

func (a *AndOrLevel) GetType() (*DataType, error) {
	switch a.Kind {
	case ANAnd, ANOr:
		return &DataType{Type: TBool}, nil
	case ANEqualityLevel:
		return a.EqualityLevel.GetType()
	}
	return nil, fmt.Errorf("andorLv type error")
}

func NewAndOrLevel(node *parse.Node) (*AndOrLevel, error) {
	switch node.Kind {
	case parse.And:
		return newAndOrLevelAnd(node)
	case parse.Or:
		return newAndOrLevelOr(node)
	}

	e, err := NewEqualityLevel(node)
	if err != nil {
		return nil, err
	}
	return &AndOrLevel{Kind: ANEqualityLevel, EqualityLevel: e}, nil
}

func newAndOrLevelAnd(node *parse.Node) (*AndOrLevel, error) {
	lhs, err := NewAndOrLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewAndOrLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &AndOrLevel{Kind: ANAnd, LHS: lhs, RHS: rhs}, nil
}

func newAndOrLevelOr(node *parse.Node) (*AndOrLevel, error) {
	lhs, err := NewAndOrLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewAndOrLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &AndOrLevel{Kind: ANOr, LHS: lhs, RHS: rhs}, nil
}

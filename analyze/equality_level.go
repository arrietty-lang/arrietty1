package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type EqualityLevel struct {
	Kind            EqualityLevelKind
	RelationalLevel *RelationalLevel
	LHS             *EqualityLevel
	RHS             *EqualityLevel
}

func (e *EqualityLevel) GetType() (*DataType, error) {
	switch e.Kind {
	case EQEqual, EQNotEqual:
		return &DataType{Type: TBool}, nil
	case EQRelationalLevel:
		return e.RelationalLevel.GetType()
	}
	return nil, fmt.Errorf("equalityLv type error")
}

func NewEqualityLevel(node *parse.Node) (*EqualityLevel, error) {
	switch node.Kind {
	case parse.Eq:
		return newEqualityLevelEq(node)
	case parse.Ne:
		return newEqualityLevelNe(node)
	}

	r, err := NewRelationalLevel(node)
	if err != nil {
		return nil, err
	}
	return &EqualityLevel{Kind: EQRelationalLevel, RelationalLevel: r}, nil
}

func newEqualityLevelEq(node *parse.Node) (*EqualityLevel, error) {
	lhs, err := NewEqualityLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewEqualityLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	lhsType, err := lhs.GetType()
	if err != nil {
		return nil, err
	}
	rhsType, err := rhs.GetType()
	if err != nil {
		return nil, err
	}
	if !isSameType(lhsType, rhsType) {
		return nil, fmt.Errorf("type miss match L:%s, R:%s", lhsType.String(), rhsType.String())
	}

	return &EqualityLevel{Kind: EQEqual, LHS: lhs, RHS: rhs}, nil
}

func newEqualityLevelNe(node *parse.Node) (*EqualityLevel, error) {
	lhs, err := NewEqualityLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewEqualityLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	lhsType, err := lhs.GetType()
	if err != nil {
		return nil, err
	}
	rhsType, err := rhs.GetType()
	if err != nil {
		return nil, err
	}
	if !isSameType(lhsType, rhsType) {
		return nil, fmt.Errorf("type miss match L:%s, R:%s", lhsType.String(), rhsType.String())
	}

	return &EqualityLevel{Kind: EQNotEqual, LHS: lhs, RHS: rhs}, nil
}

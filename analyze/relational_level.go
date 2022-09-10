package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type RelationalLevel struct {
	Kind     RelationalLevelKind
	AddLevel *AddLevel
	LHS      *RelationalLevel
	RHS      *RelationalLevel
}

func (r *RelationalLevel) GetType() (*DataType, error) {
	switch r.Kind {
	case RELt, RELe, REGt, REGe:
		return &DataType{Type: TBool}, nil
	case REAddLevel:
		return r.AddLevel.GetType()
	}
	return nil, fmt.Errorf("relationalLv type error")
}

func NewRelationalLevel(node *parse.Node) (*RelationalLevel, error) {
	switch node.Kind {
	case parse.Lt:
		return newRelationalLevelLt(node)
	case parse.Le:
		return newRelationalLevelLe(node)
	case parse.Gt:
		return newRelationalLevelGt(node)
	case parse.Ge:
		return newRelationalLevelGe(node)
	}

	a, err := NewAddLevel(node)
	if err != nil {
		return nil, err
	}
	return &RelationalLevel{Kind: REAddLevel, AddLevel: a}, nil
}

func newRelationalLevelLt(node *parse.Node) (*RelationalLevel, error) {
	lhs, err := NewRelationalLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewRelationalLevel(node.Rhs)
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

	return &RelationalLevel{Kind: RELt, LHS: lhs, RHS: rhs}, nil
}

func newRelationalLevelLe(node *parse.Node) (*RelationalLevel, error) {
	lhs, err := NewRelationalLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewRelationalLevel(node.Rhs)
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

	return &RelationalLevel{Kind: RELe, LHS: lhs, RHS: rhs}, nil
}

func newRelationalLevelGt(node *parse.Node) (*RelationalLevel, error) {
	lhs, err := NewRelationalLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewRelationalLevel(node.Rhs)
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

	return &RelationalLevel{Kind: REGt, LHS: lhs, RHS: rhs}, nil
}

func newRelationalLevelGe(node *parse.Node) (*RelationalLevel, error) {
	lhs, err := NewRelationalLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewRelationalLevel(node.Rhs)
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

	return &RelationalLevel{Kind: REGe, LHS: lhs, RHS: rhs}, nil
}

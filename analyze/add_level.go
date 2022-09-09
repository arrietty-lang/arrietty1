package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type AddLevel struct {
	Kind     AddLevelKind
	MulLevel *MulLevel
	LHS      *AddLevel
	RHS      *AddLevel
}

func (a *AddLevel) GetType() (*DataType, error) {
	switch a.Kind {
	case ADMulLevel:
		return a.MulLevel.GetType()
	case ADAdd, ADSub:
		lhs, err := a.LHS.GetType()
		if err != nil {
			return nil, err
		}
		rhs, err := a.RHS.GetType()
		if err != nil {
			return nil, err
		}

		// 片方がfloatだった場合、もう片方がintでもfloatでもfloatになる
		if lhs.Type == TFloat && (rhs.Type == TFloat || rhs.Type == TInt) {
			return lhs, nil
		}
		if (lhs.Type == TFloat || lhs.Type == TInt) && rhs.Type == TFloat {
			return rhs, nil
		}
		// どっちもIntなら
		if lhs.Type == TInt && rhs.Type == TInt {
			return lhs, nil
		}
	}
	return nil, fmt.Errorf("add type error")
}

func NewAddLevel(node *parse.Node) (*AddLevel, error) {
	switch node.Kind {
	case parse.Add:
		return newAddLevelAdd(node)
	case parse.Sub:
		return newAddLevelSub(node)
	}

	m, err := NewMulLevel(node)
	if err != nil {
		return nil, err
	}
	return &AddLevel{Kind: ADMulLevel, MulLevel: m}, nil
}

func newAddLevelAdd(node *parse.Node) (*AddLevel, error) {
	lhs, err := NewAddLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewAddLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &AddLevel{Kind: ADAdd, LHS: lhs, RHS: rhs}, nil
}

func newAddLevelSub(node *parse.Node) (*AddLevel, error) {
	lhs, err := NewAddLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewAddLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &AddLevel{Kind: ADSub, LHS: lhs, RHS: rhs}, nil
}

package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type MulLevel struct {
	Kind       MulLevelKind
	UnaryLevel *UnaryLevel
	LHS        *MulLevel
	RHS        *MulLevel
}

func (m *MulLevel) GetType() (*DataType, error) {
	switch m.Kind {
	case MUUnaryLevel:
		return m.UnaryLevel.GetType()
	case MUMul, MUDiv, MUMod:
		lhs, err := m.LHS.GetType()
		if err != nil {
			return nil, err
		}
		rhs, err := m.RHS.GetType()
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

	return nil, fmt.Errorf("mulLv is type error")
}

func NewMulLevel(node *parse.Node) (*MulLevel, error) {
	switch node.Kind {
	case parse.Mul:
		return newMulLevelMul(node)
	case parse.Div:
		return newMulLevelDiv(node)
	case parse.Mod:
		return newMulLevelMod(node)
	}

	u, err := NewUnaryLevel(node)
	if err != nil {
		return nil, err
	}
	return &MulLevel{Kind: MUUnaryLevel, UnaryLevel: u}, nil
}

func newMulLevelMul(node *parse.Node) (*MulLevel, error) {
	lhs, err := NewMulLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewMulLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &MulLevel{Kind: MUMul, LHS: lhs, RHS: rhs}, nil
}

func newMulLevelDiv(node *parse.Node) (*MulLevel, error) {
	lhs, err := NewMulLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewMulLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &MulLevel{Kind: MUDiv, LHS: lhs, RHS: rhs}, nil
}

func newMulLevelMod(node *parse.Node) (*MulLevel, error) {
	lhs, err := NewMulLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	rhs, err := NewMulLevel(node.Rhs)
	if err != nil {
		return nil, err
	}

	return &MulLevel{Kind: MUMod, LHS: lhs, RHS: rhs}, nil
}

package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type UnaryLevel struct {
	Kind         UnaryLevelKind
	PrimaryLevel *PrimaryLevel
}

func (u *UnaryLevel) GetType() (*DataType, error) {
	switch u.Kind {
	case UNPrimaryLevel:
		return u.PrimaryLevel.GetType()
	case UNPlus:
		// int, float
		// "+X"と記述できる型
		pr, err := u.PrimaryLevel.GetType()
		if err != nil {
			return nil, err
		}
		if pr.Type == TFloat || pr.Type == TInt {
			return pr, nil
		}
	case UNMinus:
		// int, float
		// "-X"と記述できる型
		pr, err := u.PrimaryLevel.GetType()
		if err != nil {
			return nil, err
		}
		if pr.Type == TFloat || pr.Type == TInt {
			return pr, nil
		}
	case UNNot:
		// bool
		// "!X"と記述できる型
		pr, err := u.PrimaryLevel.GetType()
		if err != nil {
			return nil, err
		}
		if pr.Type == TBool {
			return pr, nil
		}
	}
	return nil, fmt.Errorf("unary type error")
}

func NewUnaryLevel(node *parse.Node) (*UnaryLevel, error) {
	switch node.Kind {
	case parse.Plus:
		return newUnaryLevelPlus(node)
	case parse.Minus:
		return newUnaryLevelMinus(node)
	case parse.Not:
		return newUnaryLevelNot(node)
	}

	p, err := NewPrimaryLevel(node)
	if err != nil {
		return nil, err
	}
	return &UnaryLevel{Kind: UNPrimaryLevel, PrimaryLevel: p}, nil
}

func newUnaryLevelPlus(node *parse.Node) (*UnaryLevel, error) {
	p, err := NewPrimaryLevel(node.Lhs)
	if err != nil {
		return nil, err
	}
	return &UnaryLevel{Kind: UNPlus, PrimaryLevel: p}, nil
}

func newUnaryLevelMinus(node *parse.Node) (*UnaryLevel, error) {
	p, err := NewPrimaryLevel(node.Lhs)
	if err != nil {
		return nil, err
	}
	return &UnaryLevel{Kind: UNMinus, PrimaryLevel: p}, nil
}

func newUnaryLevelNot(node *parse.Node) (*UnaryLevel, error) {
	p, err := NewPrimaryLevel(node.Lhs)
	if err != nil {
		return nil, err
	}
	return &UnaryLevel{Kind: UNNot, PrimaryLevel: p}, nil
}

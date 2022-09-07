package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type PrimaryLevel struct {
	Kind        PrimaryLevelKind
	AccessLevel *AccessLevel
}

func (p *PrimaryLevel) GetType() (*DataType, error) {
	switch p.Kind {
	case PRAccessLevel:
		return p.AccessLevel.GetType()
	}

	panic(fmt.Sprintf("primaryLevel %d is not support getType", p.Kind))
}

func NewPrimaryLevel(node *parse.Node) (*PrimaryLevel, error) {
	switch node.Kind {
	}

	a, err := NewAccessLevel(node)
	if err != nil {
		return nil, err
	}
	return &PrimaryLevel{Kind: PRAccessLevel, AccessLevel: a}, nil
}

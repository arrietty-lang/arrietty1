package analyze

import "github.com/x0y14/arrietty/parse"

type KVLevel struct {
	Key   string
	Value *UnaryLevel
}

func NewKVLevel(node *parse.Node) (*KVLevel, error) {
	if node.Kind != parse.KV {
		return nil, NewUnexpectNodeErr(node)
	}

	key := node.Lhs.S
	valueNode := node.Rhs

	value, err := NewUnaryLevel(valueNode)
	if err != nil {
		return nil, err
	}

	return &KVLevel{Key: key, Value: value}, nil
}

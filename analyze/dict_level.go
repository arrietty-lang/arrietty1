package analyze

import "github.com/x0y14/arrietty/parse"

type DictLevel struct {
	KVs []*KVLevel
}

func NewDictLevel(node *parse.Node) (*DictLevel, error) {
	var kvs []*KVLevel

	for _, kvNode := range node.Children {
		kv, err := NewKVLevel(kvNode)
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, kv)
	}

	return &DictLevel{KVs: kvs}, nil
}

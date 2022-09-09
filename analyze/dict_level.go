package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type DictLevel struct {
	KVs []*KVLevel
}

func (d *DictLevel) IsEmpty() bool {
	return len(d.KVs) == 0
}

func (d *DictLevel) GetType() (*DataType, error) {
	t := &DataType{Type: TDict}

	// 空
	if d.IsEmpty() {
		return t, nil
	}

	// 現時点ではstring固定
	// todo : key、string以外の型に対応
	baseKeyType := &DataType{Type: TString}
	baseValueType, err := d.KVs[0].Value.GetType()
	if err != nil {
		return nil, err
	}
	for _, kv := range d.KVs {
		vt, err := kv.Value.GetType()
		if err != nil {
			return nil, err
		}
		if !isSameType(baseValueType, vt) {
			return nil, fmt.Errorf("dict.values contains different types")
		}
	}
	t.Key = baseKeyType
	t.Value = baseValueType

	return t, nil
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

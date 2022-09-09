package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type ListLevel struct {
	Items []*UnaryLevel
}

func (l *ListLevel) IsEmpty() bool {
	return len(l.Items) == 0
}

func (l *ListLevel) GetType() (*DataType, error) {
	t := &DataType{Type: TList}

	// シンプルに数えることでサイズを取得
	t.Size = len(l.Items)
	if l.IsEmpty() {
		return t, nil
	}

	baseType, err := l.Items[0].GetType()
	if err != nil {
		return nil, err
	}

	// 型anyは導入していないので、Itemsから異なる方が取得できた瞬間にエラー
	for _, item := range l.Items {
		itemType, err := item.GetType()
		if err != nil {
			return nil, err
		}
		if !isSameType(baseType, itemType) {
			return nil, fmt.Errorf("list contains different types")
		}
	}
	t.Item = baseType

	return t, nil
}

func NewListLevel(node *parse.Node) (*ListLevel, error) {
	var items []*UnaryLevel
	for _, itemNode := range node.Children {
		item, err := NewUnaryLevel(itemNode)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &ListLevel{Items: items}, nil
}

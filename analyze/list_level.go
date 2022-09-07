package analyze

import "github.com/x0y14/arrietty/parse"

type ListLevel struct {
	Size  int
	Items []*UnaryLevel
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

	return &ListLevel{Size: len(items), Items: items}, nil
}

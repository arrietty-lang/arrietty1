package analyze

import "github.com/x0y14/arrietty/parse"

func Analyze(nodes []*parse.Node) ([]*TopLevel, error) {
	var trees []*TopLevel
	for _, n := range nodes {
		top, err := NewToplevel(n)
		if err != nil {
			return nil, err
		}
		trees = append(trees, top)
	}
	return trees, nil
}

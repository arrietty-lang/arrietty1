package analyze

import "github.com/x0y14/arrietty/parse"

func analyzeNode(node *parse.Node) *SemanticNode {

	return nil
}

func Analyze(nodes []*parse.Node) []*SemanticNode {
	var trees []*SemanticNode
	for _, n := range nodes {
		trees = append(trees, analyzeNode(n))
	}
	return trees
}

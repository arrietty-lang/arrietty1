package interpret

import "github.com/x0y14/arrietty/parse"

func if_(scope *Storage, node *parse.Node) (*Object, error) {
	condObj, err := eval(scope, node.Cond)
	if err != nil {
		return nil, err
	}

	if condObj.Literal.Kind == True {
		return eval(scope, node.Children[0])
	}

	return nil, nil
}

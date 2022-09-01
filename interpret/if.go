package interpret

import "github.com/x0y14/arrietty/parse"

func if_(scope *Storage, node *parse.Node) (*Object, error) {
	condObj, err := eval(scope, node.Cond)
	if err != nil {
		return nil, err
	}

	ifBlock := node.Children[0]
	if condObj.Literal.Kind == True {
		return eval(scope, ifBlock)
	}

	return nil, nil
}

func ifElse(scope *Storage, node *parse.Node) (*Object, error) {
	condObj, err := eval(scope, node.Cond)
	if err != nil {
		return nil, err
	}

	ifBlock := node.Children[0]
	elseBlock := node.Children[1]

	if condObj.Literal.Kind == True {
		return eval(scope, ifBlock)
	} else {
		return eval(scope, elseBlock)
	}
}

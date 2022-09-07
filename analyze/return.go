package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type Return struct {
	Value *ExprLevel
}

func NewReturn(node *parse.Node) (*Return, error) {
	retType, _ := isDefinedFunc(currentFunction)

	if node.Children == nil {
		// voidではないのに、戻り値が提供されていない
		if retType.Type != TVoid {
			return nil, fmt.Errorf("return value is not provided")
		}
		return nil, nil
	}

	actualRetValue, err := NewExprLevel(node.Children[0])
	if err != nil {
		return nil, err
	}
	t, err := actualRetValue.GetType()
	if err != nil {
		return nil, err
	}

	if !isSameType(retType, t) {
		return nil, fmt.Errorf("return type miss match want:%s, found:%s", retType.String(), t.String())
	}

	return &Return{Value: actualRetValue}, nil
}

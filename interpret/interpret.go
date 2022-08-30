package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

var globalStorage *Storage

func init() {
	globalStorage = NewStorage()
}

func eval(scope *Storage, node *parse.Node) (*Object, error) {
	switch node.Kind {
	case parse.Ident:
		return Load(scope, node.Str)
	case parse.Float:
		return NewFloat(node.NumFloat), nil
	case parse.Int:
		return NewInt(node.NumInt), nil
	case parse.String:
		return NewString(node.Str), nil
	case parse.Raw:
		return NewRaw(node.Str), nil
	case parse.Array:
		return NewArray(node.Children)
	case parse.Dict:
		return NewDict(node.Children)
	case parse.Function:
		id := node.Children[0]
		params := node.Children[1]
		body := node.Children[2]
		f := NewFn(params.Children, body)
		err := Store(globalStorage, id.Str, f)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case parse.Return:
		return eval(scope, node.Children[0])
	}
	return nil, nil
}

func Interpret(nodes []*parse.Node) (*Object, error) {
	for _, node := range nodes {
		v, err := eval(globalStorage, node)
		if err != nil {
			return nil, err
		}
		if v != nil {
			fmt.Println(v.String())
		}
	}

	main, err := Load(globalStorage, "main")
	if err != nil {
		return nil, fmt.Errorf("need main")
	}

	return main.Exec(nil)
}

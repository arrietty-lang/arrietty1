package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
	"log"
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
		return NewArray(scope, node.Children)
	case parse.Dict:
		return NewDict(scope, node.Children)
	case parse.True:
		return NewTrue(), nil
	case parse.False:
		return NewFalse(), nil
	case parse.Null:
		return NewNull(), nil
	case parse.Add:
		return add(scope, node)
	case parse.Sub:
		return sub(scope, node)
	case parse.Mul:
		return mul(scope, node)
	case parse.Div:
		return div(scope, node)
	case parse.Mod:
		return mod(scope, node)
	case parse.Eq:
		return eq(scope, node)
	case parse.Ne:
		return ne(scope, node)
	case parse.Lt:
		return lt(scope, node)
	case parse.Le:
		return le(scope, node)
	case parse.Gt:
		return gt(scope, node)
	case parse.Ge:
		return ge(scope, node)
	case parse.Assign:
		// 保存する値
		rhs, err := eval(scope, node.Rhs)
		if err != nil {
			return nil, err
		}
		if node.Lhs.Kind == parse.Access {
			// 格納先本体の解決
			ident := node.Lhs.Children[0]
			identObj, err := eval(scope, ident)
			if err != nil {
				return nil, err
			}
			// 添字の解決
			index := node.Lhs.Children[1]
			indexObj, err := eval(scope, index)
			if err != nil {
				return nil, err
			}
			// dict
			if identObj.Literal.Kind == Dict {
				identObj.KVS[indexObj.Str] = rhs
			} else if identObj.Literal.Kind == Array {
				identObj.Items[indexObj.NumInt] = rhs
			} else {
				log.Fatalf("assign unsupport: %d", identObj.Literal.Kind)
			}

		} else if node.Lhs.Kind == parse.Ident {
			err = Store(scope, node.Lhs.Str, rhs)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
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
	case parse.Call:
		id := node.Children[0]
		args := node.Children[1]
		f, err := Load(globalStorage, id.Str)
		if err != nil {
			return nil, err
		}
		return f.Exec(scope, args.Children)
	case parse.Access:
		id := node.Children[0]
		dest := node.Children[1]
		d, err := eval(scope, dest)
		if err != nil {
			return nil, err
		}
		obj, err := eval(scope, id)
		if err != nil {
			return nil, err
		}
		if obj.Literal.Kind == Dict {
			return obj.KVS[d.Str], nil
		} else {
			return obj.Items[d.NumInt], nil
		}
	}
	return nil, nil
}

func Import(nodes []*parse.Node) error {
	for _, node := range nodes {
		_, err := eval(globalStorage, node)
		if err != nil {
			return err
		}
	}
	return nil
}

func Run(fName string, args []*parse.Node) (*Object, error) {
	f, err := Load(globalStorage, fName)
	if err != nil {
		return nil, err
	}
	return f.Exec(globalStorage, args)
}

func Interpret(nodes []*parse.Node) (*Object, error) {
	for _, node := range nodes {
		_, err := eval(globalStorage, node)
		if err != nil {
			return nil, err
		}
		//if v != nil {
		//	fmt.Println(v.String())
		//}
	}

	main, err := Load(globalStorage, "main")
	if err != nil {
		return nil, fmt.Errorf("need main")
	}

	return main.Exec(globalStorage, nil)
}

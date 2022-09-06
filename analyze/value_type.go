package analyze

import "github.com/x0y14/arrietty/parse"

type ValueType struct {
	Type  TypeKind
	Size  *ExprLevel
	Item  *ValueType
	Key   *ValueType
	Value *ValueType
}

func NewValueType(node *parse.Node) (*ValueType, error) {
	switch node.Kind {
	case parse.Float:
		return &ValueType{Type: TFloat}, nil
	case parse.Int:
		return &ValueType{Type: TInt}, nil
	case parse.String:
		return &ValueType{Type: TString}, nil
	case parse.Bool:
		return &ValueType{Type: TBool}, nil
	case parse.Void:
		return &ValueType{Type: TVoid}, nil
	case parse.Dict:
		keyType, err := NewValueType(node.Children[0])
		if err != nil {
			return nil, err
		}
		valueType, err := NewValueType(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &ValueType{Type: TDict, Key: keyType, Value: valueType}, nil
	case parse.List:
		var size *ExprLevel = nil
		if node.Children != nil {
			// todo
			// s, err = NewExpr(node.children[0]
			// if err
			// size = s
		}
		itemType, err := NewValueType(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &ValueType{Type: TList, Size: size, Item: itemType}, nil
	}

	return nil, NewUnsupportedTypeErr(node)
}

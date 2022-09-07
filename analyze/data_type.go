package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type DataType struct {
	Type  TypeKind
	Size  *ExprLevel
	Item  *DataType
	Key   *DataType
	Value *DataType
}

func isSameType(d1, d2 *DataType) bool {
	if d1.Type != d2.Type {
		return false
	}

	// todo
	// list, dict
	if d1.Type == TList || d1.Type == TDict ||
		d2.Type == TList || d2.Type == TDict {
		return false
	}

	return true
}

func (d *DataType) String() string {
	switch d.Type {
	case TBool:
		return "bool"
	case TFloat:
		return "float"
	case TInt:
		return "int"
	case TString:
		return "string"
	case TVoid:
		return "void"
	case TDict:
		return fmt.Sprintf("dict[%s]%s", d.Key.String(), d.Value.String())
	case TList:
		return fmt.Sprintf("[?]%s", d.Item.String())
	}
	return "unknown"
}

func NewDataType(node *parse.Node) (*DataType, error) {
	switch node.Kind {
	case parse.Float:
		return &DataType{Type: TFloat}, nil
	case parse.Int:
		return &DataType{Type: TInt}, nil
	case parse.String:
		return &DataType{Type: TString}, nil
	case parse.Bool:
		return &DataType{Type: TBool}, nil
	case parse.Void:
		return &DataType{Type: TVoid}, nil
	case parse.Dict:
		keyType, err := NewDataType(node.Children[0])
		if err != nil {
			return nil, err
		}
		valueType, err := NewDataType(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &DataType{Type: TDict, Key: keyType, Value: valueType}, nil
	case parse.List:
		var size *ExprLevel = nil
		if node.Children != nil {
			// todo
			// s, err = NewExpr(node.children[0]
			// if err
			// size = s
		}
		itemType, err := NewDataType(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &DataType{Type: TList, Size: size, Item: itemType}, nil
	}

	return nil, NewUnsupportedTypeErr(node)
}

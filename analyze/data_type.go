package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type DataType struct {
	Type  TypeKind
	Size  int
	Item  *DataType
	Key   *DataType
	Value *DataType
}

// 代入が可能かを確認する関数
func isAssignable(destination *DataType, t *DataType) bool {
	if destination.Type == TAny {
		return true
	}
	// そもそも一致しない場合は弾く
	if destination.Type != t.Type {
		return false
	}

	switch t.Type {
	case TList: // listは空(item=nil)、大きさの問題があるので確認をする
		// []int = []のようにからだった場合、値から推論ができないためItem=nilになる.この場合、型にとらわれないため、代入を可能として良い
		if t.Item == nil {
			return true
		}
		// listの中の型が一致していることを確認(厳密でなくて良い)
		//   -> 再帰的に確認される可能性があるので、厳密なチェックをかけると多分ほとんど通らなくなる
		if !isAssignable(destination.Item, t.Item) {
			return false
		}
		// 要素数を確認
		// 格納先が動的なサイズを許容していたら、入れるものの大きさは考慮しなくて良い
		if destination.Size == -1 {
			return true
		}
		// オーバーしていれば弾く
		if destination.Size < t.Size {
			return false
		}
		return true
	case TDict: // dictは空(value=nil)の可能性があるので確認する
		// dict[string]int = {}のような場合、左辺から値を推論できないためValue=nilになる
		if t.Value == nil {
			return true
		}
		// keyの型をチェック
		if !isAssignable(destination.Key, t.Key) {
			return false
		}
		// 値に厳密でない型チェックをかける
		if !isAssignable(destination.Value, t.Value) {
			return false
		}
		return true
	}

	// 特殊な型でないのであれば一致を確認する
	return isSameType(destination, t)
}

// 厳密に型が一致するかをチェックする関数
func isSameType(d1, d2 *DataType) bool {
	if d1.Type != d2.Type {
		return false
	}

	switch d1.Type {
	case TDict:
		// 空のdictは全てのdictと型が一致する(ということにする)
		if d1.Value == nil || d2.Value == nil {
			return true
		}
		return isSameType(d1.Key, d2.Key) && isSameType(d1.Value, d2.Value)
	case TList:
		// 空のlistは全てのlistと型が一致する(ということにする)
		if d1.Item == nil || d2.Item == nil {
			return true
		}
		return isSameType(d1.Item, d2.Item) && d1.Size == d2.Size
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
		return fmt.Sprintf("[%d]%s", d.Size, d.Item.String())
	case TAny:
		return "any"
	}
	return "unknown"
}

func NewDataTypeFromNode(node *parse.Node) (*DataType, error) {
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
		keyType, err := NewDataTypeFromNode(node.Children[0])
		if err != nil {
			return nil, err
		}
		valueType, err := NewDataTypeFromNode(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &DataType{Type: TDict, Key: keyType, Value: valueType}, nil
	case parse.List:
		// listの型定義は
		// children = [size, type]ってなってる
		size := -1
		if node.Children[0] != nil {
			size = node.Children[0].I
		}
		if size == 0 {
			return nil, fmt.Errorf("zero size list declaration is not support")
		}
		itemType, err := NewDataTypeFromNode(node.Children[1])
		if err != nil {
			return nil, err
		}
		return &DataType{Type: TList, Size: size, Item: itemType}, nil
	case parse.Any:
		return &DataType{Type: TAny}, nil
	}

	return nil, NewUnsupportedTypeErr(node)
}

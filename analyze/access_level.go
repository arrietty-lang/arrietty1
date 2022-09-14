package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
	"log"
)

type AccessLevel struct {
	Kind         AccessLevelKind
	LiteralLevel *LiteralLevel

	Src   *AccessLevel
	Index *ExprLevel
}

func (a *AccessLevel) GetType() (*DataType, error) {
	switch a.Kind {
	case ACLiteralLevel:
		return a.LiteralLevel.GetType()
	case ACListIndex:
		src, err := a.Src.GetType()
		if err != nil {
			return nil, err
		}
		return src.Item, nil
	case ACDictIndex:
		src, err := a.Src.GetType()
		if err != nil {
			return nil, err
		}
		return src.Value, nil
	case ACUnknownIndex:
		return a.Src.GetType()
	}
	return nil, fmt.Errorf("accessLv %d is not support getType", a.Kind)
}

func NewAccessLevel(node *parse.Node) (*AccessLevel, error) {
	switch node.Kind {
	case parse.Access:
		return newAccessLevelIndex(node)
	}

	l, err := NewLiteralLevel(node)
	if err != nil {
		return nil, err
	}
	return &AccessLevel{Kind: ACLiteralLevel, LiteralLevel: l}, nil
}

func newAccessLevelIndex(node *parse.Node) (*AccessLevel, error) {
	srcNode := node.Lhs
	indexNode := node.Rhs

	src, err := NewAccessLevel(srcNode)
	if err != nil {
		return nil, err
	}

	index, err := NewExprLevel(indexNode)
	if err != nil {
		return nil, err
	}
	indexType, err := index.GetType()
	if err != nil {
		return nil, err
	}

	var indexKind = ACUnknownIndex

	// 左辺をidentとして読み取ることができたら、シンボルテーブルから型データを取り出す
	if src.Kind == ACLiteralLevel && src.LiteralLevel.Kind == LIdent {
		varSymbol, ok := currentFunc.IsDefinedLocalVar(src.LiteralLevel.Ident)
		if ok {
			if varSymbol.DataType.Type == TList {
				if indexType.Type != TInt {
					return nil, fmt.Errorf("list index expect int, but got %s", indexType.String())
				}
				indexKind = ACListIndex
			} else if varSymbol.DataType.Type == TDict {
				if !isSameType(varSymbol.DataType.Key, indexType) {
					return nil, fmt.Errorf("dict key type is not match: expect %s, but got %s", varSymbol.DataType.Key.String(), indexType.String())
				}
				indexKind = ACDictIndex
			}
		}
	}

	// dict, listならアクセスlv経由でGetTypeすれば、型が取得できるはず
	if src.Kind == ACDictIndex || src.Kind == ACListIndex {
		t, err := src.GetType()
		if err != nil {
			return nil, err
		}
		if t.Type == TList {
			if indexType.Type != TInt {
				return nil, fmt.Errorf("list index expect int, but got %s", indexType.String())
			}
			indexKind = ACListIndex
		} else if t.Type == TDict {
			if !isSameType(indexType, t.Key) {
				return nil, fmt.Errorf("dict key type is not match: expect %s, but got %s", t.Key.String(), indexType.String())
			}
			indexKind = ACDictIndex
		} else {
			log.Fatalf("unsupported access ")
		}
	}

	return &AccessLevel{Kind: indexKind, Src: src, Index: index}, nil
}

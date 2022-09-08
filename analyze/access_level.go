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

	var indexType = ACUnknownIndex

	// 左辺をidentとして読み取ることができたら、シンボルテーブルから型データを取り出す
	if src.Kind == ACLiteralLevel && src.LiteralLevel.Kind == LIdent {
		typ, ok := isDefinedVariable(currentFunction, src.LiteralLevel.Ident)
		if ok {
			if typ.Type == TList {
				indexType = ACListIndex
			} else if typ.Type == TDict {
				indexType = ACDictIndex
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
			indexType = ACListIndex
		} else if t.Type == TDict {
			indexType = ACDictIndex
		} else {
			log.Fatalf("unsupported access ")
		}
	}

	return &AccessLevel{Kind: indexType, Src: src, Index: index}, nil
}

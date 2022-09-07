package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
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
	case ACDictIndex, ACListIndex:
		// 判明しているのであれば、シンボルテーブルからより詳細なデータを得られるはず
		t, ok := isDefinedVariable(currentFunction, a.LiteralLevel.Ident)
		if !ok {
			return nil, fmt.Errorf("can't call getType from undefined ident")
		}
		return t, nil
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
	if src.LiteralLevel.Kind == LIdent {
		typ, ok := isDefinedVariable(currentFunction, src.LiteralLevel.Ident)
		if ok {
			if typ.Type == TList {
				indexType = ACListIndex
			} else if typ.Type == TDict {
				indexType = ACDictIndex
			}
		}
	}

	return &AccessLevel{Kind: indexType, Src: src, Index: index}, nil
}

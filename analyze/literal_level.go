package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type LiteralLevel struct {
	Kind LiteralLevelKind

	ExprLevel *ExprLevel // for parentheses

	Ident    string
	CallArgs []*ExprLevel

	Atom      *Atom
	DictLevel *DictLevel
	ListLevel *ListLevel
}

func (l *LiteralLevel) GetType() (*DataType, error) {
	switch l.Kind {
	case LParentheses:
		return l.ExprLevel.GetType()
	case LIdent:
		v, ok := isDefinedVariable(currentFunction, l.Ident)
		if !ok {
			// シンボルテーブルに定義されていない変数から型情報を引き出そうとした
			return nil, NewUndefinedErr(l.Ident)
		}
		return v, nil
	case LCall:
		t, ok := isDefinedFunc(l.Ident)
		if !ok {
			// 同じく定義されていない関数から戻り値を取得しようとした
			return nil, NewUndefinedErr(l.Ident)
		}
		return t, nil
	case LAtom:
		return l.Atom.GetType()
	case LList:
		return l.ListLevel.GetType()
	case LDict:
		return l.DictLevel.GetType()
	}
	return nil, fmt.Errorf("literalLv %d is not support getType", l.Kind)
}

func NewLiteralLevel(node *parse.Node) (*LiteralLevel, error) {
	switch node.Kind {
	case parse.Parenthesis:
		return newLiteralLevelParentheses(node)
	case parse.Ident:
		return newLiteralLevelIdent(node)
	case parse.Call:
		return newLiteralLevelCall(node)
	case parse.Float, parse.Int,
		parse.String, parse.RawString,
		parse.True, parse.False,
		parse.Null:
		return newLiteralLevelAtom(node)
	case parse.List:
		return newLiteralLevelList(node)
	case parse.Dict:
		return newLiteralLevelDict(node)
	}

	return nil, NewUnexpectNodeErr(node)
}

func newLiteralLevelParentheses(node *parse.Node) (*LiteralLevel, error) {
	if node.Kind != parse.Parenthesis {
		return nil, NewUnexpectNodeErr(node)
	}

	expr, err := NewExprLevel(node.Lhs)
	if err != nil {
		return nil, err
	}

	return &LiteralLevel{Kind: LParentheses, ExprLevel: expr}, nil
}

func newLiteralLevelAtom(node *parse.Node) (*LiteralLevel, error) {
	a, err := NewAtom(node)
	if err != nil {
		return nil, err
	}
	return &LiteralLevel{Kind: LAtom, Atom: a}, nil
}

func newLiteralLevelIdent(node *parse.Node) (*LiteralLevel, error) {
	if node.Kind != parse.Ident {
		return nil, NewUnexpectNodeErr(node)
	}

	ident := node.S

	return &LiteralLevel{Kind: LIdent, Ident: ident}, nil
}

func newLiteralLevelCall(node *parse.Node) (*LiteralLevel, error) {
	argsNode := node.Children[1]
	if argsNode != nil && argsNode.Kind != parse.Args {
		return nil, NewUnexpectNodeErr(argsNode)
	}

	identNode := node.Children[0]
	ident := identNode.S
	if argsNode == nil {
		return &LiteralLevel{Kind: LCall, Ident: ident, CallArgs: nil}, nil
	}

	var args []*ExprLevel
	for _, argNode := range argsNode.Children {
		arg, err := NewExprLevel(argNode)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	return &LiteralLevel{Kind: LCall, Ident: ident, CallArgs: args}, nil
}

func newLiteralLevelList(node *parse.Node) (*LiteralLevel, error) {
	l, err := NewListLevel(node)
	if err != nil {
		return nil, err
	}

	return &LiteralLevel{Kind: LList, ListLevel: l}, nil
}

func newLiteralLevelDict(node *parse.Node) (*LiteralLevel, error) {
	d, err := NewDictLevel(node)
	if err != nil {
		return nil, err
	}

	return &LiteralLevel{Kind: LDict, DictLevel: d}, nil
}

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
		identType, ok := currentFunc.IsDefinedLocalVar(l.Ident)
		if !ok {
			// シンボルテーブルに定義されていない変数から型情報を引き出そうとした
			return nil, NewUndefinedErr(l.Ident)
		}
		return identType.DataType, nil
	case LCall:
		funcDecl, ok := currentPkg.IsDefinedFunc(l.Ident)
		if !ok {
			f, ok := builtinPkg.IsDefinedFunc(l.Ident)
			if !ok {

				// 定義されていない関数から戻り値を取得しようとした
				return nil, NewUndefinedErr(l.Ident)
			}
			funcDecl = f
		}
		return funcDecl.ReturnType, nil
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
	functionName := identNode.S

	funcDecl, ok := currentPkg.IsDefinedFunc(functionName)
	if !ok {
		f, ok := builtinPkg.IsDefinedFunc(functionName)
		if !ok {
			return nil, fmt.Errorf("the function you tried to call is not defined: %s", functionName)
		}
		funcDecl = f
	}

	functionParams := funcDecl.Params

	// 引数の個数が一致するかの確認
	// 呼び出し時に渡された引数がnil(0個)だった場合
	if argsNode == nil {
		if len(functionParams) != 0 {
			return nil, fmt.Errorf("%s call param error want: %d, reserve: %d", functionName, len(functionParams), 0)
		}
		return &LiteralLevel{Kind: LCall, Ident: functionName, CallArgs: nil}, nil
	}
	if len(functionParams) != len(argsNode.Children) {
		return nil, fmt.Errorf("%s call param error want: %d, reserve: %d", functionName, len(functionParams), len(argsNode.Children))
	}

	// 引数を前から順番に型を検証していく
	var args []*ExprLevel
	for i, argNode := range argsNode.Children {
		arg, err := NewExprLevel(argNode)
		if err != nil {
			return nil, err
		}

		// 引数の型を取得
		argType, err := arg.GetType()
		if err != nil {
			return nil, fmt.Errorf("filed to getType : params.%d of function %s", i, functionName)
		}

		// 型検証
		paramData := funcDecl.Params[i]
		if !isAssignable(paramData.DataType, argType) {
			return nil, fmt.Errorf("%s's param: %s, reserve invalid type arg. want: %s, reserve: %s", functionName, paramData.Ident, paramData.DataType.String(), argType.String())
		}

		args = append(args, arg)
	}

	return &LiteralLevel{Kind: LCall, Ident: functionName, CallArgs: args}, nil
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

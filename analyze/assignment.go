package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type Assignment struct {
	Kind     AssignmentKind
	Ident    string
	AccessLv *AccessLevel

	Value *AndOrLevel
}

func NewAssignment(node *parse.Node) (*Assignment, error) {
	// lhs: var decl or ident or access
	// rhs: andor

	if node.Kind == parse.ShortVarDecl {
		// ident := RHS
		ident := node.Lhs.S
		// すでに定義されている
		// 宣言+割り当てなので宣言されてはならない
		if !isDefinableIdent(currentFunction, ident) {
			return nil, NewAlreadyDefinedErr(currentFunction, ident)
		}

		// 右辺代入するデータを解析、型を取得
		val, err := NewAndOrLevel(node.Rhs)
		if err != nil {
			return nil, err
		}

		t, err := val.GetType()
		if err != nil {
			return nil, err
		}
		// シンボルテーブルに記録
		err = defineVar(currentFunction, ident, t)
		if err != nil {
			return nil, err
		}

		return &Assignment{Kind: ToDefinedIdent, Ident: ident, Value: val}, nil
	}

	switch node.Lhs.Kind {
	case parse.VarDecl:
		// var ident type = RHS

		// 左辺の型を解析
		decl, err := NewVarDecl(node.Lhs)
		if err != nil {
			return nil, err
		}

		// 右辺代入するデータを解析、型を取得
		val, err := NewAndOrLevel(node.Rhs)
		if err != nil {
			return nil, err
		}
		t, err := val.GetType()
		if err != nil {
			return nil, err
		}

		// 右辺左辺の一致を確認
		if isSameType(decl.Type, t) {
			return nil, fmt.Errorf("type miss match L:%s  R:%s", decl.Type.String(), t.String())
		}

		return &Assignment{Kind: ToDefinedIdent, Ident: decl.Ident, Value: val}, nil

	case parse.Ident:
		// ident = RHS
		// シンボルテーブルに記録されているはず
		ident := node.Lhs.S
		identType, ok := isDefinedVariable(currentFunction, ident)
		if !ok {
			return nil, NewUndefinedErr(ident)
		}

		// 右辺代入するデータを解析、型を取得
		val, err := NewAndOrLevel(node.Rhs)
		if err != nil {
			return nil, err
		}
		valueType, err := val.GetType()
		if err != nil {
			return nil, err
		}

		// 両辺の型が一致していることを確認
		if !isSameType(identType, valueType) {
			return nil, fmt.Errorf("assign type miss match L:%s, R:%s", identType.String(), valueType.String())
		}
		return &Assignment{Kind: ToDefinedIdent, Ident: ident, Value: val}, nil
	case parse.Access:
		// list[N] = RHS
		// dict[KEY] = RHS
		// todo
	}

	return nil, fmt.Errorf("unimplemented list, dict")
}

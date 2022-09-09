package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type Assignment struct {
	Kind     AssignmentKind
	Ident    string
	AccessLv *AccessLevel
	Inline   bool // 宣言と代入が１行で行われている時にそのことをインタプリタに教えるために使う

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
		// 長すぎw
		if t.Type == TList && val.EqualityLevel.RelationalLevel.AddLevel.MulLevel.UnaryLevel.PrimaryLevel.AccessLevel.LiteralLevel.ListLevel.IsEmpty() {
			return nil, fmt.Errorf("can't assign empty-list to ident in short-variable-declaration")
		}
		if t.Type == TDict && val.EqualityLevel.RelationalLevel.AddLevel.MulLevel.UnaryLevel.PrimaryLevel.AccessLevel.LiteralLevel.DictLevel.IsEmpty() {
			return nil, fmt.Errorf("can't assign empty-dict to ident in short-variable-declaration")
		}
		// シンボルテーブルに記録
		err = defineVar(currentFunction, ident, t)
		if err != nil {
			return nil, err
		}

		return &Assignment{Kind: ToDefinedIdent, Ident: ident, Value: val, Inline: true}, nil
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
		if !isAssignable(decl.Type, t) {
			return nil, fmt.Errorf("[assign(to var-decl)] type miss match L:%s  R:%s", decl.Type.String(), t.String())
		}

		return &Assignment{Kind: ToDefinedIdent, Ident: decl.Ident, Value: val, Inline: true}, nil

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
		if !isAssignable(identType, valueType) {
			return nil, fmt.Errorf("[assign(to ident)] type miss match L:%s, R:%s", identType.String(), valueType.String())
		}
		return &Assignment{Kind: ToDefinedIdent, Ident: ident, Value: val}, nil
	case parse.Access:

		var assignKind AssignmentKind = ToUnknown

		acc, err := NewAccessLevel(node.Lhs)
		if err != nil {
			return nil, err
		}

		destinationType, err := acc.GetType()
		if err != nil {
			return nil, err
		}
		if acc.Kind == ACListIndex {
			assignKind = ToListIndex
		} else if acc.Kind == ACDictIndex {
			assignKind = ToDictKey
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

		if !isAssignable(destinationType, valueType) {
			return nil, fmt.Errorf("can't assign %s to %s", valueType.String(), destinationType.String())
		}

		return &Assignment{Kind: assignKind, AccessLv: acc, Value: val}, nil

		//var destination AssignmentKind
		//// 入れる場所と、入れるものの型が一致しているかを確認するだけ
		//
		//accessNode := node.Lhs
		//base := accessNode.Lhs
		//
		//// ident[?] = ?を扱える
		//if base.Kind == parse.Ident {
		//	// identをいただいているので、型をとってくる
		//
		//	acc, err := NewAccessLevel(accessNode)
		//	if err != nil {
		//		return nil, err
		//	}
		//
		//	accType, err := acc.GetType()
		//	if err != nil {
		//		return nil, err
		//	}
		//	if accType.Type == TDict {
		//		destination = ToDictKey
		//	} else if accType.Type == TList {
		//		destination = ToListIndex
		//	}
		//
		//	// 右辺代入するデータを解析、型を取得
		//	val, err := NewAndOrLevel(node.Rhs)
		//	if err != nil {
		//		return nil, err
		//	}
		//	valueType, err := val.GetType()
		//	if err != nil {
		//		return nil, err
		//	}
		//	if !isAssignable(accType, valueType) {
		//		return nil, fmt.Errorf("can't assign %s to %s", valueType.String(), accType.String())
		//	}
		//
		//	return &Assignment{Kind: destination, Ident: base.S, Value: val}, nil
		//}
		//
		//return NewAssignment(node.Lhs)

	}

	return nil, fmt.Errorf("unsupport assign")
}

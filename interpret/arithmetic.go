package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

func add(scope *Storage, node *parse.Node) (*Object, error) {
	// 左辺評価
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	// 右辺評価
	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}

	// どちらもIntならば
	if lhsObj.Literal.IsInt() && rhsObj.Literal.IsInt() {
		return NewInt(lhsObj.NumInt + rhsObj.NumInt), nil
	}
	// どちらもFloatならば
	if lhsObj.Literal.IsFloat() && rhsObj.Literal.IsFloat() {
		return NewFloat(lhsObj.NumFloat + rhsObj.NumFloat), nil
	}
	// L:F, R:I
	if lhsObj.Literal.IsFloat() && rhsObj.Literal.IsInt() {
		return NewFloat(lhsObj.NumFloat + float64(rhsObj.NumInt)), nil
	}
	// L:I, R:F
	if lhsObj.Literal.IsInt() && rhsObj.Literal.IsFloat() {
		return NewFloat(float64(lhsObj.NumInt) + rhsObj.Literal.NumFloat), nil
	}
	// どちらもStringならば
	if lhsObj.IsString() && rhsObj.IsString() {
		return NewString(lhsObj.Literal.Str + rhsObj.Literal.Str), nil
	}

	// todo : raw

	return nil, fmt.Errorf("unsupported combination: %s + %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

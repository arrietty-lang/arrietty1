package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

func isSupportedByEqNe(kind LiteralKind) bool {
	for _, k := range []LiteralKind{Float, Int, String, Raw, True, False, Null} {
		if k == kind {
			return true
		}
	}
	return false
}

func eq(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByEqNe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by eq: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByEqNe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by eq: %s", rhsObj.Literal.Kind.String())
	}

	// -> 1.0 != 1
	// 検討
	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case True, False, Null:
			return NewTrue(), nil
		case String, Raw:
			if lhsObj.Literal.Str == rhsObj.Literal.Str {
				return NewTrue(), nil
			}
			return NewFalse(), nil
		case Float:
			if lhsObj.Literal.NumFloat == rhsObj.Literal.NumFloat {
				return NewTrue(), nil
			}
			return NewFalse(), nil
		case Int:
			if lhsObj.NumInt == rhsObj.NumInt {
				return NewTrue(), nil
			}
			return NewFalse(), nil
		default:
			return NewFalse(), nil
		}
	} else {
		if lhsObj.Literal.Kind == True && rhsObj.Literal.Kind == False {
			return NewFalse(), nil
		}
		if lhsObj.Literal.Kind == False && rhsObj.Literal.Kind == True {
			return NewFalse(), nil
		}
	}
	return nil, fmt.Errorf("type miss match: %s == %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

func ne(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByEqNe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by ne: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByEqNe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by ne: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case True, False, Null:
			return NewFalse(), nil
		case String, Raw:
			if lhsObj.Literal.Str == rhsObj.Literal.Str {
				return NewFalse(), nil
			}
			return NewTrue(), nil
		case Float:
			if lhsObj.Literal.NumFloat == rhsObj.Literal.NumFloat {
				return NewFalse(), nil
			}
			return NewTrue(), nil
		case Int:
			if lhsObj.NumInt == rhsObj.NumInt {
				return NewFalse(), nil
			}
			return NewTrue(), nil
		}
	} else {
		if lhsObj.Literal.Kind == True && rhsObj.Literal.Kind == False {
			return NewTrue(), nil
		}
		if lhsObj.Literal.Kind == False && rhsObj.Literal.Kind == True {
			return NewTrue(), nil
		}
	}
	return nil, fmt.Errorf("type miss match: %s != %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

func isSupportedByLtLeGtGe(kind LiteralKind) bool {
	for _, k := range []LiteralKind{Float, Int} {
		if k == kind {
			return true
		}
	}
	return false
}

func lt(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by lt: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by lt: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case Float:
			if lhsObj.Literal.NumFloat < rhsObj.Literal.NumFloat {
				return NewTrue(), nil
			}
		case Int:
			if lhsObj.Literal.NumInt < rhsObj.Literal.NumInt {
				return NewTrue(), nil
			}
		}
		return NewFalse(), nil
	}

	return nil, fmt.Errorf("type miss match: %s < %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

func le(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by le: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by le: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case Float:
			if lhsObj.Literal.NumFloat <= rhsObj.Literal.NumFloat {
				return NewTrue(), nil
			}
		case Int:
			if lhsObj.Literal.NumInt <= rhsObj.Literal.NumInt {
				return NewTrue(), nil
			}
		}
		return NewFalse(), nil
	}

	return nil, fmt.Errorf("type miss match: %s <= %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

func gt(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by gt: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by gt: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case Float:
			if lhsObj.Literal.NumFloat > rhsObj.Literal.NumFloat {
				return NewTrue(), nil
			}
		case Int:
			if lhsObj.Literal.NumInt > rhsObj.Literal.NumInt {
				return NewTrue(), nil
			}
		}
		return NewFalse(), nil
	}

	return nil, fmt.Errorf("type miss match: %s > %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

func ge(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by ge: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isSupportedByLtLeGtGe(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by ge: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		switch lhsObj.Literal.Kind {
		case Float:
			if lhsObj.Literal.NumFloat >= rhsObj.Literal.NumFloat {
				return NewTrue(), nil
			}
		case Int:
			if lhsObj.Literal.NumInt >= rhsObj.Literal.NumInt {
				return NewTrue(), nil
			}
		}
		return NewFalse(), nil
	}

	return nil, fmt.Errorf("type miss match: %s >= %s",
		lhsObj.Literal.Kind.String(),
		rhsObj.Literal.Kind.String())
}

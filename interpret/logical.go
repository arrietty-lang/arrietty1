package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

func isBool(kind LiteralKind) bool {
	switch kind {
	case True, False:
		return true
	}
	return false
}

func and(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isBool(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by and: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isBool(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by and: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		// どちらもTrue
		if lhsObj.Literal.Kind == True {
			return NewTrue(), nil
		}
		// どちらもFalse
		if lhsObj.Literal.Kind == False {
			return NewFalse(), nil
		}
	}

	return NewFalse(), nil
}

func or(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isBool(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by or: %s", lhsObj.Literal.Kind.String())
	}

	rhsObj, err := eval(scope, node.Rhs)
	if err != nil {
		return nil, err
	}
	if !isBool(rhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by or: %s", rhsObj.Literal.Kind.String())
	}

	if lhsObj.Literal.Kind == rhsObj.Literal.Kind {
		// どちらもTrue
		if lhsObj.Literal.Kind == True {
			return NewTrue(), nil
		}
		// どちらもFalse
		if lhsObj.Literal.Kind == False {
			return NewFalse(), nil
		}
	}

	return NewTrue(), nil
}

func not(scope *Storage, node *parse.Node) (*Object, error) {
	lhsObj, err := eval(scope, node.Lhs)
	if err != nil {
		return nil, err
	}
	if !isBool(lhsObj.Literal.Kind) {
		return nil, fmt.Errorf("unsupported by not: %s", lhsObj.Literal.Kind.String())
	}

	// true -> false
	if lhsObj.Literal.Kind == True {
		return NewFalse(), nil
	}

	// else (false) -> true
	return NewTrue(), nil
}

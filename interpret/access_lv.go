package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAccess(mem *Memory, accessLv *analyze.AccessLevel) (*Object, error) {
	switch accessLv.Kind {
	case analyze.ACLiteralLevel:
		return evalLiteral(mem, accessLv.LiteralLevel)
	case analyze.ACDictIndex:
		index, err := evalExpr(mem, accessLv.Index)
		if err != nil {
			return nil, err
		}
		src, err := evalAccess(mem, accessLv.Src)
		if err != nil {
			return nil, err
		}
		return src.D[index.S], nil
	case analyze.ACListIndex:
		index, err := evalExpr(mem, accessLv.Index)
		if err != nil {
			return nil, err
		}
		src, err := evalAccess(mem, accessLv.Src)
		if err != nil {
			return nil, err
		}
		return src.L[index.I], nil
	case analyze.ACUnknownIndex:
	}
	return nil, fmt.Errorf("unimplemented: %s", accessLv.Kind.String())
}

func softEvalAccess(mem *Memory, accessLv *analyze.AccessLevel) (*Object, *Object, error) {
	switch accessLv.Kind {
	case analyze.ACDictIndex:
		index, err := evalExpr(mem, accessLv.Index)
		if err != nil {
			return nil, nil, err
		}
		src, err := evalAccess(mem, accessLv.Src)
		if err != nil {
			return nil, nil, err
		}
		return src, index, nil
	case analyze.ACListIndex:
		index, err := evalExpr(mem, accessLv.Index)
		if err != nil {
			return nil, nil, err
		}
		src, err := evalAccess(mem, accessLv.Src)
		if err != nil {
			return nil, nil, err
		}
		return src, index, nil
	case analyze.ACUnknownIndex:
	}
	return nil, nil, fmt.Errorf("unimplemented: %s", accessLv.Kind.String())
}

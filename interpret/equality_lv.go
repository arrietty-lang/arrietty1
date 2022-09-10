package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalEquality(mem *Memory, equalityLv *analyze.EqualityLevel) (*Object, error) {
	switch equalityLv.Kind {
	case analyze.EQRelationalLevel:
		return evalRelation(mem, equalityLv.RelationalLevel)
	case analyze.EQEqual:
		lhs, err := evalEquality(mem, equalityLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalEquality(mem, equalityLv.RHS)
		if err != nil {
			return nil, err
		}
		return eq(lhs, rhs)
	case analyze.EQNotEqual:
		lhs, err := evalEquality(mem, equalityLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalEquality(mem, equalityLv.RHS)
		if err != nil {
			return nil, err
		}
		return ne(lhs, rhs)
	}
	return nil, fmt.Errorf("unimplemented: %s", equalityLv.Kind.String())
}

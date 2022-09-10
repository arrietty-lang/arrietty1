package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalRelation(mem *Memory, relationalLv *analyze.RelationalLevel) (*Object, error) {
	switch relationalLv.Kind {
	case analyze.REAddLevel:
		return evalAdd(mem, relationalLv.AddLevel)
	case analyze.RELt:
		lhs, err := evalRelation(mem, relationalLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalRelation(mem, relationalLv.RHS)
		if err != nil {
			return nil, err
		}
		return lt(lhs, rhs)
	case analyze.RELe:
		lhs, err := evalRelation(mem, relationalLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalRelation(mem, relationalLv.RHS)
		if err != nil {
			return nil, err
		}
		return le(lhs, rhs)
	case analyze.REGt:
		lhs, err := evalRelation(mem, relationalLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalRelation(mem, relationalLv.RHS)
		if err != nil {
			return nil, err
		}
		return gt(lhs, rhs)
	case analyze.REGe:
		lhs, err := evalRelation(mem, relationalLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalRelation(mem, relationalLv.RHS)
		if err != nil {
			return nil, err
		}
		return ge(lhs, rhs)
	}
	return nil, fmt.Errorf("unimplemented: %s", relationalLv.Kind.String())
}

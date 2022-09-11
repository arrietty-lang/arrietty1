package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAndOr(mem *Memory, andorLv *analyze.AndOrLevel) (*Object, error) {
	switch andorLv.Kind {
	case analyze.ANEqualityLevel:
		return evalEquality(mem, andorLv.EqualityLevel)
	case analyze.ANAnd:
		lhs, err := evalAndOr(mem, andorLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalAndOr(mem, andorLv.RHS)
		if err != nil {
			return nil, err
		}
		switch {
		case lhs.Kind == OTrue && rhs.Kind == OTrue:
			return NewTrueObject(), nil
		case lhs.Kind == OTrue && rhs.Kind == OFalse:
			return NewFalseObject(), nil
		case lhs.Kind == OFalse && rhs.Kind == OTrue:
			return NewFalseObject(), nil
		case lhs.Kind == OFalse && rhs.Kind == OFalse:
			return NewFalseObject(), nil
		}
	case analyze.ANOr:
		lhs, err := evalAndOr(mem, andorLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalAndOr(mem, andorLv.RHS)
		if err != nil {
			return nil, err
		}
		switch {
		case lhs.Kind == OTrue && rhs.Kind == OTrue:
			return NewTrueObject(), nil
		case lhs.Kind == OTrue && rhs.Kind == OFalse:
			return NewTrueObject(), nil
		case lhs.Kind == OFalse && rhs.Kind == OTrue:
			return NewTrueObject(), nil
		case lhs.Kind == OFalse && rhs.Kind == OFalse:
			return NewFalseObject(), nil
		}
	}
	return nil, fmt.Errorf("unimplemented: %s", andorLv.Kind.String())
}

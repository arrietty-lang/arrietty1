package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalMul(mem *Memory, mulLv *analyze.MulLevel) (*Object, error) {
	switch mulLv.Kind {
	case analyze.MUUnaryLevel:
		return evalUnary(mem, mulLv.UnaryLevel)
	case analyze.MUMul:
		lhs, err := evalMul(mem, mulLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalMul(mem, mulLv.RHS)
		if err != nil {
			return nil, err
		}

		switch {
		case lhs.Kind == OFloat && rhs.Kind == OFloat:
			return NewFloatObject(lhs.F * rhs.F), nil
		case lhs.Kind == OInt && rhs.Kind == OInt:
			return NewIntObject(lhs.I * rhs.I), nil
		case lhs.Kind == OFloat && rhs.Kind == OInt:
			return NewFloatObject(lhs.F * float64(rhs.I)), nil
		case lhs.Kind == OInt && rhs.Kind == OFloat:
			return NewFloatObject(float64(lhs.I) * rhs.F), nil
		}
	case analyze.MUDiv:
		lhs, err := evalMul(mem, mulLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalMul(mem, mulLv.RHS)
		if err != nil {
			return nil, err
		}

		switch {
		case lhs.Kind == OFloat && rhs.Kind == OFloat:
			if rhs.F == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			return NewFloatObject(lhs.F / rhs.F), nil
		case lhs.Kind == OInt && rhs.Kind == OInt:
			if rhs.I == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			return NewIntObject(lhs.I / rhs.I), nil
		case lhs.Kind == OFloat && rhs.Kind == OInt:
			if rhs.I == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			return NewFloatObject(lhs.F / float64(rhs.I)), nil
		case lhs.Kind == OInt && rhs.Kind == OFloat:
			if rhs.F == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			return NewFloatObject(float64(lhs.I) / rhs.F), nil
		}
	case analyze.MUMod:
		lhs, err := evalMul(mem, mulLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalMul(mem, mulLv.RHS)
		if err != nil {
			return nil, err
		}

		switch {
		case lhs.Kind == OInt && rhs.Kind == OInt:
			if rhs.I == 0 {
				return nil, fmt.Errorf("devision by zero")
			}
			return NewIntObject(lhs.I % rhs.I), nil
		}
	}
	return nil, fmt.Errorf("unimplemented: %s", mulLv.Kind.String())
}

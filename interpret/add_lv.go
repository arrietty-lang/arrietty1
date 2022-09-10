package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAdd(mem *Memory, addLv *analyze.AddLevel) (*Object, error) {
	switch addLv.Kind {
	case analyze.ADMulLevel:
		return evalMul(mem, addLv.MulLevel)
	case analyze.ADAdd:
		lhs, err := evalAdd(mem, addLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalAdd(mem, addLv.RHS)
		if err != nil {
			return nil, err
		}
		// どちらもInt
		if lhs.Kind == OInt && rhs.Kind == OInt {
			return NewIntObject(lhs.I + rhs.I), nil
		}
		// どちらもFloat
		if lhs.Kind == OFloat && rhs.Kind == OFloat {
			return NewFloatObject(lhs.F + rhs.F), nil
		}
		// int + float
		if lhs.Kind == OInt && rhs.Kind == OFloat {
			return NewFloatObject(float64(lhs.I) + rhs.F), nil
		}
		// float + int
		if lhs.Kind == OFloat && rhs.Kind == OInt {
			return NewFloatObject(lhs.F + float64(rhs.I)), nil
		}
		if lhs.Kind == OString {
			return NewStringObject(lhs.S + rhs.S), nil
		}
	case analyze.ADSub:
		lhs, err := evalAdd(mem, addLv.LHS)
		if err != nil {
			return nil, err
		}
		rhs, err := evalAdd(mem, addLv.RHS)
		if err != nil {
			return nil, err
		}
		// どちらもInt
		if lhs.Kind == OInt && rhs.Kind == OInt {
			return NewIntObject(lhs.I - rhs.I), nil
		}
		// どちらもFloat
		if lhs.Kind == OFloat && rhs.Kind == OFloat {
			return NewFloatObject(lhs.F - rhs.F), nil
		}
		// int - float
		if lhs.Kind == OInt && rhs.Kind == OFloat {
			return NewFloatObject(float64(lhs.I) - rhs.F), nil
		}
		// float - int
		if lhs.Kind == OFloat && rhs.Kind == OInt {
			return NewFloatObject(lhs.F - float64(rhs.I)), nil
		}
	}
	return nil, fmt.Errorf("unimplemented: %s", addLv.Kind.String())
}

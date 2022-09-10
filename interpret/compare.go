package interpret

import "fmt"

func lt(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F < rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I < rhs.I), nil
	default:
		return nil, fmt.Errorf("%s is not supported by LT", lhs.Kind.String())
	}
}
func le(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F <= rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I <= rhs.I), nil
	default:
		return nil, fmt.Errorf("%s is not supported by LE", lhs.Kind.String())
	}
}
func gt(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F > rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I > rhs.I), nil
	default:
		return nil, fmt.Errorf("%s is not supported by GT", lhs.Kind.String())
	}
}
func ge(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F >= rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I >= rhs.I), nil
	default:
		return nil, fmt.Errorf("%s is not supported by GE", lhs.Kind.String())
	}
}

func eq(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F == rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I == rhs.I), nil
	case OString:
		return NewBoolObject(lhs.S == rhs.S), nil
	case OTrue, OFalse, ONull:
		return NewBoolObject(lhs.Kind == rhs.Kind), nil
	case OList:
		if len(lhs.L) != len(rhs.L) {
			return NewFalseObject(), nil
		}
		for i, item := range lhs.L {
			e, err := eq(item, rhs.L[i])
			if err != nil {
				return nil, err
			}
			if e.Kind == OFalse {
				return NewFalseObject(), nil
			}
		}
		return NewTrueObject(), nil
	}

	return nil, fmt.Errorf("%s unsuppoertd by eq", lhs.Kind.String())
}

func ne(lhs, rhs *Object) (*Object, error) {
	switch lhs.Kind {
	case OFloat:
		return NewBoolObject(lhs.F != rhs.F), nil
	case OInt:
		return NewBoolObject(lhs.I != rhs.I), nil
	case OString:
		return NewBoolObject(lhs.S != rhs.S), nil
	case OTrue, OFalse, ONull:
		return NewBoolObject(lhs.Kind != rhs.Kind), nil
	case OList:
		if len(lhs.L) == len(rhs.L) {
			return NewFalseObject(), nil
		}
		for i, item := range lhs.L {
			e, err := eq(item, rhs.L[i])
			if err != nil {
				return nil, err
			}
			if e.Kind == OFalse {
				return NewTrueObject(), nil
			}
		}
		return NewFalseObject(), nil
	}

	return nil, fmt.Errorf("%s unsuppoertd by ne", lhs.Kind.String())
}

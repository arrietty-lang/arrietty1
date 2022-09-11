package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalUnary(mem *Memory, unaryLv *analyze.UnaryLevel) (*Object, error) {
	switch unaryLv.Kind {
	case analyze.UNPrimaryLevel:
		return evalPrimary(mem, unaryLv.PrimaryLevel)
	case analyze.UNPlus:
		// +なので何もしないで返して良い
		return evalPrimary(mem, unaryLv.PrimaryLevel)
	case analyze.UNMinus:
		// float, intを想定していてそれ以外は受け付けない
		d, err := evalPrimary(mem, unaryLv.PrimaryLevel)
		if err != nil {
			return nil, err
		}
		switch d.Kind {
		case OFloat:
			return NewFloatObject(-d.F), nil
		case OInt:
			return NewIntObject(-d.I), nil
		}
	case analyze.UNNot:
		// true, falseしか受け付けてない
		b, err := evalPrimary(mem, unaryLv.PrimaryLevel)
		if err != nil {
			return nil, err
		}
		switch b.Kind {
		case OTrue:
			return NewFalseObject(), nil
		case OFalse:
			return NewTrueObject(), nil
		}
	}
	return nil, fmt.Errorf("unimplemented: %s", unaryLv.Kind.String())
}

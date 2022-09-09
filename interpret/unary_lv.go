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
	case analyze.UNMinus:
	case analyze.UNNot:

	}
	return nil, fmt.Errorf("unimplemented: %s", unaryLv.Kind.String())
}

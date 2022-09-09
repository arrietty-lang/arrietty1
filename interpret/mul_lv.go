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
	case analyze.MUDiv:
	case analyze.MUMod:

	}
	return nil, fmt.Errorf("unimplemented: %s", mulLv.Kind.String())
}

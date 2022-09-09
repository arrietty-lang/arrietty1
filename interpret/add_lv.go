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
	case analyze.ADSub:

	}
	return nil, fmt.Errorf("unimplemented: %s", addLv.Kind.String())
}

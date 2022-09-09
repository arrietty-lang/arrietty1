package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalExpr(mem *Memory, exprLv *analyze.ExprLevel) (*Object, error) {
	switch exprLv.Kind {
	case analyze.EXAssignLevel:
		return evalAssign(mem, exprLv.AssignLevel)
	}
	return nil, fmt.Errorf("unimplemented: %s", exprLv.Kind.String())
}

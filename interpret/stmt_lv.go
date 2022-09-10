package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalStmt(mem *Memory, stmtLv *analyze.StmtLevel) (*Object, error) {
	switch stmtLv.Kind {
	case analyze.STComment:
		return nil, nil
	case analyze.STExprLevel:
		return evalExpr(mem, stmtLv.ExprLevel)
	case analyze.STReturn:
		return statReturn(mem, stmtLv.Return)
	case analyze.STIfElse:
		return statIfElse(mem, stmtLv.IfElse)
	case analyze.STWhile:
		return statWhile(mem, stmtLv.While)
	case analyze.STFor:
		return statFor(mem, stmtLv.For)
	}
	return nil, fmt.Errorf("unimplemented: %s", stmtLv.Kind.String())
}

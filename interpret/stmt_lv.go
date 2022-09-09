package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func ret(mem *Memory, stmtLv *analyze.StmtLevel) (*Object, error) {
	return evalExpr(mem, stmtLv.Return.Value)
}

func evalStmt(mem *Memory, stmtLv *analyze.StmtLevel) (*Object, error) {
	switch stmtLv.Kind {
	case analyze.STComment:
		return nil, nil
	case analyze.STExprLevel:
		return evalExpr(mem, stmtLv.ExprLevel)
	case analyze.STReturn:
		return ret(mem, stmtLv)

	case analyze.STIfElse:
	case analyze.STWhile:
	case analyze.STFor:

	}
	return nil, fmt.Errorf("unimplemented: %s", stmtLv.Kind.String())
}

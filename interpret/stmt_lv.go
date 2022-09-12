package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalStmt(mem *Memory, stmtLv *analyze.StmtLevel) (*Object, bool, error) {
	switch stmtLv.Kind {
	case analyze.STComment:
		return nil, false, nil
	case analyze.STExprLevel:
		v, err := evalExpr(mem, stmtLv.ExprLevel)
		return v, false, err
	case analyze.STReturn:
		return statReturn(mem, stmtLv.Return)
	case analyze.STIfElse:
		return statIfElse(mem, stmtLv.IfElse)
	case analyze.STWhile:
		return statWhile(mem, stmtLv.While)
	case analyze.STFor:
		return statFor(mem, stmtLv.For)
	case analyze.STBlock:
		return statBlock(mem, stmtLv.Block)
	}
	return nil, false, fmt.Errorf("unimplemented: %s", stmtLv.Kind.String())
}

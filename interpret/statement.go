package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func statReturn(mem *Memory, stmt *analyze.Return) (*Object, error) {
	v, err := evalExpr(mem, stmt.Value)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return NewReturnValue(v), nil
}
func statIfElse(mem *Memory, stmt *analyze.IfElse) (*Object, error) {
	cond, err := evalExpr(mem, stmt.Cond)
	if err != nil {
		return nil, err
	}

	if cond.Kind == OTrue {
		return evalBlock(mem, stmt.IfBlock)
	}

	if cond.Kind == OFalse {
		if stmt.ElseBlock == nil {
			return nil, nil
		}
		return evalBlock(mem, stmt.ElseBlock)
	}

	return nil, fmt.Errorf("invalid condition: %s", cond.Kind.String())
}

func statWhile(mem *Memory, stmt *analyze.While) (*Object, error) {
	for {
		cond, err := evalExpr(mem, stmt.Cond)
		if err != nil {
			return nil, err
		}
		if cond.Kind == OFalse {
			return nil, nil
		}
		ret, err := evalBlock(mem, stmt.WhileBlock)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			return ret, err
		}
	}
}

func statFor(mem *Memory, stmt *analyze.For) (*Object, error) {
	_, err := evalExpr(mem, stmt.Init)
	if err != nil {
		return nil, err
	}
	for {
		cond, err := evalExpr(mem, stmt.Cond)
		if err != nil {
			return nil, err
		}
		if cond.Kind == OFalse {
			return nil, nil
		}
		ret, err := evalBlock(mem, stmt.ForBlock)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			return ret, nil
		}
		_, err = evalExpr(mem, stmt.Loop)
		if err != nil {
			return nil, err
		}
	}
}

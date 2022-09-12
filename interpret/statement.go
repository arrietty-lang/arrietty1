package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func statReturn(mem *Memory, stmt *analyze.Return) (*Object, bool, error) {
	// return;
	if stmt == nil {
		return nil, true, nil
	}

	// 戻り値の解析
	ret, err := evalExpr(mem, stmt.Value)
	if err != nil {
		return nil, false, err
	}

	// return xxx;
	return ret, true, nil
}

func statIfElse(mem *Memory, stmt *analyze.IfElse) (*Object, bool, error) {
	// 条件式の解析
	cond, err := evalExpr(mem, stmt.Cond)
	if err != nil {
		return nil, false, err
	}

	switch cond.Kind {
	case OTrue:
		r, IsReturn, err := evalStmt(mem, stmt.IfBlock)
		if err != nil {
			return nil, false, err
		}
		if IsReturn {
			return r, IsReturn, nil
		}
		return nil, false, nil
	case OFalse:
		if stmt.ElseBlock == nil {
			return nil, false, nil
		}
		r, isReturn, err := evalStmt(mem, stmt.ElseBlock)
		if err != nil {
			return nil, false, err
		}
		if isReturn {
			return r, isReturn, nil
		}
		return nil, false, nil
	}

	return nil, false, fmt.Errorf("invalid condition")
}

func statWhile(mem *Memory, stmt *analyze.While) (*Object, bool, error) {
	for {
		// 条件式の解析
		cond, err := evalExpr(mem, stmt.Cond)
		if err != nil {
			return nil, false, err
		}
		// 条件に沿わなかったのでループから抜ける
		if cond.Kind == OFalse {
			return nil, false, nil
		}

		// ブロックの解析
		r, isReturn, err := evalStmt(mem, stmt.WhileBlock)
		if err != nil {
			return nil, false, err
		}
		// もし戻り値とマークされたものが返ってきたら返す
		if isReturn {
			return r, isReturn, nil
		}
	}
}

func statFor(mem *Memory, stmt *analyze.For) (*Object, bool, error) {
	// 初期化式を解析
	_, err := evalExpr(mem, stmt.Init)
	if err != nil {
		return nil, false, err
	}

	for {
		// 条件式を解析
		cond, err := evalExpr(mem, stmt.Cond)
		if err != nil {
			return nil, false, err
		}

		// 条件に沿わなかったのでループを抜ける
		if cond.Kind == OFalse {
			return nil, false, err
		}

		// ブロックの中身を解析
		r, isReturn, err := evalStmt(mem, stmt.ForBlock)
		if err != nil {
			return nil, false, err
		}
		// もし戻り値としてマークされていたら返す
		if isReturn {
			return r, isReturn, nil
		}

		// ループ式を解析
		_, err = evalExpr(mem, stmt.Loop)
		if err != nil {
			return nil, false, err
		}
	}
}

func statBlock(mem *Memory, stmt *analyze.Block) (*Object, bool, error) {
	for _, statement := range stmt.Body {
		r, isReturn, err := evalStmt(mem, statement)
		if err != nil {
			return nil, false, err
		}
		if isReturn {
			return r, isReturn, nil
		}
	}
	return nil, false, nil
}

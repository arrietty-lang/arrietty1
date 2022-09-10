package interpret

import "github.com/x0y14/arrietty/analyze"

func evalBlock(mem *Memory, stmts []*analyze.StmtLevel) (*Object, error) {
	for _, stmt := range stmts {
		// 1行ごとに結果が返ってくる
		result, err := evalStmt(mem, stmt)
		if err != nil {
			return nil, err
		}
		if result != nil && result.IsReturnValue {
			return result, nil
		}
	}
	return nil, nil
}

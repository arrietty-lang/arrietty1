package interpret

import (
	"github.com/x0y14/arrietty/analyze"
)

func args2Objs(mem *Memory, args []*analyze.ExprLevel) ([]*Object, error) {
	var objs []*Object
	for _, arg := range args {
		obj, err := evalExpr(mem, arg)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func ExecFunction(mem *Memory, f *analyze.FuncDef, args []*analyze.ExprLevel) (*Object, error) {
	// 関数で発生したデータを保存するための領域を作成
	localMem := NewMemory(nil, nil)

	// 関数の戻り値
	var ret *Object = nil

	// パラメーターを宣言
	for _, param := range f.Params {
		err := localMem.DeclareVar(param.Ident)
		if err != nil {
			return nil, err
		}
	}

	// args(expr)をobjに変換してから引数を割り当てる
	objectArgs, err := args2Objs(mem, args)
	if err != nil {
		return nil, err
	}
	for i, arg := range objectArgs {
		paramName := f.Params[i].Ident
		err := localMem.AssignVar(paramName, arg, false) // 型も識別子も解決された状態でくるから気にしなくていいと思うけど念の為上書きを規制する
		if err != nil {
			return nil, err
		}
	}

	// bodyを実行
	for _, stmt := range f.Body {
		// 1行ごとに結果が返ってくる
		result, err := evalStmt(localMem, stmt)
		if err != nil {
			return nil, err
		}
		// もしReturnが実施されたなら結果を返すために保存
		if stmt.Kind == analyze.STReturn {
			ret = result
		}
	}

	return ret, nil
}

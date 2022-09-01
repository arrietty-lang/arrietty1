package interpret

import "github.com/x0y14/arrietty/parse"

func while_(scope *Storage, node *parse.Node) (*Object, error) {
	whileBlock := node.Children[0]

	for {
		// 条件確認
		condObj, err := eval(scope, node.Cond)
		if err != nil {
			return nil, err
		}
		if condObj.Literal.Kind != True {
			break
		}

		// ループの中身実行
		r, err := eval(scope, whileBlock)
		if err != nil {
			return nil, err
		}
		if r != nil && r.IsResult {
			return r, nil
		}
	}

	return nil, nil
}

func for_(scope *Storage, node *parse.Node) (*Object, error) {
	// 初期化式を実行
	_, err := eval(scope, node.Init)
	if err != nil {
		return nil, err
	}

	forBlock := node.Children[0]

	for {
		// 条件を確認
		condObj, err := eval(scope, node.Cond)
		if err != nil {
			return nil, err
		}
		if condObj.Literal.Kind != True {
			break
		}

		// ループの中身実行
		r, err := eval(scope, forBlock)
		if err != nil {
			return nil, err
		}
		if r != nil && r.IsResult {
			return r, nil
		}

		// ループ毎に実行される式を実行
		_, err = eval(scope, node.Loop)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

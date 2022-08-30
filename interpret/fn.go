package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type Fn struct {
	Params []*Object
	Body   *parse.Node
	Local  *Storage
}

func (f *Fn) Exec(args []*Object) (*Object, error) {
	// パラメータと受け取った引数の個数が一致することを確認
	if len(f.Params) != len(args) {
		return nil, fmt.Errorf("the number of parameters(expect %d) & arguments(actual %d) do not match", len(f.Params), len(args))
	}

	// 関数スコープストレージにデータを保存
	for i, param := range f.Params {
		err := Store(f.Local, param.Str, args[i])
		if err != nil {
			return nil, err
		}
	}

	// evalは式の結果を返してくれるが、return式以外の式は結果として返却をしないので個別のノードで計算を行う
	// -> 一括で実行しない
	var result *Object = nil
	for _, n := range f.Body.Children {
		r, err := eval(f.Local, n)
		if err != nil {
			return nil, err
		}
		if (*n).Kind == parse.Return {
			result = r
		}
	}
	return result, nil
}

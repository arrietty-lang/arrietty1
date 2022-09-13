package analyze

import "fmt"

var symbols map[string]map[string]*DataType
var currentFunction string

func init() {
	symbols = map[string]map[string]*DataType{}
	setBuiltIn() // 最初につける
}

func defineVar(funcScope string, id string, dataType *DataType) error {
	_, ok := isDefinedFunc(id)
	if ok {
		return NewAlreadyDefinedErr("file-toplevel", id)
	}

	_, ok = isDefinedVariable(funcScope, id)
	if ok {
		return NewAlreadyDefinedErr(funcScope, id)
	}

	symbols[funcScope][id] = dataType
	return nil
}

func isDefinedFunc(name string) (*DataType, bool) {
	v, ok := symbols[name]
	if ok {
		// 戻り値を型として返してあげる
		return v[""], true
	}
	return nil, false
}

func isDefinedVariable(funcScope string, id string) (*DataType, bool) {
	// スコープ内に定義されているかを確認
	t, ok := symbols[funcScope][id]
	if ok {
		return t, true
	}

	return nil, false
}

func isDefinableIdent(funcScope string, id string) bool {
	_, ok := isDefinedFunc(id)
	if ok {
		return false
	}

	_, ok = symbols[funcScope][id]
	if ok {
		return false
	}
	return true
}

func getFuncParams(name string) (map[string]*FuncParam, error) {
	_, ok := isDefinedFunc(name)
	if !ok {
		return nil, fmt.Errorf("can't get undefined function's paramaters")
	}

	params := map[string]*FuncParam{}

	for paramName := range symbols[name] {
		if paramName == "" {
			// 関数の戻り値
			continue
		}
		params[paramName] = &FuncParam{
			Ident: paramName,
			Type:  symbols[name][paramName],
		}
	}

	return params, nil
}

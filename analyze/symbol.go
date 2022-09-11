package analyze

var symbols map[string]map[string]*DataType
var currentFunction string

func init() {
	symbols = map[string]map[string]*DataType{}
	setBuiltIn() // 最初につける
}

func setBuiltIn() {
	symbols["strlen"] = map[string]*DataType{
		"":  {Type: TInt},    // 戻り値
		"v": {Type: TString}, // 引数1
	}
	symbols["len"] = map[string]*DataType{
		"":  {Type: TInt},
		"v": {Type: TList, Item: &DataType{Type: TAny}},
	}
	symbols["append"] = map[string]*DataType{
		"":   {Type: TVoid},
		"to": {Type: TList, Item: &DataType{Type: TAny}},
		"v":  {Type: TAny},
	}
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

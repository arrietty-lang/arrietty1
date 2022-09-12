package analyze

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
	symbols["print"] = map[string]*DataType{
		"":  {Type: TVoid},
		"v": {Type: TString},
	}
	symbols["itos"] = map[string]*DataType{
		"":  {Type: TString},
		"v": {Type: TInt},
	}
	symbols["split"] = map[string]*DataType{
		"":       {Type: TList, Item: &DataType{Type: TString}},
		"target": {Type: TString},
		"sep":    {Type: TString},
	}
	symbols["keys"] = map[string]*DataType{
		"": {Type: TList, Item: &DataType{Type: TString}},
		"d": {Type: TDict,
			Key:   &DataType{Type: TString},
			Value: &DataType{Type: TAny}},
	}
	symbols["stoi"] = map[string]*DataType{
		"":  {Type: TInt},
		"v": {Type: TString},
	}
}

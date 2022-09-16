package analyze

func attachBuiltin() {
	// 領域確保
	b, err := symbolTable.DeclarePkg("builtin")
	if err != nil {
		panic(err)
	}

	builtinPkg = b
	builtinFunc := []struct {
		name string
		f    *FunctionSymbol
	}{
		{
			"strlen",
			&FunctionSymbol{
				Public:     false,
				Ident:      "strlen",
				ReturnType: &DataType{Type: TInt},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TString}},
				},
				LocalVariables: nil,
			},
		},
		{
			"len",
			&FunctionSymbol{
				Public:     false,
				Ident:      "len",
				ReturnType: &DataType{Type: TInt},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TList, Item: &DataType{Type: TAny}, Size: -1}},
				},
				LocalVariables: nil,
			},
		},
		{
			"append",
			&FunctionSymbol{
				Public:     false,
				Ident:      "append",
				ReturnType: &DataType{Type: TVoid},
				Params: []*VariableSymbol{
					{Ident: "to", DataType: &DataType{Type: TList, Item: &DataType{Type: TAny}}},
					{Ident: "v", DataType: &DataType{Type: TAny}},
				},
				LocalVariables: nil,
			},
		},
		{
			"print",
			&FunctionSymbol{
				Public:     false,
				Ident:      "print",
				ReturnType: &DataType{Type: TVoid},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TString}},
				},
				LocalVariables: nil,
			},
		},
		{
			"itos",
			&FunctionSymbol{
				Public:     false,
				Ident:      "itos",
				ReturnType: &DataType{Type: TString},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TInt}},
				},
				LocalVariables: nil,
			},
		},
		{
			"split",
			&FunctionSymbol{
				Public:     false,
				Ident:      "split",
				ReturnType: &DataType{Type: TList, Item: &DataType{Type: TString}},
				Params: []*VariableSymbol{
					{Ident: "target", DataType: &DataType{Type: TString}},
					{Ident: "sep", DataType: &DataType{Type: TString}},
				},
				LocalVariables: nil,
			},
		},
		{
			"keys",
			&FunctionSymbol{
				Public:     false,
				Ident:      "keys",
				ReturnType: &DataType{Type: TList, Item: &DataType{Type: TString}},
				Params: []*VariableSymbol{
					{Ident: "d", DataType: &DataType{Type: TDict, Key: &DataType{Type: TString}, Value: &DataType{Type: TAny}}},
				},
				LocalVariables: nil,
			},
		},
		{
			"stoi",
			&FunctionSymbol{
				Public:     false,
				Ident:      "stoi",
				ReturnType: &DataType{Type: TInt},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TString}},
				},
				LocalVariables: nil,
			},
		},
		{
			"as_string",
			&FunctionSymbol{
				Public:     false,
				Ident:      "as_string",
				ReturnType: &DataType{Type: TString},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TAny}},
				},
				LocalVariables: nil,
			},
		},
		{
			"as_float",
			&FunctionSymbol{
				Public:     false,
				Ident:      "as_float",
				ReturnType: &DataType{Type: TFloat},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TAny}},
				},
				LocalVariables: nil,
			},
		},
		{
			"as_int",
			&FunctionSymbol{
				Public:     false,
				Ident:      "as_int",
				ReturnType: &DataType{Type: TInt},
				Params: []*VariableSymbol{
					{Ident: "v", DataType: &DataType{Type: TAny}},
				},
				LocalVariables: nil,
			},
		},
	}

	for _, builtin := range builtinFunc {
		//symbolTable.packages["builtin"].Functions[builtin.name] = builtin.f
		funcDecl, err := builtinPkg.DeclareFunc(builtin.name)
		if err != nil {
			panic(err)
		}
		funcDecl.Public = builtin.f.Public
		funcDecl.Ident = builtin.f.Ident
		funcDecl.ReturnType = builtin.f.ReturnType
		funcDecl.Params = builtin.f.Params
		funcDecl.LocalVariables = builtin.f.LocalVariables
	}
}

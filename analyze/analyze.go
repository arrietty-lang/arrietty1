package analyze

import "github.com/x0y14/arrietty/parse"

var Loaded map[string]*TopLevel

func CleanUp() {
	currentFunction = ""
	symbols = map[string]map[string]*DataType{}
	Loaded = map[string]*TopLevel{}
}

func Init() {
	Loaded = map[string]*TopLevel{}
	LoadBuiltIn()
}

func LoadBuiltIn() {
	Loaded["str_len"] = &TopLevel{
		Kind: TPFuncDef,
		FuncDef: &FuncDef{
			ReturnType: &DataType{Type: TInt},
			Ident:      "str_len",
			Params: []*FuncParam{
				{Ident: "v", Type: &DataType{Type: TString}},
			},
			Body: nil,
		},
		Comment: nil,
	}
}

func Analyze(nodes []*parse.Node) (map[string]*TopLevel, error) {
	Init()
	for _, n := range nodes {
		top, err := NewToplevel(n)
		if err != nil {
			return nil, err
		}
		if top.Kind == TPFuncDef {
			Loaded[top.FuncDef.Ident] = top
		}
	}
	return Loaded, nil
}

package analyze

import "github.com/x0y14/arrietty/parse"

func CleanUp() {
	currentFunction = ""
	symbols = map[string]map[string]*DataType{}
	setBuiltIn() // 掃除後もつける
}

func Analyze(nodes []*parse.Node) (map[string]*TopLevel, error) {
	scripts := map[string]*TopLevel{}
	for _, n := range nodes {
		top, err := NewToplevel(n)
		if err != nil {
			return nil, err
		}
		if top.Kind == TPFuncDef {
			scripts[top.FuncDef.Ident] = top
		}
	}
	return scripts, nil
}

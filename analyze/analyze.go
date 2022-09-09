package analyze

import "github.com/x0y14/arrietty/parse"

func Analyze(nodes []*parse.Node) (map[string]*TopLevel, error) {
	script := map[string]*TopLevel{}

	for _, n := range nodes {
		top, err := NewToplevel(n)
		if err != nil {
			return nil, err
		}
		if top.Kind == TPFuncDef {
			script[top.FuncDef.Ident] = top
		}
	}
	return script, nil
}

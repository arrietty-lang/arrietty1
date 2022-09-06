package analyze

import "github.com/x0y14/arrietty/parse"

type FuncDef struct {
	ReturnType *ValueType
	Ident      string
	Params     []*FuncParam
	Body       []*StmtLevel
}

func NewFuncDef(retType, name, params, body *parse.Node) (*FuncDef, error) {

	returnType, err := NewValueType(retType)
	if err != nil {
		return nil, err
	}

	functionName := name.S

	var ps []*FuncParam
	for _, paramNode := range params.Children {
		p, err := NewFuncParam(paramNode)
		if err != nil {
			return nil, err
		}

		if !isAvailableVarIdent(functionName, p.Ident) {
			return nil, NewAlreadyDefinedErr(functionName, p.Ident)
		}

		ps = append(ps, p)
	}

	// todo: body

	return &FuncDef{
		ReturnType: returnType,
		Ident:      functionName,
		Params:     ps,
		Body:       nil,
	}, nil
}

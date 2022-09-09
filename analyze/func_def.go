package analyze

import "github.com/x0y14/arrietty/parse"

type FuncDef struct {
	ReturnType *DataType
	Ident      string
	Params     []*FuncParam
	Body       []*StmtLevel
}

func NewFuncDef(retType, name, params, body *parse.Node) (*FuncDef, error) {
	functionName := name.S

	// body解析を行う際に必要になるので設定しておく
	currentFunction = functionName

	// 戻り値解析
	returnType, err := NewDataTypeFromNode(retType)
	if err != nil {
		return nil, err
	}

	// 変数の宣言データを保存する領域を確保
	symbols[functionName] = map[string]*DataType{}
	// Identとして使用できない空白キーに関数の戻り値を設定
	symbols[functionName][""] = returnType

	var ps []*FuncParam
	if params != nil {
		for _, paramNode := range params.Children {
			p, err := NewFuncParam(paramNode)
			if err != nil {
				return nil, err
			}

			// すでに変数として定義されている
			if !isDefinableIdent(functionName, p.Ident) {
				return nil, NewAlreadyDefinedErr(functionName, p.Ident)
			}

			// ローカル変数として宣言してあげる
			err = defineVar(functionName, p.Ident, p.Type)
			if err != nil {
				return nil, err
			}

			ps = append(ps, p)
		}
	}

	var stmts []*StmtLevel
	for _, stmtNode := range body.Children {
		s, err := NewStmtLevel(stmtNode)
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, s)
	}

	return &FuncDef{
		ReturnType: returnType,
		Ident:      functionName,
		Params:     ps,
		Body:       stmts,
	}, nil
}

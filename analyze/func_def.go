package analyze

import (
	"fmt"
	"github.com/x0y14/arrietty/parse"
)

type FuncDef struct {
	ReturnType *DataType
	Ident      string
	Params     []*FuncParam
	Body       *StmtLevel
}

func NewFuncDef(retType, name, params, body *parse.Node) (*FuncDef, error) {
	functionName := name.S

	// funcを宣言
	funcDecl, err := currentPkg.DeclareFunc(functionName)
	if err != nil {
		return nil, err
	}
	currentFunc = funcDecl
	funcDecl.Ident = functionName

	// 戻り値解析
	returnType, err := NewDataTypeFromNode(retType)
	if err != nil {
		return nil, err
	}
	funcDecl.ReturnType = returnType

	// パラメータの解析
	var funcParams []*FuncParam
	if params != nil {
		for _, paramNode := range params.Children {
			param, err := NewFuncParam(paramNode)
			if err != nil {
				return nil, err
			}

			// パラメータを設定
			funcDecl.Params = append(funcDecl.Params, &VariableSymbol{
				Public:   false,
				Ident:    param.Ident,
				DataType: param.Type,
			})

			// パラメータをローカル変数として宣言
			funcParamDecl, err := funcDecl.DeclareLocalVar(param.Ident)
			if err != nil {
				return nil, fmt.Errorf("failed to declare param of %s(): %v", currentFunc, err)
			}
			funcParamDecl.Ident = param.Ident
			funcParamDecl.DataType = param.Type

			funcParams = append(funcParams, param)
		}
	}

	block, err := NewBlock(body)
	if err != nil {
		return nil, err
	}

	return &FuncDef{
		ReturnType: returnType,
		Ident:      functionName,
		Params:     funcParams,
		Body:       &StmtLevel{Kind: STBlock, Block: block},
	}, nil
}

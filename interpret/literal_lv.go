package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalLiteral(mem *Memory, literalLv *analyze.LiteralLevel) (*Object, error) {
	switch literalLv.Kind {
	case analyze.LParentheses:
		return evalExpr(mem, literalLv.ExprLevel)
	case analyze.LIdent:
		return mem.GetVar(literalLv.Ident)
	case analyze.LCall:
		if IsBuiltInFunc(literalLv.Ident) {
			return ExecBuiltIn(literalLv.Ident, mem, literalLv.CallArgs)
		}
		f, err := FileMem.GetFunc(literalLv.Ident)
		if err != nil {
			return nil, err
		}
		return ExecFunction(mem, f, literalLv.CallArgs)
	case analyze.LAtom:
		return ConvertAtomToObject(literalLv.Atom)
	case analyze.LList:
		return ConvertListToObject(mem, literalLv.ListLevel)
	case analyze.LDict:
		return ConvertDictToObject(mem, literalLv.DictLevel)
	}
	return nil, fmt.Errorf("unimplemented: %s", literalLv.Kind.String())
}

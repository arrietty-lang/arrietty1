package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
	"strings"
)

func evalLiteral(mem *Memory, literalLv *analyze.LiteralLevel) (*Object, error) {
	switch literalLv.Kind {
	case analyze.LParentheses:
		return evalExpr(mem, literalLv.ExprLevel)
	case analyze.LIdent:
		return mem.GetVar(literalLv.Ident)
	case analyze.LCall:
		// other pkg
		if strings.Contains(literalLv.Ident, ".") {
			pkgFunc := strings.Split(literalLv.Ident, ".")
			pkgName := pkgFunc[0]
			fnName := pkgFunc[1]
			f, err := runtimeMem.Packages[pkgName].GetFunc(fnName)
			if err != nil {
				return nil, err
			}
			v, _, err := ExecFunction(mem, f, literalLv.CallArgs)
			return v, err
		}
		// builtin
		if IsBuiltInFunc(literalLv.Ident) {
			return ExecBuiltIn(literalLv.Ident, mem, literalLv.CallArgs)
		}
		// current pkg
		f, err := runtimeMem.Packages[currentPkg].GetFunc(literalLv.Ident)
		if err != nil {
			return nil, err
		}
		v, _, err := ExecFunction(mem, f, literalLv.CallArgs)
		return v, err
	case analyze.LAtom:
		return ConvertAtomToObject(literalLv.Atom)
	case analyze.LList:
		return ConvertListToObject(mem, literalLv.ListLevel)
	case analyze.LDict:
		return ConvertDictToObject(mem, literalLv.DictLevel)
	}
	return nil, fmt.Errorf("unimplemented: %s", literalLv.Kind.String())
}

package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func assign(mem *Memory, assignment *analyze.Assignment) error {
	switch assignment.Kind {
	case analyze.ToDefinedIdent:
		value, err := evalAndOr(mem, assignment.Value)
		if err != nil {
			return err
		}
		// 同時に宣言をする必要があるか
		// インラインなら個別に宣言がされていないので、宣言をしてあげる必要がある
		if assignment.Inline {
			err = mem.DeclareVar(assignment.Ident)
			if err != nil {
				return err
			}
		}
		err = mem.AssignVar(assignment.Ident, value, true)
		if err != nil {
			return err
		}
		return nil
	case analyze.ToDictKey, analyze.ToListIndex:
		value, err := evalAndOr(mem, assignment.Value)
		if err != nil {
			return err
		}
		src, index, err := softEvalAccess(mem, assignment.AccessLv)
		if err != nil {
			return err
		}
		err = src.AssignWithIndex(index, value)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unimplemented: %s", assignment.Kind.String())
}

func evalAssign(mem *Memory, assignLv *analyze.AssignLevel) (*Object, error) {
	switch assignLv.Kind {
	case analyze.AAndOrLevel:
		return evalAndOr(mem, assignLv.AndOrLevel)
	case analyze.AVarDecl:
		// 宣言するだけ
		err := mem.DeclareVar(assignLv.VarDecl.Ident)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case analyze.AAssign:
		// 代入するだけ
		err := assign(mem, assignLv.Assignment)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, fmt.Errorf("unimplemented: %s", assignLv.Kind.String())
}

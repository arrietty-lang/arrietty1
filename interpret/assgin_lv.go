package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAssign(mem *Memory, assignLv *analyze.AssignLevel) (*Object, error) {
	switch assignLv.Kind {
	case analyze.AAndOrLevel:
		return evalAndOr(mem, assignLv.AndOrLevel)
	case analyze.AVarDecl:
	case analyze.AAssign:
	}
	return nil, fmt.Errorf("unimplemented: %s", assignLv.Kind.String())
}

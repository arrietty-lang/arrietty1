package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalEquality(mem *Memory, equalityLv *analyze.EqualityLevel) (*Object, error) {
	switch equalityLv.Kind {
	case analyze.EQRelationalLevel:
		return evalRelation(mem, equalityLv.RelationalLevel)
	case analyze.EQEqual:
	case analyze.EQNotEqual:

	}
	return nil, fmt.Errorf("unimplemented: %s", equalityLv.Kind.String())
}

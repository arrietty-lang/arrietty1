package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAndOr(mem *Memory, andorLv *analyze.AndOrLevel) (*Object, error) {
	switch andorLv.Kind {
	case analyze.ANEqualityLevel:
		return evalEquality(mem, andorLv.EqualityLevel)
	case analyze.ANAnd:
	case analyze.ANOr:

	}
	return nil, fmt.Errorf("unimplemented: %s", andorLv.Kind.String())
}

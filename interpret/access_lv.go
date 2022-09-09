package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalAccess(mem *Memory, accessLv *analyze.AccessLevel) (*Object, error) {
	switch accessLv.Kind {
	case analyze.ACLiteralLevel:
		return evalLiteral(mem, accessLv.LiteralLevel)
	case analyze.ACDictIndex:
	case analyze.ACListIndex:
	case analyze.ACUnknownIndex:

	}
	return nil, fmt.Errorf("unimplemented: %s", accessLv.Kind.String())
}

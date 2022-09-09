package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalRelation(mem *Memory, relationalLv *analyze.RelationalLevel) (*Object, error) {
	switch relationalLv.Kind {
	case analyze.REAddLevel:
		return evalAdd(mem, relationalLv.AddLevel)
	case analyze.RELt:
	case analyze.RELe:
	case analyze.REGt:
	case analyze.REGe:

	}
	return nil, fmt.Errorf("unimplemented: %s", relationalLv.Kind.String())
}

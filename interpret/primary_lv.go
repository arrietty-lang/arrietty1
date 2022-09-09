package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

func evalPrimary(mem *Memory, primaryLv *analyze.PrimaryLevel) (*Object, error) {
	switch primaryLv.Kind {
	case analyze.PRAccessLevel:
		return evalAccess(mem, primaryLv.AccessLevel)
	}
	return nil, fmt.Errorf("unimplemented: %s", primaryLv.Kind.String())
}

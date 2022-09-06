package analyze

import "github.com/x0y14/arrietty/parse"

var decls map[string]map[string]*ValueType

func isFuncDefined(name string) bool {
	_, ok := decls[name]
	if ok {
		return true
	}
	return false
}

func isAvailableVarIdent(funcScope string, id string) bool {
	if isFuncDefined(id) {
		return false
	}
	_, ok := decls[funcScope][id]
	if ok {
		return false
	}
	return true
}

func init() {
	decls = map[string]map[string]*ValueType{}
}

func stmtLevel(node *parse.Node) (*StmtLevel, error) {
	switch node.Kind {
	case parse.Comment:
		return NewStmtLevelComment(node.S)
	}
	return nil, nil
}

func toplevel(node *parse.Node) (*TopLevel, error) {
	switch node.Kind {
	case parse.FuncDef:
		return NewTopLevelFuncDef(node)
	case parse.Comment:
		return NewTopLevelComment(node)
	}

	return nil, NewUnexpectNodeErr(node)
}

func Analyze(nodes []*parse.Node) ([]*TopLevel, error) {
	var trees []*TopLevel
	for _, n := range nodes {
		top, err := toplevel(n)
		if err != nil {
			return nil, err
		}
		trees = append(trees, top)
	}
	return trees, nil
}

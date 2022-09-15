package interpret

import (
	"fmt"
	"github.com/x0y14/arrietty/analyze"
)

var currentPkg string
var runtimeMem *RuntimeMemory

func Setup() {
	runtimeMem = &RuntimeMemory{Packages: map[string]*Memory{}}
}

func LoadScript(pkgName string, script map[string]*analyze.TopLevel) error {
	if _, ok := runtimeMem.Packages[pkgName]; ok {
		return fmt.Errorf("pkg: %s is already loaded", pkgName)
	}
	runtimeMem.Packages[pkgName] = NewMemory(nil, script)
	return nil
}

func Interpret(pkgName string, fnName string) (*Object, error) {
	currentPkg = pkgName
	if _, ok := runtimeMem.Packages[pkgName]; !ok {
		return nil, fmt.Errorf("pkg: %s is not loaded", pkgName)
	}

	mainFunc, err := runtimeMem.Packages[currentPkg].GetFunc(fnName)
	if err != nil {
		return nil, err
	}
	v, _, err := ExecFunction(runtimeMem.Packages[currentPkg], mainFunc, nil)
	return v, err
}

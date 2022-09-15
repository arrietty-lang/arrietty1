package apm

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetArriettyPackagesPath() (string, error) {
	aPath := os.Getenv("ARRIETTY_PATH")
	if aPath == "" {
		return "", fmt.Errorf("not found env: ARRIETTY_PATH")
	}
	return filepath.Join(aPath, "packages"), nil
}

func GetInstalledPackageList() ([]string, error) {
	pkgPath, err := GetArriettyPackagesPath()
	if err != nil {
		return nil, err
	}

	var packages []string
	err = filepath.Walk(pkgPath, func(path string, info fs.FileInfo, err error) error {
		packages = append(packages, path)
		return nil
	})

	return packages, err
}

func GetPackageInfo(pkgName string) (*PkgInfo, error) {
	aPath, err := GetArriettyPackagesPath()
	if err != nil {
		return nil, err
	}

	pkgPath := filepath.Join(aPath, pkgName)
	pkgJsonPath := filepath.Join(pkgPath, "pkg.json")

	bytes, err := os.ReadFile(pkgJsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read pkg.json of %s: %v", pkgName, err)
	}
	return UnmarshalPkgJson(bytes)
}

// IsPkgInstalled return ( version, isInstalled )
func IsPkgInstalled(pkgName string) bool {
	packages, err := GetInstalledPackageList()
	if err != nil {
		return false
	}

	for _, pkg := range packages {
		if pkg == pkgName {
			return true
		}
	}
	return false
}

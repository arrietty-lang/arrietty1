package apm

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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
		if pkgPath == path {
			return nil
		}
		s := strings.Split(path, "/")
		pkgName := s[len(s)-1]
		packages = append(packages, pkgName)
		return nil
	})

	return packages, err
}

func GetPackageInfo(pkgName string) (*PkgInfo, error) {
	pkgPath, err := GetPackagePath(pkgName)
	if err != nil {
		return nil, err
	}
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

func GetPackagePath(pkgName string) (string, error) {
	aPath, err := GetArriettyPackagesPath()
	if err != nil {
		return "", err
	}

	pkgPath := filepath.Join(aPath, pkgName)
	return pkgPath, nil
}

func GetArrFilePathsInPackage(pkgName string) ([]string, error) {
	var filePaths []string
	if !IsPkgInstalled(pkgName) {
		return nil, fmt.Errorf("%s is not installed", pkgName)
	}

	pkgPath, err := GetPackagePath(pkgName)
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(pkgPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".arr") {
			return nil
		}
		filePaths = append(filePaths, path)
		return nil
	})

	return filePaths, err
}

func GetCurrentPackageInfo(root string) (*PkgInfo, error) {
	pkgJsonPath := filepath.Join(root, "pkg.json")

	bytes, err := os.ReadFile(pkgJsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read pkg.json: %v", err)
	}
	return UnmarshalPkgJson(bytes)
}

func GetArrFilePathsInCurrent(root string) ([]string, error) {
	var arrFiles []string
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".arr") {
			return nil
		}
		arrFiles = append(arrFiles, path)
		return nil
	})

	return arrFiles, err
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// InitPackage `go mod init ...`のようなパッケージ初期化コマンドに使用
func InitPackage(root string, packageName string) error {
	// すでにpkg.jsonが存在していないかを確認
	pkgJsonPath := filepath.Join(root, "pkg.json")
	if FileExists(pkgJsonPath) {
		return fmt.Errorf("pkg.json already exists: %v", pkgJsonPath)
	}

	pkg := PkgInfo{
		Name:    packageName,
		Version: "0.0.1",
		Deps:    make([]Dependencies, 0),
	}

	bytes, err := json.MarshalIndent(pkg, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(pkgJsonPath, bytes, 0644)
	return err
}

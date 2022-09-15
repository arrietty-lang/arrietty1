package apm

import (
	"encoding/json"
	"fmt"
)

type PkgInfo struct {
	Name    string         `json:"name,omitempty"`
	Version string         `json:"version,omitempty"`
	Deps    []Dependencies `json:"deps,omitempty"`
}

type Dependencies struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Url     string `json:"url,omitempty"`
}

func UnmarshalPkgJson(b []byte) (*PkgInfo, error) {
	var pkgJson PkgInfo
	err := json.Unmarshal(b, &pkgJson)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pkg.json: %v", err)
	}
	return &pkgJson, nil
}

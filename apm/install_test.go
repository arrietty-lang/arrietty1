package apm

import "testing"

func TestInstallPkgFromGithub(t *testing.T) {
	target := "https://github.com/x0y14/arrietty_json"
	err := InstallPkgFromGithub(target)
	if err != nil {
		t.Fatalf("failed to install pkg: %v", err)
	}
}

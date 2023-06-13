package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		baseFilePath   string
		newFilePath    string
		policyFilePath string
		exitCode       int
	}{
		{
			name:           "Denied Diff",
			baseFilePath:   "../example/base.yaml",
			newFilePath:    "../example/new.yaml",
			policyFilePath: "../config.yaml",
			exitCode:       ExitCodeDenied,
		},
		{
			name:           "Equal Manifests",
			baseFilePath:   "../test/allow-equal-manifests/base.yaml",
			newFilePath:    "../test/allow-equal-manifests/new.yaml",
			policyFilePath: "../test/allow-equal-manifests/config.yaml",
			exitCode:       ExitCodeAllowed,
		},
		{
			name:           "Allowed Diff(different image tag)",
			baseFilePath:   "../test/allow-image-tag-diff/base.yaml",
			newFilePath:    "../test/allow-image-tag-diff/new.yaml",
			policyFilePath: "../test/allow-image-tag-diff/config.yaml",
			exitCode:       ExitCodeAllowed,
		},
		{
			name:           "Allowed Diff(different image tag)",
			baseFilePath:   "../test/allow-metadata-recursive-diff/base.yaml",
			newFilePath:    "../test/allow-metadata-recursive-diff/new.yaml",
			policyFilePath: "../test/allow-metadata-recursive-diff/config.yaml",
			exitCode:       ExitCodeAllowed,
		},
		{
			name:           "Denied Diff(different number of docs)",
			baseFilePath:   "../test/deny-different-number-of-docs/base.yaml",
			newFilePath:    "../test/deny-different-number-of-docs/new.yaml",
			policyFilePath: "../test/deny-different-number-of-docs/config.yaml",
			exitCode:       ExitCodeDenied,
		},
		{
			name:           "Denied Diff(different port)",
			baseFilePath:   "../test/deny-different-port/base.yaml",
			newFilePath:    "../test/deny-different-port/new.yaml",
			policyFilePath: "../test/deny-different-port/config.yaml",
			exitCode:       ExitCodeDenied,
		},
		{
			name:           "Denied Diff(different port)",
			baseFilePath:   "../test/deny-different-port/base-noexist.yaml",
			newFilePath:    "../test/deny-different-port/new.yaml",
			policyFilePath: "../test/deny-different-port/config.yaml",
			exitCode:       ExitCodeSomethingError,
		},
		{
			name:           "Denied Diff(different port)",
			baseFilePath:   "../test/deny-different-port/base.yaml",
			newFilePath:    "../test/deny-different-port/new.yaml",
			policyFilePath: "../test/deny-different-port/config-noexist.yaml",
			exitCode:       ExitCodeSomethingError,
		},
	}

	for _, test := range tests {
		code := Run([]string{test.baseFilePath, test.newFilePath}, test.policyFilePath)
		if code != test.exitCode {
			t.Errorf("expected: %v, got: %v", test.exitCode, code)
		}
	}
}

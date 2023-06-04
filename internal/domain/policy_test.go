package domain_test

import (
	"eval-yaml-diff/internal/domain"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name     string
		policies []domain.Policy
		diff     domain.Diff
		expected bool
	}{
		{
			name: "Matching policy exists",
			policies: []domain.Policy{
				{Path: "path1", ChangeType: "change"},
				{Path: "path2", ChangeType: "add"},
			},
			diff:     domain.Diff{Path: "path1", ChangeType: "change"},
			expected: true,
		},
		{
			name: "Matching policy does not exist",
			policies: []domain.Policy{
				{Path: "path1", ChangeType: "change"},
				{Path: "path2", ChangeType: "add"},
			},
			diff:     domain.Diff{Path: "path3", ChangeType: "change"},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pc := domain.NewPolicyChecker(tc.policies)

			result := pc.Check(tc.diff)

			if result != tc.expected {
				t.Errorf("expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}

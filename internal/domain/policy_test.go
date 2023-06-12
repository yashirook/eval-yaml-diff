package domain_test

import (
	"eval-yaml-diff/internal/domain"
	"reflect"
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
				{Path: ".metadata.name", ChangeType: "change", Recursive: false},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".spec.containers[0].image", ChangeType: "change"},
			expected: true,
		},
		{
			name: "Matching policy does not exist",
			policies: []domain.Policy{
				{Path: ".metadata.name", ChangeType: "change", Recursive: false},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".spec.containers[0].name", ChangeType: "change"},
			expected: false,
		},
		{
			name: "Matching policy does exist(recursive pettern)",
			policies: []domain.Policy{
				{Path: ".metadata", ChangeType: "change", Recursive: true},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".metadata.name", ChangeType: "change"},
			expected: true,
		},
		{
			name: "Matching policy does not exist(recursive pettern)",
			policies: []domain.Policy{
				{Path: ".metadata", ChangeType: "change", Recursive: true},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".spec", ChangeType: "add"},
			expected: false,
		},
		{
			name: "Matching policy does exist(change type all pettern)",
			policies: []domain.Policy{
				{Path: ".metadata", ChangeType: "all", Recursive: true},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".metadata.name", ChangeType: "add"},
			expected: true,
		},
		{
			name: "Matching policy does not exist(recursive unintentionally)",
			policies: []domain.Policy{
				{Path: ".metadata", ChangeType: "all", Recursive: true},
				{Path: ".spec.containers[0].image", ChangeType: "change", Recursive: false},
			},
			diff:     domain.Diff{Path: ".metadatahoge", ChangeType: "add"},
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

func TestCheckAll(t *testing.T) {
	pc := domain.PolicyChecker{
		Policies: []domain.Policy{
			{Path: "path1", ChangeType: domain.ChangeTypeAdd},
			{Path: "path2", ChangeType: domain.ChangeTypeDelete},
			{Path: "path4", ChangeType: domain.ChangeTypeAll},
		},
	}

	diffs := domain.DiffList{
		{Path: "path1", ChangeType: domain.ChangeTypeAdd},
		{Path: "path2", ChangeType: domain.ChangeTypeChange},
		{Path: "path3", ChangeType: domain.ChangeTypeDelete},
		{Path: "path4", ChangeType: domain.ChangeTypeAdd},
	}

	expected := domain.DiffList{
		{Path: "path1", ChangeType: domain.ChangeTypeAdd, Allowed: true},
		{Path: "path2", ChangeType: domain.ChangeTypeChange},
		{Path: "path3", ChangeType: domain.ChangeTypeDelete},
		{Path: "path4", ChangeType: domain.ChangeTypeAdd, Allowed: true},
	}

	result := pc.CheckAll(diffs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CheckAll result mismatch\nExpected: %+v\nActual: %+v", expected, result)
	}
}

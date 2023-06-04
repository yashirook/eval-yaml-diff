package domain

import "fmt"

type Policy struct {
	Path       string     `yaml:"path"`
	ChangeType ChangeType `yaml:"changeType"`
	Recursive  bool       `yaml:"recursive"`
}

type PolicyChecker struct {
	Policies []Policy
}

func NewPolicyChecker(policies []Policy) PolicyChecker {
	return PolicyChecker{
		Policies: policies,
	}
}

func (pc PolicyChecker) CheckAll(diffs DiffList) error {
	evaluatedDiffs := make([]Diff, 0)
	for _, diff := range diffs {
		if ok := pc.Check(diff); ok {
			evaluatedDiffs = append(evaluatedDiffs, diff.Allow())
		}
	}
	fmt.Println(evaluatedDiffs)
	return nil
}

func (pc PolicyChecker) Check(diff Diff) bool {
	for _, policy := range pc.Policies {
		if diff.Path == policy.Path && diff.ChangeType == policy.ChangeType {
			return true
		}
	}

	return false
}

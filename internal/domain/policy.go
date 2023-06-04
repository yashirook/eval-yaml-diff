package domain

import "fmt"

type Policy struct {
	Path       string
	ChangeType ChangeType
	Recursive  bool
}

type PolicyChecker struct {
	Policies []Policy
}

func NewPolicyChecker(policies []Policy) PolicyChecker {
	return PolicyChecker{
		Policies: policies,
	}
}

func (pc PolicyChecker) ScanAll(diffs DiffList) error {
	evaluatedDiffs := make([]Diff, 0)
	for _, diff := range diffs {
		if allowed, err := pc.Check(diff); allowed {
			evaluatedDiffs = append(evaluatedDiffs, diff.Allow())
		} else if err != nil {
			return err
		}
	}
	fmt.Println(evaluatedDiffs)
	return nil
}

func (pc PolicyChecker) Check(diff Diff) (bool, error) {
	for _, policy := range pc.Policies {
		if diff.Path == policy.Path && diff.ChangeType == policy.ChangeType {
			return true, nil
		}
	}

	return false, nil
}
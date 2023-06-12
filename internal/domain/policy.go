package domain

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

func (pc PolicyChecker) CheckAll(diffs DiffList) DiffList {
	newDiffs := make([]Diff, 0)
	for _, diff := range diffs {
		if ok := pc.Check(diff); ok {
			newDiffs = append(newDiffs, diff.Allow())
		} else {
			newDiffs = append(newDiffs, diff)
		}
	}
	return newDiffs
}

func (pc PolicyChecker) Check(diff Diff) bool {
	for _, policy := range pc.Policies {
		matchPath := diff.Path == policy.Path
		if policy.Recursive {
			matchPath = len(diff.Path) >= len(policy.Path) && diff.Path[:len(policy.Path)+1] == policy.Path+"."
		}

		if matchPath && (policy.ChangeType == ChangeTypeAll || diff.ChangeType == policy.ChangeType) {
			return true
		}
	}

	return false
}

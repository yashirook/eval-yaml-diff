package domain

type Config struct {
	AllowedPolicies `yaml:"allowedPolicies"`
}

type AllowedPolicies []Policy

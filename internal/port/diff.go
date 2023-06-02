package port

import (
	"github.com/sters/yaml-diff/yamldiff"
)

type DiffPort interface {
	Get(source, target string) ([]*yamldiff.YamlDiff, error)
}

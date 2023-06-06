package port

import "eval-yaml-diff/internal/domain"

type YAMLDocsPort interface {
	Get(path string) (domain.YamlDocs, error)
}

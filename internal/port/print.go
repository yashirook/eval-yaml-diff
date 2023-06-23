package port

import "eval-yaml-diff/internal/domain"

type PrintPort interface {
	Print(domain.DiffList) error
}

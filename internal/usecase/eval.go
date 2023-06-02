package usecase

import (
	"eval-yaml-diff/internal/port"
	"log"
)

type Eval struct {
	DiffPort port.DiffPort
}

func (e Eval) Do(source, target string) error {
	diff, err := e.DiffPort.Get(source, target)
	if err != nil {
		return err
	}

	log.Println(diff)
	return nil
}

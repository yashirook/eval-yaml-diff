package main

import (
	"eval-yaml-diff/internal/usecase"
	"os"
)

func main() {
	eval := usecase.Eval{}

	err := eval.Do("../example/source.yaml", "../example/target.yaml")
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

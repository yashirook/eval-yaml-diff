package main

import (
	"eval-yaml-diff/internal/gateway"
	"eval-yaml-diff/internal/usecase"
	"os"
)

func main() {
	eval := usecase.Eval{
		DiffPort: &gateway.Diff{},
	}

	err := eval.Do("../example/source.yaml", "../example/target.yaml")
	if err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

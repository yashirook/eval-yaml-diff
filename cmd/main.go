package main

import (
	"eval-yaml-diff/internal/gateway"
	"eval-yaml-diff/internal/usecase"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatalln("Please specify the paths of the YAML files using the first and second arguments.")
	}

	eval := usecase.Eval{
		YAMLDocsPort: &gateway.LocalYAMLDocsGateway{},
	}

	err := eval.Do(args[1], args[2])
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

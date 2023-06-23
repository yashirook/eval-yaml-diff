package main

import (
	"eval-yaml-diff/internal/domain"
	"eval-yaml-diff/internal/gateway"
	"eval-yaml-diff/internal/usecase"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	ExitCodeAllowed        = 0
	ExitCodeSomethingError = 1
	ExitCodeDenied         = 2
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.yaml", "config file for policys")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalln("Please specify the paths of the YAML files using the first and second arguments.")
	}

	os.Exit(Run(args, configPath))
}

func Run(args []string, cp string) int {
	configData, err := ioutil.ReadFile(cp)
	if err != nil {
		log.Println("Failed to read config file")
		return ExitCodeSomethingError
	}

	var config domain.Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Println("Failed to parse config file. Please check file format.")
		return ExitCodeSomethingError
	}

	eval := usecase.Eval{
		YAMLDocsPort: &gateway.LocalYAMLDocsGateway{},
		PrintPort:    &gateway.PrintGateway{},
		Config:       config,
	}

	err = eval.Do(args[0], args[1])
	if err == usecase.DifferentDocumentNumberError || err == usecase.DeniedDiffExistError {
		fmt.Println(err)
		return ExitCodeDenied
	}
	if err != nil {
		fmt.Println(err)
		return ExitCodeSomethingError
	}

	return ExitCodeAllowed
}

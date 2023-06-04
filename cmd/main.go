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

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("Failed to read config file")
	}

	var config domain.Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalln("Failed to parse config file. Please check file format.")
	}

	fmt.Println(config)

	Run(args, config)
}

func Run(args []string, config domain.Config) {
	eval := usecase.Eval{
		YAMLDocsPort: &gateway.LocalYAMLDocsGateway{},
		Config:       config,
	}

	err := eval.Do(args[0], args[1])
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

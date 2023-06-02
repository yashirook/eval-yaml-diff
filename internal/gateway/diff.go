package gateway

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sters/yaml-diff/yamldiff"
)

type Diff struct {
}

func (d Diff) Get(source, target string) ([]*yamldiff.YamlDiff, error) {
	sourceRawYaml, err := yamldiff.Load(load(source))
	if err != nil {
		return nil, err
	}

	targetRawYaml, err := yamldiff.Load(load(target))
	if err != nil {
		return nil, err
	}

	yamlDiffs := yamldiff.Do(sourceRawYaml, targetRawYaml)
	for _, d := range yamlDiffs {
		fmt.Println(d.Dump(), "")
	}
	return yamlDiffs, nil
}

func load(f string) string {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Printf("%+v, %s", err, f)

		return ""
	}

	b, err := io.ReadAll(file)
	if err != nil {
		log.Printf("%+v, %s", err, f)

		return ""
	}

	return string(b)
}

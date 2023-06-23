package gateway

import (
	"bytes"
	"eval-yaml-diff/internal/domain"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type LocalYAMLDocsGateway struct{}

func (l LocalYAMLDocsGateway) Get(path string) (domain.YamlDocs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	yamlDocs, err := decodeYaml(data)
	if err != nil {
		return nil, err
	}

	return yamlDocs, nil
}

func decodeYaml(data []byte) (domain.YamlDocs, error) {
	yamls := make(domain.YamlDocs, 0)

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	for {
		var data domain.Yaml

		err := decoder.Decode(&data)
		if err == io.EOF {
			break
		}
		if err != nil {
			return yamls, err
		}

		yamls = append(yamls, data)
	}
	return yamls, nil

}

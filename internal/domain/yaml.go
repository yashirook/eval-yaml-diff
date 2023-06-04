package domain

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v2"
)

type Yaml interface{}

type YamlDocs []Yaml

func NewYamlDocs(raw []byte) (YamlDocs, error) {
	yamls := make(YamlDocs, 0)

	decoder := yaml.NewDecoder(bytes.NewReader(raw))
	for {
		var data Yaml

		if err := decoder.Decode(&data); err != nil {
			if err == io.EOF {
				break
			}
			return yamls, err
		}

		yamls = append(yamls, data)
	}

	return yamls, nil
}

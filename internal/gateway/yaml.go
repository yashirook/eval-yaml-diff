// gatewayは外部から取った情報をdomainに変換する
package gateway

import (
	"eval-yaml-diff/internal/domain"
	"io/ioutil"
)

type LocalYAMLDocsGateway struct{}

// driverでioutilをスタブ化させるとテスト書きやすいかも
func (l LocalYAMLDocsGateway) Get(path string) (domain.YamlDocs, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	yamlDocs, err := domain.NewYamlDocs(data)
	if err != nil {
		return nil, err
	}

	return yamlDocs, nil
}

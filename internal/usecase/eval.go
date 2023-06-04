package usecase

import (
	"eval-yaml-diff/internal/domain"
	"fmt"
	"io/ioutil"
	"log"
)

type Eval struct {
}

func (e Eval) Do(source, target string) error {
	// TODO: YAML Docsのスライスを取得するところをGatewayのレイヤにまとめる。
	oldYAMLData, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	newYAMLData, err := ioutil.ReadFile(target)
	if err != nil {
		return err
	}

	oldYamlDocs, err := domain.NewYamlDocs(oldYAMLData)
	if err != nil {
		return err
	}

	newYamlDocs, err := domain.NewYamlDocs(newYAMLData)
	if err != nil {
		return err
	}

	// TODO: ドキュメントの数が違う場合にいい感じに処理できるようにする
	if len(oldYamlDocs) != len(newYamlDocs) {
		fmt.Println("Different number of yaml documents")
		return err
	}

	diffFinder := domain.DiffFinder{}
	// TODO: マルチドキュメントの場合にいい感じに差分を取得できるようにする
	for i, oldYamlDoc := range oldYamlDocs {
		diff, err := diffFinder.Find(oldYamlDoc, newYamlDocs[i])
		if err != nil {
			return err
		}
		log.Println(diff)

	}

	return nil
}

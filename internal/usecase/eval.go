package usecase

import (
	"eval-yaml-diff/internal/domain"
	"eval-yaml-diff/internal/port"
	"fmt"
	"log"
)

type Eval struct {
	YAMLDocsPort port.YAMLDocsPort
}

func (e Eval) Do(baseline, new string) error {
	baseYamlDocs, err := e.YAMLDocsPort.Get(baseline)
	if err != nil {
		return err
	}
	newYamlDocs, err := e.YAMLDocsPort.Get(new)
	if err != nil {
		return err
	}

	// TODO: ドキュメントの数が違う場合にいい感じに処理できるようにする
	if len(baseYamlDocs) != len(newYamlDocs) {
		fmt.Println("Different number of yaml documents")
		return err
	}

	diffFinder := domain.DiffFinder{}
	diffs := make(domain.DiffList, 0)
	// TODO: マルチドキュメントの場合にいい感じに差分を取得できるようにする
	for i, baseYamlDoc := range baseYamlDocs {
		diff, err := diffFinder.Find(baseYamlDoc, newYamlDocs[i])
		if err != nil {
			return err
		}
		log.Println(diff)
		diffs = append(diffs, diff...)
	}

	policies := []domain.Policy{
		{Path: ".spec.template.metadata.labels.version", ChangeType: domain.ChangeTypeAdd, Recursive: false},
	}
	pc := domain.NewPolicyChecker(policies)
	pc.ScanAll(diffs)

	return nil
}

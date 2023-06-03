package usecase

import (
	"eval-yaml-diff/internal/domain"
	"io/ioutil"
	"log"
)

type Eval struct {
}

func (e Eval) Do(source, target string) error {
	oldYAMLData, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	newYAMLData, err := ioutil.ReadFile(target)
	if err != nil {
		return err
	}

	diffFinder := domain.DiffFinder{}

	oldYamlDocs, err := domain.NewYamlDocs(oldYAMLData)
	if err != nil {
		return err
	}

	newYamlDocs, err := domain.NewYamlDocs(newYAMLData)
	if err != nil {
		return err
	}

	for i, oldYamlDoc := range oldYamlDocs {
		diff, err := diffFinder.Find(oldYamlDoc, newYamlDocs[i])
		if err != nil {
			return err
		}
		log.Println(diff)

	}

	// if len(oldData) != len(newData) {
	// 	fmt.Println("Different number of yaml documents")
	// 	return err
	// }

	return nil
}

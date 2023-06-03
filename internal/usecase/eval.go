package usecase

import (
	"eval-yaml-diff/internal/domain"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
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

	var oldData interface{}
	var newData interface{}

	if err := yaml.Unmarshal(oldYAMLData, &oldData); err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(newYAMLData, &newData); err != nil {
		log.Fatal(err)
	}

	diffFinder := domain.DiffFinder{}
	diff, err := diffFinder.Find(oldData, newData)
	if err != nil {
		return err
	}

	log.Println(diff)
	return nil
}

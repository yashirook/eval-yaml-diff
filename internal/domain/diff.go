package domain

import (
	"strconv"
)

type Diff struct {
	ChangeType ChangeType // 変更の種類（追加、削除、変更など）
	Path       string     // 変更が発生した場所（パス）
	Allowed    bool
}

type ChangeType string

const (
	ChangeTypeAdd    ChangeType = "add"
	ChangeTypeDelete ChangeType = "delete"
	ChangeTypeChange ChangeType = "change"
	ChangeTypeAll    ChangeType = "all"
)

func newDiff(changeType ChangeType, path string) Diff {
	return Diff{ChangeType: changeType, Path: path}
}

func (d Diff) Allow() Diff {
	d.Allowed = true
	return d
}

type DiffList []Diff

type DiffFinder struct{}

func (df DiffFinder) Find(oldData, newData interface{}) (DiffList, error) {
	var diffs DiffList

	findDifferences(oldData, newData, "", &diffs)

	return diffs, nil
}

func findDifferences(oldData, newData interface{}, path string, diffs *DiffList) {
	switch oldValue := oldData.(type) {
	case map[interface{}]interface{}:
		newValue, ok := newData.(map[interface{}]interface{})
		if !ok {
			// データの型が一致しない場合、変更が発生したと判断してDiffを追加する
			*diffs = append(*diffs, Diff{ChangeType: ChangeTypeChange, Path: path})
			return
		}

		for key, value := range oldValue {
			newPath := path + "." + key.(string)
			if newValue[key] == nil {
				// 新しいデータにキーが存在しない場合、削除が発生したと判断してDiffを追加する
				*diffs = append(*diffs, Diff{ChangeType: ChangeTypeDelete, Path: newPath})
			} else {
				findDifferences(value, newValue[key], newPath, diffs)
			}
		}

		for key := range newValue {
			if oldValue[key] == nil {
				// 古いデータにキーが存在しない場合、追加が発生したと判断してDiffを追加する
				newPath := path + "." + key.(string)
				*diffs = append(*diffs, Diff{ChangeType: ChangeTypeAdd, Path: newPath})
			}
		}

	case []interface{}:
		newValue, ok := newData.([]interface{})
		if !ok {
			// データの型が一致しない場合、変更が発生したと判断してDiffを追加する
			*diffs = append(*diffs, Diff{ChangeType: ChangeTypeChange, Path: path})
			return
		}

		// スライスの要素数が異なる場合、変更が発生したと判断してDiffを追加する
		if len(oldValue) != len(newValue) {
			*diffs = append(*diffs, Diff{ChangeType: ChangeTypeChange, Path: path})
			return
		}

		for i := range oldValue {
			findDifferences(oldValue[i], newValue[i], path+"["+strconv.Itoa(i)+"]", diffs)
		}

	default:
		if oldValue != newData {
			// 値が異なる場合、変更が発生したと判断してDiffを追加する
			*diffs = append(*diffs, Diff{ChangeType: ChangeTypeChange, Path: path})
		}
	}

}

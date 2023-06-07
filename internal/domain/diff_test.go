package domain_test

import (
	"eval-yaml-diff/internal/domain"
	"reflect"
	"testing"
)

func TestFindDifferences(t *testing.T) {
	testCases := []struct {
		Name          string
		OldData       interface{}
		NewData       interface{}
		ExpectedDiffs domain.DiffList
	}{
		{
			Name:          "Map_DifferentValue",
			OldData:       map[interface{}]interface{}{"key1": "value1"},
			NewData:       map[interface{}]interface{}{"key1": "value2"},
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeChange, Path: ".key1"}},
		},
		{
			Name:          "Map_Deletion",
			OldData:       map[interface{}]interface{}{"key1": "value1"},
			NewData:       map[interface{}]interface{}{},
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeDelete, Path: ".key1"}},
		},
		{
			Name:          "Map_Addition",
			OldData:       map[interface{}]interface{}{},
			NewData:       map[interface{}]interface{}{"key1": "value1"},
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeAdd, Path: ".key1"}},
		},
		{
			Name:          "Map_Same",
			OldData:       map[interface{}]interface{}{"key1": "value1"},
			NewData:       map[interface{}]interface{}{"key1": "value1"},
			ExpectedDiffs: nil,
		},
		{
			Name:          "Slice_DifferentValue",
			OldData:       []interface{}{"value1", "value2", "value3"},
			NewData:       []interface{}{"value1", "valueX", "value3"},
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeChange, Path: "[1]"}},
		},
		{
			Name:          "Slice_DifferentLength",
			OldData:       []interface{}{"value1", "value2", "value3"},
			NewData:       []interface{}{"value1", "value2"},
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeChange, Path: ""}},
		},
		{
			Name:          "Slice_Empty",
			OldData:       []interface{}{},
			NewData:       []interface{}{},
			ExpectedDiffs: nil,
		},
		{
			Name:          "Value_Different",
			OldData:       "value1",
			NewData:       "value2",
			ExpectedDiffs: domain.DiffList{{ChangeType: domain.ChangeTypeChange, Path: ""}},
		},
		{
			Name:          "Value_Same",
			OldData:       "value",
			NewData:       "value",
			ExpectedDiffs: nil,
		},
	}

	for _, tc := range testCases {
		df := domain.DiffFinder{}
		t.Run(tc.Name, func(t *testing.T) {
			diffs, err := df.Find(tc.OldData, tc.NewData)
			if err != nil {
				t.Errorf("Find returned an error: %v", err)
			}

			if !reflect.DeepEqual(diffs, tc.ExpectedDiffs) {
				t.Errorf("findDifferences result mismatch\nExpected: %+v\nActual: %+v", tc.ExpectedDiffs, diffs)
			}
		})
	}
}

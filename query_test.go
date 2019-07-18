package binder

import (
	"reflect"
	"testing"
)

func TestBindFromQuery(t *testing.T) {
	type test struct {
		StrField        string `json:"str_field"`
		IntField        int    `json:"int_field"`
		FieldWithoutTag int
		IntSliceField   []int    `json:"int_slice_field"`
		StrSliceField   []string `json:"str_slice_field"`
		BoolField       bool     `json:"bool_field"`
	}

	expected := test{
		StrField:      "some string",
		IntField:      10,
		IntSliceField: []int{1, 2, 3, 4, 5},
		StrSliceField: []string{"one", "two"},
		BoolField:     true,
	}

	source := map[string][]string{}
	source["int_field"] = []string{"10"}
	source["str_field"] = []string{"some string"}
	source["int_slice_field"] = []string{"1,2,3,4,5"}
	source["str_slice_field"] = []string{"one,two"}
	source["bool_field"] = []string{"true"}

	var actual test

	err := BindFromQuery(source, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected not equal expected \n %+v \n %+v", expected, actual)
	}
	//TODO provide more tests
}

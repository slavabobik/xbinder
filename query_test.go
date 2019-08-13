package binder

import (
	"reflect"
	"strings"
	"testing"
)

func TestFromQuery_Bind(t *testing.T) {
	type s struct {
		Int             int    `json:"int"`
		Int8            int8   `json:"int_8"`
		Int16           int16  `json:"int_16"`
		Int32           int32  `json:"int_32"`
		Int64           int64  `json:"int_64"`
		String          string `json:"string"`
		Bool            bool   `json:"bool"`
		CapitalizedBool bool   `json:"capitalized_bool"`
	}

	want := s{
		Int:             4,
		Int8:            8,
		Int16:           16,
		Int32:           32,
		Int64:           64,
		String:          "value",
		Bool:            true,
		CapitalizedBool: true,
	}

	query := map[string][]string{
		"int":              {"4"},
		"int_8":            {"8"},
		"int_16":           {"16"},
		"int_32":           {"32"},
		"int_64":           {"64"},
		"string":           {"value"},
		"bool":             {"true"},
		"capitalized_bool": {"True"},
	}

	var got s
	err := FromQuery(&got, query)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected value, got:%+v, want:%+v", got, want)
	}
}

func TestFromQuery_BindSlice(t *testing.T) {
	type s struct {
		IntSlice         []int    `json:"int_slice"`
		StringSlice      []string `json:"string_slice"`
		MissedValueSlice []int    `json:"missed_value_slice"`
	}

	want := s{
		IntSlice:         []int{1, 2, 3, 4},
		StringSlice:      []string{"one", "two"},
		MissedValueSlice: []int{1, 3, 4},
	}

	query := map[string][]string{
		"int_slice":          {"1,2,3,4"},
		"string_slice":       {"one,two", "three"},
		"missed_value_slice": {"1,,3,4,"},
	}

	var got s
	err := FromQuery(&got, query)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected value, got:%+v, want:%+v", got, want)
	}
}

func TestFromQuery_BindIntSliceError(t *testing.T) {
	query := map[string][]string{"f": {"str"}}
	s := struct {
		F []int `json:"f"`
	}{}

	err := FromQuery(&s, query)
	if err == nil {
		t.Fatal("no error")
	}

	if err.Error() != "can't parse value: str" {
		t.Fatalf("unexpected error: %#v", err)
	}
}

func TestFromQuery_BindError_UnaddressableField(t *testing.T) {
	query := map[string][]string{"f": {"str"}}
	s := struct {
		f []int `json:"f"`
	}{}

	err := FromQuery(&s, query)
	if err == nil {
		t.Fatal("no error")
	}

	if err.Error() != "field f is not addressable" {
		t.Fatalf("unexpected error: %#v", err)
	}
}

func TestFromQuery_BindTypeError(t *testing.T) {
	err := FromQuery(struct{}{}, map[string][]string{})
	if err == nil || err.Error() != "invalid interface type" {
		t.Fatal("no expected error")
	}
}

func TestFromQuery_UnsupportedTypes(t *testing.T) {
	f := func(s interface{}) {
		t.Helper()
		err := FromQuery(s, map[string][]string{"f": {"1"}})
		if err == nil {
			t.Fatal("no error")
		}

		if !strings.Contains(err.Error(), "unsupported type") {
			t.Fatalf("no expected error: %#v", err)
		}
	}

	s1 := struct {
		F uint `json:"f"`
	}{}
	f(&s1)

	s2 := struct {
		F uint8 `json:"f"`
	}{}
	f(&s2)

	s3 := struct {
		F uint16 `json:"f"`
	}{}
	f(&s3)

	s4 := struct {
		F uint32 `json:"f"`
	}{}
	f(&s4)

	s5 := struct {
		F uint64 `json:"f"`
	}{}
	f(&s5)

	s6 := struct {
		F float32 `json:"f"`
	}{}
	f(&s6)

	s7 := struct {
		F float64 `json:"f"`
	}{}
	f(&s7)

	s8 := struct {
		F *int `json:"f"`
	}{}
	f(&s8)
}

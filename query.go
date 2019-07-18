package binder

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const defaultSeparator = ","

//TODO add comment
func BindFromQuery(source url.Values, v interface{}) error {
	value := reflect.ValueOf(v).Elem()

	for i := 0; i < value.NumField(); i++ {
		tag, ok := value.Type().Field(i).Tag.Lookup("json")
		if !ok || tag == "" || tag == "-" {
			continue
		}

		queryParamValue := source.Get(tag)
		if queryParamValue == "" {
			continue
		}

		switch value.Field(i).Interface().(type) {
		case int:
			p, err := strconv.Atoi(queryParamValue)
			if err != nil {
				return err
			}
			value.Field(i).SetInt(int64(p))
		case string:
			value.Field(i).SetString(queryParamValue)
		case bool:
			if queryParamValue == "true" {
				value.Field(i).SetBool(true)
			}
		case []int:
			values := strings.Split(queryParamValue, defaultSeparator)
			if len(values) == 0 {
				continue
			}
			//TODO remove duplicate code
			value.Field(i).Set(reflect.MakeSlice(value.Field(i).Type(), len(values), len(values)))
			for idx, s := range values {
				n, err := strconv.Atoi(s)
				if err != nil {
					return err
				}

				value.Field(i).Index(idx).Set(reflect.ValueOf(n))
			}

		case []string:
			values := strings.Split(queryParamValue, defaultSeparator)
			if len(values) == 0 {
				continue
			}
			value.Field(i).Set(reflect.MakeSlice(value.Field(i).Type(), len(values), len(values)))
			for idx, s := range values {
				value.Field(i).Index(idx).Set(reflect.ValueOf(s))
			}

			log.Println(queryParamValue)
		default:
			return fmt.Errorf("unsupported type")
		}
	}

	return nil
}

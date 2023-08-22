package utils

import (
	"errors"
	"fmt"
	"reflect"
)

// 读取结构体的指定字段
func GetStructField(input interface{}, key string) (value interface{}, err error) {
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)
	if rt.Kind() != reflect.Struct {
		return value, errors.New("输入参数必须为struct")
	}

	keyExist := false
	for i := 0; i < rt.NumField(); i++ {
		curField := rv.Field(i)
		if rt.Field(i).Name == key {
			switch curField.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				keyExist = true
				value = curField.Interface()
			default:
				return value, errors.New("key must be int float or string")
			}
		}
	}

	if !keyExist {
		return value, errors.New(fmt.Sprintf("key %s not found in %s's field", key, rt))
	}
	return
}

// 读取slice的字段，并返回slice
func GetSliceStructField(input interface{}, key string) (value []interface{}, err error) {
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)
	if rt.Kind() != reflect.Slice {
		return value, errors.New("输入参数必须为slice")
	}
	var val []interface{}

	for i := 0; i < rv.Len(); i++ {
		sv := rv.Index(i)
		st := sv.Type()
		if st.Kind() != reflect.Struct {
			return value, errors.New("输入参数必须为struct")
		}

		for j := 0; j < sv.NumField(); j++ {
			if st.Field(j).Name == key {
				val = append(val, sv.Field(j).Interface())
			}
		}
	}
	return val, nil
}

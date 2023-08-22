package utils

import (
	"fmt"
	"reflect"
)

func GetStructFieldsByReflect(obj interface{}, jsonNames *[]string, fieldNames *[]string) error {
	s := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	if s.Kind() == reflect.Ptr {
		obj = reflect.Value.Elem(s).Interface()
		s = reflect.ValueOf(obj)
		t = reflect.TypeOf(obj)
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("obj is not struct type")
	}
	for i := 0; i < s.NumField(); i++ {
		tField := t.Field(i)
		if tField.Type.Kind() == reflect.Struct && tField.Anonymous {
			//fmt.Println("13", tField.Name, tField.Type.Kind(), tField.Anonymous)
			if err := GetStructFieldsByReflect(s.Field(i).Interface(), jsonNames, fieldNames); err != nil {
				return err
			}
		} else {
			//fmt.Println("14", tField.Name, tField.Type.Kind(), tField.Anonymous)
			name := tField.Name
			// 获取struct的tag
			tag := tField.Tag.Get("json")
			if len(tag) > 0 {
				name = tag
			}
			if tag != "-" {
				for idx, n := range *jsonNames {
					n := n
					if n == name {
						*jsonNames = append((*jsonNames)[:idx], (*jsonNames)[idx+1:]...)
					}
				}
				for idx, n := range *fieldNames {
					n := n
					if n == tField.Name {
						*fieldNames = append((*fieldNames)[:idx], (*fieldNames)[idx+1:]...)
					}
				}
				*jsonNames = append(*jsonNames, name)
				*fieldNames = append(*fieldNames, tField.Name)
			}

		}

	}
	return nil
}

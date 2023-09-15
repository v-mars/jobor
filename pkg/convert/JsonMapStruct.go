package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func StructToMapOut(m interface{}, out *map[string]interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return err
	}
	var t = reflect.TypeOf(m)
	var v = reflect.ValueOf(m)
	if v.Kind() == reflect.Ptr {
		m = reflect.Value.Elem(v).Interface()
		v = reflect.ValueOf(m)
		t = reflect.TypeOf(m)
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := t.Field(i)
		if field.CanInterface() {
			var fieldName = strings.ToLower(tField.Name)
			var value = field.Interface()
			var keepData = tField.Tag.Get("keep_data")
			var toStr = tField.Tag.Get("to_str")
			var toTime = tField.Tag.Get("to_time")
			var ignore = tField.Tag.Get("ignore")
			var jsonTag = tField.Tag.Get("json")
			if len(strings.Split(jsonTag, ",")) > 0 {
				fieldName = strings.Split(jsonTag, ",")[0]
			}
			_, ok := (*out)[fieldName]
			//fmt.Println("to :", jsonTag, ignore, tField.Name, field.String())
			if ok && keepData == "yes" {
				(*out)[fieldName] = value
			}
			if ok && toStr == "yes" {
				bytes, err := json.Marshal(value)
				if err != nil {
					return err
				}
				(*out)[fieldName] = string(bytes)
				//fmt.Println("to map fieldName3:", fieldName, value)
			}
			if ignore == "yes" {
				delete(*out, fieldName)
			}
			if toTime != "" && ToInt64(value) != 0 {
				(*out)[fieldName] = time.UnixMilli(ToInt64(value))
			}
			if ignore == "no" {
				(*out)[fieldName] = ""
			}
		}
	}
	return nil
}

func StructToMap(m interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}
	tmp, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	//fmt.Println("StructToMap jsonStr:", string(tmp))
	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return nil, err
	}
	var t = reflect.TypeOf(m)
	var v = reflect.ValueOf(m)
	if v.Kind() == reflect.Ptr {
		m = reflect.Value.Elem(v).Interface()
		v = reflect.ValueOf(m)
		t = reflect.TypeOf(m)
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tField := t.Field(i)
		if field.CanInterface() {
			var fieldName = strings.ToLower(tField.Name)
			var value = field.Interface()
			var keepData = tField.Tag.Get("keep_data")
			var toStr = tField.Tag.Get("to_str")
			var ignore = tField.Tag.Get("ignore")
			var toTime = tField.Tag.Get("to_time")
			var jsonTag = tField.Tag.Get("json")
			if len(strings.Split(jsonTag, ",")) > 0 {
				fieldName = strings.Split(jsonTag, ",")[0]
			}
			_, ok := out[fieldName]
			//fmt.Println("to map fieldName1:", jsonTag, fieldName, keepData, ok)
			if ok && keepData == "yes" {
				//fmt.Println("to map fieldName2:", fieldName, value)
				out[fieldName] = value
			}
			if ok && toStr == "yes" {
				bytes, err := json.Marshal(value)
				if err != nil {
					return nil, err
				}
				out[fieldName] = string(bytes)
				//fmt.Println("to map fieldName3:", fieldName, value)
			}
			if ignore == "yes" {
				delete(out, fieldName)
			}
			if toTime != "" && ToInt64(value) != 0 {
				out[fieldName] = time.UnixMilli(ToInt64(value))
			}
			if ignore == "no" {
				out[fieldName] = ""
			}
		}
	}
	return out, nil
}

func StructToMapSlice(m interface{}) ([]map[string]interface{}, error) {
	var out []map[string]interface{}
	tmp, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func StructToMapViaReflect(m interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	elem := reflect.ValueOf(&m).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		r[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return r
}

// StructToMapByReflect ToMap 结构体转为Map[string]interface{}
func StructToMapByReflect(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	relType := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := relType.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		} else {
			out[relType.Field(i).Name] = v.Interface()
		}
	}
	return out, nil
}

func AnyToAny(src interface{}, out interface{}) error {
	tmp, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return err
	}
	return nil
}

package convert

import (
	"encoding/json"
	"reflect"
)


func StructToMapOut(m interface{}, out *map[string]interface{}) error {
	tmp, err := json.Marshal(m)
	if err!=nil{
		return err
	}
	err = json.Unmarshal(tmp, &out)
	if err!=nil{
		return err
	}
	return nil
}

func StructToMap(m interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}
	tmp, err := json.Marshal(m)
	if err!=nil{
		return nil,err
	}
	err = json.Unmarshal(tmp, &out)
	if err!=nil{
		return nil,err
	}
	return out, nil
}

func StructToMapSlice(m interface{}) ([]map[string]interface{}, error) {
	var out []map[string]interface{}
	tmp, err := json.Marshal(m)
	if err!=nil{
		return nil,err
	}
	err = json.Unmarshal(tmp, &out)
	if err!=nil{
		return nil,err
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
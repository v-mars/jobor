package utils

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
)

// 定义JSON操作
var (
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal       = json.Marshal
	JSONUnmarshal     = json.Unmarshal
	JSONMarshalIndent = json.MarshalIndent
	JSONNewDecoder    = json.NewDecoder
	JSONNewEncoder    = json.NewEncoder
)

// JSONMarshalToString JSON编码为字符串
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

func PrettyJson(data interface{}) (string, error) {
	const (
		empty = ""
		tab   = "\t"
	)
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

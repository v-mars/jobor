package utils

import (
	"bytes"
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

// 定义JSON操作
var (
	Json              = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal       = Json.Marshal
	JSONUnmarshal     = Json.Unmarshal
	JSONMarshalIndent = Json.MarshalIndent
	JSONNewDecoder    = Json.NewDecoder
	JSONNewEncoder    = Json.NewEncoder
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
	encoder := Json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func StructToJsonFormatted(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(data)
}

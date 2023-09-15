package convert

import (
	"bytes"
	"encoding/json"
)

// StructToJsonBytes struct 转成json防止特殊字符转义
func StructToJsonBytes(src interface{}) (*bytes.Buffer, error) {
	bf := bytes.NewBuffer([]byte{})
	jE := json.NewEncoder(bf)
	if err := jE.Encode(src); err != nil {
		return bf, err
	}
	return bf, nil
}

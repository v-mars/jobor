package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"strings"
)

func AnyToAny(src interface{}, des interface{}) error {
	bytesData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytesData, des)
	return err
}

func AnyToAnyV2(src interface{}, des interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(src)
	err = json.NewDecoder(&buf).Decode(&des)
	return err
}

func GobAnyToAny(src interface{}, des interface{}) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(src)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(bytes.NewReader(buf.Bytes()))
	//dec := gob.NewDecoder(&buf)
	err = decoder.Decode(des)
	return err
}

func Uri2map(uri string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if len(uri) < 1 { // 空字符串
		return m, errors.New("uri is none")
	}
	if uri[0:1] == "?" { // 有没有包含？,有的话忽略。
		uri = uri[1:]
	}

	// 首先分解&（sessionid=22222&token=3333 )变成
	// sessionid=222222 和 token=3333
	pars := strings.Split(uri, "&")
	for _, par := range pars {
		// 然后分解 sessionid=222222 和 token=3333
		parkv := strings.Split(par, "=")
		m[parkv[0]] = parkv[1] // 等号前面是key,后面是value
	}
	//log.Println(m)
	return m, nil
}

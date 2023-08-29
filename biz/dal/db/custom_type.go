package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
	"log"
	"time"
)

type IntArray []int

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *IntArray) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %s", value)
	}
	var err error
	result := IntArray{}
	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &result)
	}
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j IntArray) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type StringArray []string

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *StringArray) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %s", value)
	}
	var err error
	result := StringArray{}
	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &result)
	}
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j StringArray) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type IntNestArray [][]int

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *IntNestArray) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %s", value)
	}
	var err error
	result := IntNestArray{}
	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &result)
	}
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j IntNestArray) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type Bool bool

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *Bool) Scan(value interface{}) error {
	*j = Bool(false)
	if value != nil {
		*j = Bool(convert.ToBool(value))
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j Bool) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	value := convert.ToString(j)
	return value, nil
}

type Int int

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *Int) Scan(value interface{}) error {
	*j = 0
	if value != nil {
		*j = Int(convert.ToInt(value))
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j Int) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	value := convert.ToString(j)
	return value, nil
}

type MillisTime string

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
//func (j MillisTime) MarshalJSON() ([]byte, error) {
//	return []byte(strconv.Quote(string(j))), nil
//}

// UnmarshalJSON 反序列化
func (j *MillisTime) UnmarshalJSON(data []byte) error {
	*j = MillisTime(data)
	return nil
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *MillisTime) Scan(value interface{}) error {
	*j = "0"
	if value != nil {
		*j = MillisTime((time.Millisecond * time.Duration(convert.ToInt64(value))).String())
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j MillisTime) Value() (driver.Value, error) {
	if len(j) == 0 || j == "0" || j == "0s" {
		return "0", nil
	}
	duration, err := time.ParseDuration(string(j))
	if err != nil {
		log.Printf("MillisTime time.ParseDuration failed: %s", err)
		return "0", nil
	}
	return convert.ToString(duration.Milliseconds()), nil
}

type SecTime string

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
//func (j SecTime) MarshalJSON() ([]byte, error) {
//
//	//formatted := ""
//	//if len(j.Value()) > 0 && j.String() != "null" {
//	//	formatted = fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
//	//}
//	return []byte(strconv.Quote(string(j))), nil
//}

// UnmarshalJSON 反序列化
func (j *SecTime) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	*j = SecTime((time.Millisecond * time.Duration(convert.ToFloat64(string(data))*1000)).String())
	return nil
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *SecTime) Scan(value interface{}) error {
	*j = "0s"
	if value != nil {
		*j = SecTime((time.Millisecond * time.Duration(convert.ToFloat64(value)*1000)).String())
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j SecTime) Value() (driver.Value, error) {
	if len(j) == 0 || j == "0" || j == "0s" {
		return "0", nil
	}
	duration, err := time.ParseDuration(string(j))
	if err != nil {
		log.Printf("SecTime time.ParseDuration failed: %s", err)
		return "0", nil
	}
	return convert.ToString(duration.Seconds()), nil
}

type AesStr string

var DefaultAesKey = "Joborb5Zxxxxx06GorMa"

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *AesStr) Scan(value interface{}) error {
	//*j = ""
	var str = convert.ToString(value)
	if len(str) > 0 {
		*j = AesStr(utils.DeTxtByAes(str, DefaultAesKey))
	}
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j AesStr) Value() (driver.Value, error) {
	if len(j) == 0 || len(string(j)) == 0 {
		return "", nil
	}
	return utils.EnTxtByAes(string(j), DefaultAesKey), nil
}

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := ""
	if len(t.String()) > 0 && t.String() != "null" {
		formatted = fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	}
	return []byte(formatted), nil
}

// UnmarshalJSON 反序列化
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	var dataStr = string(data)
	if string(data) == "null" || string(data) == "0001-01-01 00:00:00" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	var format = time.RFC3339
	if len(dataStr) <= 21 {
		format = "2006-01-02 15:04:05"
	}

	// Fractional seconds are handled implicitly by Parse.
	var tt time.Time
	if len(dataStr) == 13 {
		tt = time.UnixMilli(convert.ToInt64(dataStr))
	} else if len(dataStr) == 10 {
		tt = time.Unix(convert.ToInt64(dataStr), 0)
	} else {
		tt, err = time.Parse(`"`+format+`"`, string(data))
	}
	fmt.Println("jsontime:", tt, tt.String(), len(dataStr))
	//tt, err2 := time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local)
	if err != nil {
		return err
	} else {
		*t = JSONTime{tt}
	}
	return err
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// DateTime format json time field by myself
type DateTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// UnmarshalJSON 反序列化
func (t *DateTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	var dataStr = string(data)
	if dataStr == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	var layout = time.RFC3339
	if len(dataStr) <= 12 {
		layout = "2006-01-02"
	}
	var tt time.Time
	if len(dataStr) == 13 {
		tt = time.UnixMilli(convert.ToInt64(dataStr))
	} else {
		tt, err = time.Parse(`"`+layout+`"`, string(data))
	}
	//tt, err2 := time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local)
	if err != nil {
		return err
	} else {
		*t = DateTime{tt}
	}
	return err
}

// Value insert timestamp into mysql need this function.
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type ZhToCode struct {
	Zh   string `json:"zh"`
	Code string `json:"code"`
}

type JsonMap map[string]interface{}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JsonMap) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %s", value)
	}
	var err error
	result := JsonMap{}
	if len(bytes) != 0 {
		err = json.Unmarshal(bytes, &result)
	}
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j JsonMap) Value() (driver.Value, error) {
	if &j == nil || len(j) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

type EncJsonMap map[string]string

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *EncJsonMap) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to unmarshal JSONB value: %s", value)
	}
	var err error
	result := EncJsonMap{}
	if len(bytes) != 0 {
		var str = convert.ToString(value)
		if len(str) > 0 {
			err = json.Unmarshal([]byte(utils.DeTxtByAes(str, DefaultAesKey)), &result)
		}
	}
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j EncJsonMap) Value() (driver.Value, error) {
	if &j == nil || len(j) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	var r = utils.EnTxtByAes(string(bytes), DefaultAesKey)
	//fmt.Println("EncJsonMap value:", string(bytes), "enc:", r)
	//return string(bytes), err
	return r, err
}

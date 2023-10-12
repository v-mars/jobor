package task

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func (u *PutTaskReq) GetParentIdsInt() []int {
	var arrayInt = make([]int, 0)
	if u.ParentTaskIds != nil {
		for _, v := range u.ParentTaskIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}
func (u *PutTaskReq) GetChildIdsInt() []int {
	var arrayInt = make([]int, 0)
	if u.ChildTaskIds != nil {
		for _, v := range u.ChildTaskIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}

func (u *PutTaskReq) GetOwnerIdsInt() []int {
	var arrayInt = make([]int, 0)
	if u.OwnerIds != nil {
		for _, v := range u.OwnerIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *KeepLog) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := KeepLog{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j KeepLog) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

//type TaskData struct {
//	Code string `json:"data"`
//	Api  struct {
//		Url         string              `json:"url"`
//		Method      string              `json:"method"`
//		ContentType string              `json:"content_type"`
//		Payload     string              `json:"payload"`
//		Header      []map[string]string `json:"header"`
//		Forms       []map[string]string `json:"forms"`
//	} `json:"api"`
//}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *TaskData) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := TaskData{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j TaskData) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

//type Notify struct {
//	Email struct {
//		Enable    bool     `json:"enable"`
//		Receivers []string `json:"receivers"` // 多个邮箱地址以逗号分割
//	} `json:"email"`
//	Webhook struct {
//		Enable bool     `json:"enable"`
//		Urls   []string `json:"urls"` // 多个api url以逗号分割
//	} `json:"webhook"`
//	// 多个飞书webhook以逗号分割
//	Lark struct {
//		Enable   bool `json:"enable"`
//		Webhooks []struct {
//			WebhookUrl string `json:"webhook_url"`
//			Secret     string `json:"secret"`
//		} `json:"webhooks"`
//	} `json:"lark"`
//	// 多个钉钉webhook以逗号分割
//	DingDing struct {
//		Enable   bool `json:"enable"`
//		Webhooks []struct {
//			WebhookUrl string `json:"webhook_url"`
//			Secret     string `json:"secret"`
//		} `json:"webhooks"`
//	} `json:"dingding"`
//}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *Notify) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Notify{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (j Notify) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

package tbs

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"strings"
)


type JoborTask struct {
	BaseID
	Name          string   `gorm:"type:varchar(128);unique_index;comment:'任务名'" json:"name" form:"name"`
	Description   string   `gorm:"type:varchar(512);default:null;comment:'任务描述" json:"description" form:"description"`
	Lang          string   `gorm:"type:varchar(16);index:idx_code;not null;comment:'任务类型:[shell,api,python,golang,e.g.]'" json:"lang" form:"lang"`
	Data          TaskData `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	UserID        *uint    `gorm:"index:user_id;comment:'关联用户id'" json:"user_id"`
	Count         int      `gorm:"comment:'执行次数'" json:"count" form:"count"`
	Expr          string   `gorm:"type:varchar(32);not null;comment:'定时任务表达式：0/1 * * ? * * * 秒分时天月星期'" json:"expr" form:"expr"`
	Timeout       int      `gorm:"default:-1;comment:'超时时间'" json:"timeout" form:"timeout"`
	Status        string   `gorm:"type:varchar(32);default:'running';comment:'定时任务状态: running,stop'" json:"status" form:"status"`
	AlarmPolicy   int      `gorm:"default:2;comment:'告警策略：{0:never,1:always,2:failed,3:success}'" json:"alarm_policy" form:"alarm_policy"`
	ExpectContent string   `gorm:"type:varchar(500);default:null;comment:'期望任务返回结果'" json:"expect_content" form:"expect_content"`
	ExpectCode    int      `gorm:"default:0;comment:'期望任务状态码'" json:"expect_code" form:"expect_code"`
	Retry         int      `gorm:"default:0;comment:'重试次数'" json:"retry" form:"retry"`
	//Deleted  bool   `gorm:"default:false;comment:'逻辑删除'" json:"deleted" form:"deleted"`
	Prev JSONTime `gorm:"default: null;type:datetime;comment:'上次执行时间'" json:"prev" form:"prev"`
	Next JSONTime `gorm:"default: null;type:datetime;comment:'下次执行时间'" json:"next" form:"next"`
	BaseByUpdate
	BaseTime
}

type JoborLog struct {
	BaseID
	Name          string   `gorm:"type:varchar(128);index:idx_code;comment:'任务名'" json:"name" form:"name"`
	Lang          string   `gorm:"type:varchar(16);index:idx_code;not null;comment:'任务类型:[shell,api,python,golang,e.g.]'" json:"lang" form:"lang"`
	TaskID        *uint    `gorm:"index:task_id;comment:'关联任务id'" json:"task_id"`
	TriggerMethod string   `gorm:"type:varchar(16);comment:'任务触发方式：auto,manual'" json:"trigger_method"`
	Expr          string   `gorm:"type:varchar(32);not null;comment:'定时任务表达式：0/1 * * ? * * * 秒分时天月星期'" json:"expr" form:"expr"`
	//Data          string   `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	Data      TaskData `gorm:"type:text;not null;comment:'任务执行详细，格式：json'" json:"data" form:"data"`
	Resp      string   `gorm:"type:text;comment:'任务返回结果'" json:"resp"`
	CostTime  float32  `gorm:"comment:'任务耗时'" json:"cost_time" form:"cost_time"`
	Result    string   `gorm:"type:varchar(16);comment:'任务结果: success,failed'" json:"result" form:"result"`
	ErrCode   int      `gorm:"comment:'错误码'" json:"err_code" form:"err_code"`
	ErrMsg    string   `gorm:"type:text;comment:'错误信息'" json:"err_msg" form:"err_msg"`
	StartTime JSONTime `gorm:"default: null;type:datetime;comment:'开始时间'" json:"start_time" form:"start_time"`
	EndTime   JSONTime `gorm:"default: null;type:datetime;comment:'结束时间'" json:"end_time" form:"end_time"`
	BaseTime
}


type TaskData struct {
	Code string      `json:"data"`
	Api  struct{
		Url         string                   `json:"url"`
		Method      string                   `json:"method"`
		ContentType string                   `json:"content_type"`
		Payload     string                   `json:"payload"`
		Header      []map[string]interface{} `json:"header"`
		Forms       []map[string]interface{} `json:"forms"`

	} `json:"api"`
}


// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
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
	//bytes, ok := value.([]byte)
	//if !ok {
	//	return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	//}
	result := TaskData{}
	err := json.Unmarshal(bytes, &result)
	*j = result
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j TaskData) Value() (driver.Value, error) {
	if &j == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}


type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value: %s", value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// MarshalJSON to output non base64 encoded []byte
func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

// UnmarshalJSON to deserialize []byte
func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

func (j JSON) String() string {
	return string(j)
}

// GormDataType gorm common data type
func (JSON) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

// JSONQueryExpression json query expression, implements clause.Expression interface to use as querier
type JSONQueryExpression struct {
	column      string
	keys        []string
	hasKeys     bool
	equals      bool
	equalsValue interface{}
}

// JSONQuery query column as json
func JSONQuery(column string) *JSONQueryExpression {
	return &JSONQueryExpression{column: column}
}

// HasKey returns clause.Expression
func (jsonQuery *JSONQueryExpression) HasKey(keys ...string) *JSONQueryExpression {
	jsonQuery.keys = keys
	jsonQuery.hasKeys = true
	return jsonQuery
}

// Keys returns clause.Expression
func (jsonQuery *JSONQueryExpression) Equals(value interface{}, keys ...string) *JSONQueryExpression {
	jsonQuery.keys = keys
	jsonQuery.equals = true
	jsonQuery.equalsValue = value
	return jsonQuery
}

// Build implements clause.Expression
func (jsonQuery *JSONQueryExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql", "sqlite":
			switch {
			case jsonQuery.hasKeys:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString(fmt.Sprintf("JSON_EXTRACT(%s, '$.%s') IS NOT NULL", stmt.Quote(jsonQuery.column), strings.Join(jsonQuery.keys, ".")))
				}
			case jsonQuery.equals:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString(fmt.Sprintf("JSON_EXTRACT(%s, '$.%s') = ", stmt.Quote(jsonQuery.column), strings.Join(jsonQuery.keys, ".")))
					stmt.AddVar(builder, jsonQuery.equalsValue)
				}
			}
		case "postgres":
			switch {
			case jsonQuery.hasKeys:
				if len(jsonQuery.keys) > 0 {
					stmt.WriteQuoted(jsonQuery.column)
					stmt.WriteString("::jsonb")
					for _, key := range jsonQuery.keys[0 : len(jsonQuery.keys)-1] {
						stmt.WriteString(" -> ")
						stmt.AddVar(builder, key)
					}

					stmt.WriteString(" ? ")
					stmt.AddVar(builder, jsonQuery.keys[len(jsonQuery.keys)-1])
				}
			case jsonQuery.equals:
				if len(jsonQuery.keys) > 0 {
					builder.WriteString(fmt.Sprintf("json_extract_path_text(%v::json,", stmt.Quote(jsonQuery.column)))

					for idx, key := range jsonQuery.keys {
						if idx > 0 {
							builder.WriteByte(',')
						}
						stmt.AddVar(builder, key)
					}
					builder.WriteString(") = ")

					if _, ok := jsonQuery.equalsValue.(string); ok {
						stmt.AddVar(builder, jsonQuery.equalsValue)
					} else {
						stmt.AddVar(builder, fmt.Sprint(jsonQuery.equalsValue))
					}
				}
			}
		}
	}
}
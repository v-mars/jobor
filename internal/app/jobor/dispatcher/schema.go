package dispatcher

import (
	"jobor/internal/models"
	"jobor/internal/models/tbs"
)

var tbName = models.JoborTask

type ShowData struct {
	//ID        uint   `json:"id"`
	//Name      string `json:"Name" form:"Name"`
	//Code      string `json:"code" form:"code"`
	tbs.JoborTask
	//TaskData TaskData `json:"task_data"`
}
func (ShowData) TableName() string {
	return tbName
}


type PostSchema struct {
	ID            uint         `json:"id"`
	Name          string       `json:"name" binding:"required"`
	Description   string       `json:"description"`
	Lang          string       `json:"lang" binding:"required"`
	Data          tbs.TaskData `json:"data" binding:"required"`
	Notify        tbs.Notify   `json:"notify"`
	Expr          string       `json:"Expr" binding:"required"`
	Timeout       int          `json:"timeout"`
	AlarmPolicy   int          `json:"alarm_policy"`
	ExpectCode    int          `json:"expect_code"`
	ExpectContent string       `json:"expect_content"`
	//RoutePolicy   int          `json:"route_policy"`
	RoutingKey    string       `json:"routing_key"`
	Retry         int          `json:"retry"`
	Status        string       `json:"status"`
	ByUpdate      string       `json:"by_update,-"`
}
func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	Name          *string       `json:"name,omitempty"`
	Description   *string       `json:"description,omitempty"`
	Lang          *string       `json:"lang,omitempty"`
	Data          *tbs.TaskData `json:"data,omitempty" form:"data"`
	Notify        *tbs.Notify   `json:"notify,omitempty" form:"notify"`
	Expr          *string       `json:"expr,omitempty"`
	Timeout       *int          `json:"timeout,omitempty"`
	AlarmPolicy   *int          `json:"alarm_policy,omitempty"`
	RoutePolicy   *int          `json:"route_policy,omitempty"`
	RoutingKey    *string       `json:"routing_key,omitempty"`
	ExpectCode    *int          `json:"expect_code,omitempty"`
	ExpectContent *string       `json:"expect_content,omitempty"`
	Retry         *int          `json:"retry,omitempty"`
	Status        *string       `json:"status,omitempty"`
	ByUpdate      *string       `json:"by_update,-"`
}

func (PutSchema) TableName() string {
	return tbName
}

type DeleteSchema struct {
	Rows []uint `json:"rows"`
}

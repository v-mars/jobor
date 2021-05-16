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
	ID            uint         `json:"-"`
	Name          string       `json:"Name" binding:"required"`
	Description   string       `json:"description"`
	Lang          string       `json:"lang" binding:"required"`
	Data          tbs.TaskData `json:"data" binding:"required"`
	Notify        tbs.Notify   `json:"notify"`
	Expr          string       `json:"Expr" binding:"required"`
	Timeout       int          `json:"timeout"`
	ExpectCode    int          `json:"expect_code"`
	ExpectContent string       `json:"expect_content"`
	RoutingKey    string       `json:"routing_key"`
	Retry         int          `json:"retry"`
	Status        string       `json:"status"`
	ByUpdate      string       `json:"by_update,-"`
}
func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	Name          *string       `json:"Name"`
	Description   *string       `json:"description"`
	Lang          *string       `json:"lang"`
	Data          *tbs.TaskData `json:"data" form:"data"`
	Notify        *tbs.Notify   `json:"notify" form:"notify"`
	Expr          *string       `json:"Expr"`
	Timeout       *int          `json:"timeout"`
	RoutingKey    *string       `json:"routing_key"`
	ExpectCode    *int          `json:"expect_code"`
	ExpectContent *string       `json:"expect_content"`
	Retry         *int          `json:"retry"`
	Status        *string       `json:"status"`
	ByUpdate      *string       `json:"by_update,-"`
}

func (PutSchema) TableName() string {
	return tbName
}

type DeleteSchema struct {
	Rows []uint `json:"rows"`
}

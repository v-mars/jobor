package role

import "jobor/internal/models"


var tbName = models.Role

type ShowData struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
}
func (ShowData) TableName() string {
	return tbName
}

type PostSchema struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ByUpdate    string `json:"by_update,-"`
	Permissions []uint `json:"permissions,omitempty"`
}
func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	ID          uint   `json:"id" binding:"required"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ByUpdate    *string `json:"by_update,-"`
	Permissions *[]uint `json:"permissions,omitempty"`
}

func (PutSchema) TableName() string {
	return tbName
}

type DeleteSchema struct {
	Rows []uint `json:"rows"`
}
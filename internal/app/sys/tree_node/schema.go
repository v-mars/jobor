package tree_node

import "jobor/internal/models"

var tbName = models.TreeNode

type QuerySchema struct {
	Mark      string `form:"mark" binding:"required"`
	Name      string `form:"name"`
	//NodeID    uint   `form:"node_id"`
	//Condition string `json:"condition"`
}

type ShowMarkList struct {
	Mark      string `form:"mark" json:"mark"`
}

func (ShowMarkList) TableName() string {
	return tbName
}

//gt: greater than 大于
//gte: greater than or equal 大于等于
//lt: less than 小于
//lte: less than or equal 小于等于


type ShowSchema struct {
	ID       uint   `json:"id" form:"id"`
	ParentID uint   `json:"parent_id"`
	Lft      uint   `json:"lft" form:"lft"`
	Rgt      uint   `json:"rgt" form:"rgt"`
	Name     string `json:"name"`
	Mark     string `json:"mark" form:"mark"`
	//Children     []ShowSchema `json:"children"`
}
func (ShowSchema) TableName() string {
	return tbName
}

type PostSchema struct {
	ParentID  uint   `json:"parent_id"`
	Name      string `json:"name" binding:"required"`
	Mark      string `json:"mark" binding:"required"`
}
func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	SelfID    uint   `json:"self_id" binding:"required"`
	ParentID  uint   `json:"parent_id" `
	SiblingID int    `json:"sibling_id" ` // 移动至： 0:第一位置 -1：末尾位置
	Mark      string `json:"mark" binding:"required"`
}

type RenameSchema struct {
	Name      string `json:"name" binding:"required"`
}

type DeleteSchema struct {
	//SelfID    uint `json:"self_id" binding:"required"`
	Mark      string `json:"mark" binding:"required"`
}
package models

import (
	"gorm.io/gorm"
)

// 条件
type Option struct {
	DB         *gorm.DB
	Order      string
	Where      string
	Value      []interface{}
	Preloads   []string
	Select     string
	Joins      string
	Group      string
	Having     string
	Scan       bool
	First      bool
	All        bool
	Pluck      string
	Omits      []string
	NullError  bool
	Debug      bool
}
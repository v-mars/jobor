package user

import (
	"jobor/internal/models"
	"jobor/internal/models/tbs"
)

var tbName = models.User

type AllUser struct {
	ID        uint     `json:"id"`
	Nickname  string   `json:"nickname"`
	Username  string   `json:"username"`
}

func (AllUser) TableName() string {
	return tbName
}

type InfoUser struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Nickname  string   `json:"nickname"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
}

func (InfoUser) TableName() string {
	return tbName
}

type Role struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
}

type User struct {
	ID         uint         `json:"id"`
	Nickname   string       `json:"nickname"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	UserType   string       `json:"user_type"`
	Status     bool         `json:"status"`
	ByUpdate   string       `json:"by_update"`
	Roles      []Role       `gorm:"many2many:user_roles;" json:"roles"`
	Ctime      tbs.JSONTime `json:"ctime" form:"ctime"`
	Mtime      tbs.JSONTime `json:"mtime" form:"mtime"`
}

func (User) TableName() string {
	return tbName
}

type PostSchema struct {
	Nickname    string  `json:"nickname" binding:"required"`
	Username    string  `json:"username" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	Phone       string  `json:"phone"`
	Roles       []uint  `json:"roles" form:"roles"`
}

func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	//ID       uint   `json:"id" binding:"required"`
	Nickname *string `json:"nickname,omitempty"`
	Username *string `json:"username,omitempty"`
	//Password string `json:"password" binding:"required"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Status   *bool   `json:"status,omitempty"`
	Roles    *[]uint `json:"roles,omitempty"`
	ByUpdate string `json:"by_update,-"`
}

func (PutSchema) TableName() string {
	return tbName
}

type RestPass struct {
	ID       uint   `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePass struct {
	OldPassword string `json:"old_password" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"re_password" binding:"required"`
}

type ChangeProfile struct {
	Nickname *string `json:"nickname,omitempty"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	ByUpdate string `json:"by_update,-"`
}

func (ChangeProfile) TableName() string {
	return tbName
}
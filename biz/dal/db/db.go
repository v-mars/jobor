package db

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

//func NewModel(Db *gorm.DB) ModelInterface {
//	return Model{GormDB: Db}
//}

type ModelOld struct {
	GormDB *gorm.DB `gorm:"-" json:"-" form:"-" all:"-"`
	ID     int      `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	Ctime  JSONTime `gorm:"not null;type:datetime;default:current_timestamp;comment:创建时间" json:"ctime" form:"ctime"`
	Mtime  JSONTime `gorm:"<-:create;not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;comment:更新时间" json:"mtime"  form:"mtime"`
}

type ModelInterface interface {
	TableName() string
	//GetById(_id interface{}, out interface{}) error
	//GetByName(name string, out interface{}) error
	Add(newRow interface{}) error
	//Mod(newRow interface{}) error
	//Del(_id interface{}) error
}

type Model struct {
	GormDB    *gorm.DB `gorm:"-" json:"-" form:"-" all:"-"`
	ID        int      `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	CreatedAt JSONTime `gorm:"not null;type:datetime;default:current_timestamp;comment:创建时间" json:"created_at" form:"created_at"`
	UpdatedAt JSONTime `gorm:"not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"  form:"updated_at"`
}

func (*Model) TableName() string {
	return "tb_model"
}

type BaseName struct {
	Name string `gorm:"type:varchar(64);comment:更新人" json:"name" form:"name"`
}

type BaseComment struct {
	Comment string `gorm:"type:text;comment:备注" json:"comment" form:"comment"`
}

type Tor struct {
	Updater string `gorm:"type:varchar(64);comment:更新人" json:"updater" form:"updater"`
	Creator string `gorm:"type:varchar(64);comment:创建人" json:"creator" form:"creator"`
}

func (d *Model) GetById(_id interface{}, out interface{}) error {
	return d.GormDB.Table(d.TableName()).First(&out, _id).Error
}

func (d *Model) GetByName(name string, out interface{}) error {
	return d.GormDB.Table(d.TableName()).Where("name=?", name).First(&out).Error
}

func (d *Model) Add(newRow interface{}) error {
	//tx := Db.Begin()
	//defer func() {
	//	tx.Rollback()
	//}()
	fmt.Println("tbname:", d.TableName())
	return d.GormDB.Table(d.TableName()).Save(newRow).Error
}

func (d *Model) Mod(Db *gorm.DB, newRow interface{}) error {
	//tx := Db.Begin()
	//defer func() {
	//	tx.Rollback()
	//}()
	return Db.Table(d.TableName()).Save(&newRow).Error
}

func (d *Model) Del(Db *gorm.DB, _id interface{}) error {
	//tx := Db.Begin()
	//defer func() {
	//	tx.Rollback()
	//}()
	return Db.Table(d.TableName()).Delete(&d, _id).Error
}

//Ctime time.Time `gorm:"not null;type:datetime;default:current_timestamp" json:"ctime" form:"ctime"`
//Mtime time.Time `gorm:"not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;" json:"mtime"  form:"mtime"`

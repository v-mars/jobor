package tbs

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Base struct {
	ID        uint   `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                     // 主键
	Ctime     time.Time `gorm:"not null;type:datetime;default:current_timestamp" json:"ctime" form:"ctime"`
	Mtime     time.Time `gorm:"not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;" json:"mtime"  form:"mtime"`
	//CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
	//UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at" form:"updated_at"` // 更新时间
	//CreatedBy uint64    `gorm:"column:created_by;default:0;not null;" json:"created_by" form:"created_by"`     // 创建人
	//UpdatedBy uint64    `gorm:"column:updated_by;default:0;not null;" json:"updated_by" form:"updated_by"`     // 更新人
}


type ByUpdateDeletedTime struct {
	BaseByUpdate
	BaseDeleted
	BaseTime
}

type ByUpdateDeleted struct {
	BaseByUpdate
	BaseDeleted
}

type ByUpdateTime struct {
	BaseByUpdate
	BaseTime
}

type BaseID struct {
	ID  uint  `gorm:"column:id;primary_key;comment:'主键id'" json:"id" form:"id"`                     // 主键
}

type BaseName struct {
	Name string `gorm:"type:varchar(64);comment:'更新人'" json:"name" form:"name"`
}

type BaseDescription struct {
	ByUpdate string `gorm:"type:varchar(64);comment:'更新人'" json:"by_update" form:"by_update"`
}

type BaseByUpdate struct {
	ByUpdate string `gorm:"type:varchar(64);comment:'更新人'" json:"by_update" form:"by_update"`
}

type BaseDeleted struct {
	Deleted  bool   `gorm:"default:false;comment:'标记是否删除';" json:"deleted" form:"deleted"`
}

type BaseTime struct {
	Ctime JSONTime `gorm:"not null;type:datetime;default:current_timestamp;comment:'创建时间'" json:"ctime" form:"ctime"`
	Mtime JSONTime `gorm:"<-:create;not null;type:datetime;default:current_timestamp on update CURRENT_TIMESTAMP;comment:'更新时间'" json:"mtime"  form:"mtime"`
}

func GetTablePrefix() string {
	return "tb_"
}



// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
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
package audit

import (
	"gorm.io/gorm"
	"jobor/biz/dal/db"
)

const (
	NameAuditLog = "audit"
)

type AuditLog struct {
	//Id        int    `gorm:"column:id;primary_key;comment:主键id" json:"id" form:"id"` // 主键
	UserId   int64         `gorm:"comment:关联用户" json:"user_id" form:"user_id"`
	Nickname string        `gorm:"type:varchar(128);comment:显示名" json:"nickname" form:"nickname"`
	User     string        `gorm:"type:varchar(128);comment:用户名" json:"user" form:"user"`
	Action   string        `gorm:"type:varchar(256);comment:具体方法描述" json:"action" form:"action"`
	Business string        `gorm:"type:varchar(128);comment:业务" json:"business" form:"business"`
	Method   string        `gorm:"type:varchar(8);comment:请求方法" json:"method" form:"method"`
	Handler  string        `gorm:"type:varchar(256);comment:处理程序/接口" json:"handler" form:"handler"`
	UserIp   string        `gorm:"type:varchar(128);comment:用户源IP" json:"user_ip" form:"user_ip"`
	HttpCode int           `gorm:"comment:返回状态" json:"http_code" form:"http_code"`
	Body     string        `gorm:"type:longtext;comment:请求Body" json:"body" form:"body"`
	RespBody string        `gorm:"type:longtext;comment:响应Body" json:"resp_body" form:"resp_body"`
	Dom      string        `gorm:"type:varchar(128);default:default;comment:所属域" json:"dom" form:"dom"`
	ObjId    int64         `gorm:"comment:关联变更对象ID" json:"obj_id" form:"obj_id"`
	CostTime db.MillisTime `gorm:"comment:耗时" json:"cost_time" form:"cost_time"`
	Status   int           `gorm:"comment:状态[0: error, 1: success, 2: forbidden, 3: exception]" json:"status" form:"status"`
	db.Model
}

func (al *AuditLog) TableName() string {
	return NameAuditLog
}

func (al *AuditLog) CreateAuditLog() error {
	return db.DB.Table(al.TableName()).Create(al).Error
}

func SelectScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		//sql := `user.id,nickname,username,email,phone,userid,empno,user_type,avatar,otp_secret,otp_enabled,updater,path,status,user.created_at,user.updated_at`
		return Db
	}
}
func SelectAllScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,username,status")
	}
}

func WhereScopes(req *AuditReq) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "user like ? and nickname like ? and action like ? and user_ip like ?"
		var sqlArgs = []interface{}{"%" + req.User + "%", "%" + req.User + "%", "%" + req.Action + "%", "%" + req.UserIp + "%"}
		if len(req.Method) > 0 {
			sql = sql + " and method=?"
			sqlArgs = append(sqlArgs, req.Method)
		}
		if req.HttpCode != nil {
			sql = sql + " and http_code=?"
			sqlArgs = append(sqlArgs, req.HttpCode)
		}
		if req.UserId != nil {
			sql = sql + " and user_id=?"
			sqlArgs = append(sqlArgs, req.UserId)
		}
		if req.ObjId != nil {
			sql = sql + " and obj_id=?"
			sqlArgs = append(sqlArgs, req.ObjId)
		}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Joins("")
	}
}
func PreloadScopes(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderScopes() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("audit_log.id desc")
	}
}

func GroupScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("audit_log.id")
	}
}

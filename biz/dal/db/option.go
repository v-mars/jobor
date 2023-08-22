package db

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

const maxPageSize = 1000

// Option 条件
type Option struct {
	DB        *gorm.DB
	Order     interface{}
	Where     interface{}
	Value     []interface{}
	Preloads  []string
	Select    interface{}
	Joins     string
	Group     string
	Having    interface{}
	Distinct  []interface{}
	Scan      bool
	First     bool
	All       bool
	Pluck     string
	Omits     []string
	NullError bool
	Debug     bool
	Table     string
}

func (o Option) Init() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(o.Table).Scopes(o.SelectO(), o.WhereO(), o.JoinsO(),
			o.Preload(), o.DistinctO(), o.OrderO(), o.GroupO(), o.HavingO())
	}
}
func (o Option) SelectO(args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Select != nil {
			return db.Select(o.Select, args)
		}
		return db
	}
}

func (o Option) WhereO(args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Where != nil {
			return db.Where(o.Where, args)
		}
		return db
	}
}

func (o Option) JoinsO(args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(o.Joins) > 0 {
			return db.Joins(o.Joins, args)
		}
		return db
	}
}

func (o Option) GroupO() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(o.Group) > 0 {
			return db.Group(o.Group)
		}
		return db
	}
}

func (o Option) OrderO() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Order != nil {
			return db.Order(o.Order)
		}
		return db
	}
}

func (o Option) HavingO(args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if o.Having != nil {
			return db.Having(o.Having, args)
		}
		return db
	}
}

func (o Option) Preload(args ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(o.Preloads) > 0 {
			for _, v := range o.Preloads {
				db = db.Preload(v)
			}
		}
		return db
	}
}

func (o Option) DistinctO() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(o.Distinct) > 0 {
			return db.Distinct(o.Distinct)
		}
		return db
	}
}

func (o Option) Paginate(c *app.RequestContext) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		qp, _ := c.GetQuery("page")
		qps, _ := c.GetQuery("pageSize")
		page, _ := strconv.Atoi(qp)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(qps)
		switch {
		case pageSize > maxPageSize:
			pageSize = maxPageSize
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (o Option) GetWithScopes(db *gorm.DB, table string, funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	for _, f := range funcs {
		if f != nil {
			scopes = append(scopes, f)
		}
	}
	return db.Table(table).Scopes(scopes...)
}

func (o Option) GetCountWithScopes(db *gorm.DB, table string, funcs ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var i int64
	var scopes []func(*gorm.DB) *gorm.DB
	for _, f := range funcs {
		if f != nil {
			scopes = append(scopes, f)
		}
	}
	err := db.Table(table).Scopes(scopes...).Count(&i).Error
	if err != nil {
		return i, err
	}
	return i, nil
}

func (o Option) PageDataWithScopes(db *gorm.DB, table string, data interface{}, pageFunc func(*gorm.DB) *gorm.DB,
	funcs ...func(*gorm.DB) *gorm.DB) (count int64, page int64, pageSize int64, err error) {
	o.DB = o.Init()(db)
	if err = db.Table(table).Scopes(funcs...).Count(&count).Error; err != nil {
		hlog.Errorf("get db[%s] data count err: %s", table, err)
		return count, page, pageSize, err
	}

	if err = db.Table(table).Scopes(pageFunc).Find(data).Error; err != nil {
		hlog.Errorf("get db[%s] page data data err: %s", table, err)
		return count, page, pageSize, err
	}
	return count, page, pageSize, nil
}

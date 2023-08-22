package model

import (
	"errors"
	"fmt"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	"regexp"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

func PluckList(dbObj *gorm.DB, model, where interface{}, out interface{}, fieldName string) error {
	return dbObj.Model(model).Where(where).Pluck(fieldName, out).Error
}

func Error(err error, nullError bool) bool {
	notFound := errors.Is(err, gorm.ErrRecordNotFound)
	if notFound && nullError == false {
		return false
	} else {
		return true
	}
}

func Init(model interface{}, o db.Option, dbObj *gorm.DB) *gorm.DB {
	if o.DB != nil {
		dbObj = o.DB
	} else if dbObj == nil {
		dbObj = db.DB
	}
	if dbObj == nil {
		panic(any("db client is nil!"))
	}
	if len(o.Table) > 0 {
		dbObj = dbObj.Table(o.Table)
	} else if model != nil {
		dbObj = dbObj.Model(model)
	}
	if o.Select != nil {
		dbObj = dbObj.Select(o.Select) // "distinct id, name, username"
	}
	if len(o.Preloads) > 0 {
		for _, v := range o.Preloads {
			dbObj = dbObj.Preload(v)
		}
	}
	if len(o.Joins) > 0 {
		dbObj = dbObj.Joins(o.Joins) // "distinct id, name, username"
	}
	if o.Where != nil {
		dbObj = dbObj.Where(o.Where, o.Value...)
	}
	if o.Order != "" {
		dbObj = dbObj.Order(o.Order)
	}
	if len(o.Group) > 0 {
		dbObj = dbObj.Group(o.Group)
	}
	if o.Having != nil {
		dbObj = dbObj.Having(o.Having)
	}
	if o.Debug {
		dbObj = dbObj.Debug()
	}
	return dbObj
}

func Exist(model interface{}, o db.Option, db *gorm.DB) (bool, error) {
	var out int64
	dbObj := Init(model, o, db)
	err := dbObj.Count(&out).Error
	if err != nil && Error(err, false) {
		return false, err
	}
	if out > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func Raw(Model interface{}, sql string, o db.Option, out *[]string, db *gorm.DB) error {
	if o.Scan {
		if err := db.Model(Model).Raw(sql, o.Value...).Scan(out).Error; err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	} else if o.Pluck != "" {
		if err := db.Model(Model).Raw(sql, o.Value...).Pluck(o.Pluck, out).Error; err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	} else {
		if err := db.Model(Model).Raw(sql, o.Value...).Find(o.Select, out).Error; err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	}
}

func GetResult(dbObj *gorm.DB, o db.Option, out interface{}) error {
	if o.Scan {
		//fmt.Println("sql query:", dbObj.Scan(pageData.List).QueryExpr())
		err := dbObj.Scan(out).Error
		if err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	} else if o.First == true {
		err := dbObj.First(out).Error
		if err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	} else if o.Pluck != "" {
		err := dbObj.Pluck(o.Pluck, out).Error
		if err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	} else {
		err := dbObj.Find(out).Error
		if err != nil && Error(err, o.NullError) {
			return err
		} else {
			return nil
		}
	}
}

type BaseSql struct {
	Model interface{}
}

type Association struct {
	Column string
	Values interface{}
}

func Query(db *gorm.DB, Model interface{}, o db.Option, pageData *response.PageDataList) error {
	dbObj := Init(Model, o, db)
	var err error
	if len(o.Distinct) > 0 {
		countObj := Init(Model, o, db)
		err = countObj.Distinct(o.Distinct).Count(&pageData.Total).Error
	} else {
		err = dbObj.Count(&pageData.Total).Error
	}
	if err != nil && Error(err, o.NullError) {
		return err
	}
	if pageData.PageSize != 0 && o.All == false {
		dbObj = dbObj.Limit(pageData.PageSize).Offset((pageData.Page - 1) * pageData.PageSize)
	} else {
		pageData.PageSize = int(pageData.Total)
		if int(pageData.Total) <= 5000 {
			pageData.PageSize = int(pageData.Total)
		} else {
			pageData.PageSize = 5000
		}
		if pageData.Page == 0 {
			pageData.Page = 1
		}
		dbObj = dbObj.Limit(pageData.PageSize).Offset((pageData.Page - 1) * pageData.PageSize)
	}
	err = GetResult(dbObj, o, pageData.List)
	return err
}

func Get(db *gorm.DB, Model interface{}, o db.Option, out interface{}) error {
	dbObj := Init(Model, o, db)
	err := GetResult(dbObj, o, out)
	return err
}

func Create(tx *gorm.DB, Model interface{}, autoCommit bool) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(Model).Error; err != nil {
		tx.Rollback()
		// Error 1062: Duplicate entry '1.1.1.2' for key 'idx_name_code'
		comp := regexp.MustCompile(`Error 1062: Duplicate entry '(.+)'? for key '(.+)'?`)
		subMatches := comp.FindAllStringSubmatch(err.Error(), -1)
		// "记录重复。具体报错："
		if len(subMatches) > 0 {
			err = errors.New(fmt.Sprintf("记录重复。具体报错：%s", subMatches[0]))
			//fmt.Println("subMatches:", subMatches)
			//if len(subMatches[0]) >= 2 {
			//	err = errors.New(fmt.Sprintf("记录重复。具体报错：%s", subMatches[0][1]))
			//} else {
			//	err = errors.New(fmt.Sprintf("记录重复。具体报错：%s", subMatches[0]))
			//}
		}
		return err
	}
	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	}
	return nil
}

func Update(tx *gorm.DB, Model interface{}, o db.Option, data map[string]interface{}, ass []Association, autoCommit bool,
) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Model(Model).Where(o.Where, o.Value...).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(ass) > 0 {
		for _, v := range ass {
			//fmt.Println(v)
			if err := tx.Model(Model).Association(v.Column).Replace(v.Values); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func Delete(tx *gorm.DB, Model interface{}, o db.Option, Associations []string, autoCommit bool) error {
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Model(Model).Where(o.Where, o.Value...).First(Model).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, v := range Associations {
		if err := tx.Model(Model).Where(o.Where, o.Value...).Association(v).Clear(); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Delete(Model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func GetById(DB *gorm.DB, Model interface{}, sqlId interface{}, out interface{}, notFoundError bool) error {
	if err := DB.Model(Model).First(out, "id=?", sqlId).Error; err != nil && Error(err, notFoundError) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		return err
	}
	return nil
}

func UpdateById(tx *gorm.DB, Model interface{}, sqlId interface{},
	data map[string]interface{}, ass []Association, autoCommit bool) error {
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.First(Model, "id = ?", sqlId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		tx.Rollback()
		return err
	}

	var OmitList []string
	for _, v := range ass {
		OmitList = append(OmitList, strings.ToLower(v.Column))
	}
	if err := tx.Model(Model).Omit(strings.Join(OmitList, ",")).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if ass != nil && len(ass) > 0 {
		for _, v := range ass {
			if err := tx.Model(Model).Association(v.Column).Replace(v.Values); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func UpdateByIdStruct(tx *gorm.DB, Model interface{}, sqlId interface{},
	data interface{}, ass []Association, autoCommit bool) error {
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.First(Model, "id = ?", sqlId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		tx.Rollback()
		return err
	}

	var OmitList []string
	for _, v := range ass {
		OmitList = append(OmitList, strings.ToLower(v.Column))
	}

	if err := tx.Model(Model).Omit(strings.Join(OmitList, ",")).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ass != nil && len(ass) > 0 {
		for _, v := range ass {
			if err := tx.Model(Model).Association(v.Column).Replace(v.Values); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func DeleteById(tx *gorm.DB, Model interface{}, sqlId interface{}, Associations []string, autoCommit bool) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != any(nil) {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//fmt.Println("model:", d.Model)
	if err := tx.Model(Model).Where("id = ?", sqlId).First(Model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		tx.Rollback()
		return err
	}

	for _, v := range Associations {
		if err := tx.Model(Model).Association(v).Clear(); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Delete(Model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if autoCommit {
		err := tx.Commit().Error
		if err != nil {
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

// ************************ WithScopes **************************

func Paginate(c *app.RequestContext) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		qp, _ := c.GetQuery("page")
		qps, _ := c.GetQuery("pageSize")
		page, _ := strconv.Atoi(qp)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(qps)
		switch {
		case pageSize > 1000:
			pageSize = 1000
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetWithScopes(db *gorm.DB, table string, funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	for _, f := range funcs {
		if f != nil {
			scopes = append(scopes, f)
		}
	}
	return db.Table(table).Scopes(scopes...)
}

func GetCountWithScopes(db *gorm.DB, table string, funcs ...func(*gorm.DB) *gorm.DB) (int64, error) {
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

const (
	Find = "find"
	Scan = "scan"
)

func PageDataWithScopes(DB *gorm.DB, table string, Type string,
	pd *response.PageDataList, noCountScopes []func(*gorm.DB) *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) (err error) {
	if err = DB.Table(table).Scopes(funcs...).Count(&pd.Total).Error; err != nil {
		hlog.Errorf("get db[%s] data count err: %s", table, err)
		return err
	}
	var pageFunc = func(DB *gorm.DB) func(db *gorm.DB) *gorm.DB {
		return func(DB *gorm.DB) *gorm.DB {
			return DB
		}
	}
	if pd.PageSize != 0 {
		pageFunc = func(DB *gorm.DB) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				offset := (pd.Page - 1) * pd.PageSize
				return db.Offset(offset).Limit(pd.PageSize)
			}
		}
	}
	switch Type {
	case Scan:
		if err = DB.Table(table).Scopes(noCountScopes...).Scopes(pageFunc(DB)).Scan(pd.List).Error; err != nil {
			hlog.Errorf("get db[%s] page data err: %s", table, err)
			return err
		}
	default:
		if err = DB.Table(table).Scopes(noCountScopes...).Scopes(pageFunc(DB)).Find(pd.List).Error; err != nil {
			hlog.Errorf("get db[%s] page data err: %s", table, err)
			return err
		}
	}
	return nil
}

func DataWithScopes(DB *gorm.DB, table string, Type string,
	list interface{}, noCountScopes []func(*gorm.DB) *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) (count int64, err error) {
	if err = DB.Table(table).Scopes(funcs...).Count(&count).Error; err != nil {
		hlog.Errorf("get db[%s] data count err: %s", table, err)
		return count, err
	}

	switch Type {
	case Scan:
		if err = DB.Table(table).Scopes(noCountScopes...).Scopes(funcs...).Scan(list).Error; err != nil {
			hlog.Errorf("get db[%s] data err: %s", table, err)
			return count, err
		}
	default:
		if err = DB.Table(table).Scopes(noCountScopes...).Scopes(funcs...).Find(list).Error; err != nil {
			hlog.Errorf("get db[%s] data err: %s", table, err)
			return count, err
		}
	}
	return count, nil
}

func GetByIdScopes(DB *gorm.DB, table string, sqlId interface{}, out interface{}, notFoundError bool,
	funcs ...func(*gorm.DB) *gorm.DB) error {
	if err := DB.Table(table).Scopes(funcs...).First(out, "id=?", sqlId).Error; err != nil && Error(err, notFoundError) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录id[%v]不存在", sqlId))
		}
		return err
	}
	return nil
}

func PreloadScopes(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}
func JoinsScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Joins(`left join org on org.id=sso_client.org_id`)
	}
}

func GetScopesList(funcs ...func(*gorm.DB) *gorm.DB) []func(*gorm.DB) *gorm.DB {
	return funcs
}

func GetByIdOrNameScopes(DB *gorm.DB, table string, sqlObj interface{}, out interface{}, notFoundError bool,
	funcs ...func(*gorm.DB) *gorm.DB) error {
	if err := DB.Table(table).Scopes(funcs...).First(out, "id=? or name=?", sqlObj, sqlObj).Error; err != nil && Error(err, notFoundError) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("该记录[%v]不存在", sqlObj))
		}
		return err
	}
	return nil
}

func GroupScopes(name string) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group(name)
	}
}

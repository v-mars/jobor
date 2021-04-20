package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"jobor/internal/app/response"
	"strings"

	"jobor/internal/app/models/db"
	"regexp"
)


// Save
func Save(value interface{})error{
	return db.DB.Save(value).Error
}

// Updates
func Updates(where interface{},value interface{})error{
	return db.DB.Model(where).Updates(value).Error
}

// Delete
func DeleteByModel(model interface{}) (count int64 ,err error){
	dbObj:= db.DB.Delete(model)
	err=dbObj.Error
	if err!=nil {
		return
	}
	count=dbObj.RowsAffected
	return
}

// Delete
func DeleteByWhere(model,where interface{}) (count int64 ,err error){
	dbObj:= db.DB.Where(where).Delete(model)
	err=dbObj.Error
	if err!=nil {
		return
	}
	count=dbObj.RowsAffected
	return
}
// Delete
func DeleteByID(model interface{},id uint64) (count int64 ,err error){
	dbObj:= db.DB.Where("id=?", id).Delete(model)
	err=dbObj.Error
	if err!=nil {
		return
	}
	count=dbObj.RowsAffected
	return
}


// GetPage

// PluckList
func PluckList(model,where interface{},out interface{},fieldName string)error{
	return db.DB.Model(model).Where(where).Pluck(fieldName, out).Error
}

func Error(err error, nullError bool) bool {
	notFound:=errors.Is(err, gorm.ErrRecordNotFound)
	if notFound && nullError == false {
		return false
	}  else {
		return true
	}
}

func DBInt(model interface{}, o Option, dbObj *gorm.DB) *gorm.DB {
	if o.DB!=nil {
		dbObj= o.DB
	} else if dbObj == nil {
		dbObj = db.DB
	}
	if dbObj == nil{
		panic("db client is nil!")
	}
	if model !=nil{
		dbObj = dbObj.Model(model)
	}
	if len(o.Select)>0{
		dbObj=dbObj.Select(o.Select)		// "distinct id, name, username"
	}
	if len(o.Preloads)>0 {
		for _,v:=range o.Preloads{
			dbObj=dbObj.Preload(v)
		}
	}
	if len(o.Joins)>0{
		dbObj=dbObj.Joins(o.Joins)		// "distinct id, name, username"
	}
	if len(o.Where)>0 {
		dbObj=dbObj.Where(o.Where,o.Value...)
	}
	if o.Order !="" {
		dbObj=dbObj.Order(o.Order)
	}
	if len(o.Group)>0{
		dbObj=dbObj.Group(o.Group)
	}
	if len(o.Having)>0{
		dbObj=dbObj.Joins(o.Having)
	}
	if o.Debug{dbObj=dbObj.Debug()}
	return dbObj
}

func Exist(model interface{}, o Option)(bool,error) {
	var out int64
	dbObj:= DBInt(model, o, db.DB)
	err := dbObj.Count(&out).Error
	if err!= nil && Error(err,false) {
		return false,err
	}
	if out > 0 {
		return true,nil
	}else {
		return false, nil
	}
}

func Raw(Model interface{},sql string, o Option, out *[]string) error {
	if o.Scan{
		if err:= db.DB.Model(Model).Raw(sql, o.Value...).Scan(out).Error; err!=nil && Error(err, o.NullError){
			return err
		} else {
			return nil
		}
	} else if o.Pluck != "" {
		if err:= db.DB.Model(Model).Raw(sql, o.Value...).Pluck(o.Pluck, out).Error; err!=nil && Error(err, o.NullError){
			return err
		} else {
			return nil
		}
	}else {
		if err:= db.DB.Model(Model).Raw(sql, o.Value...).Find(o.Select, out).Error; err!=nil && Error(err, o.NullError){
			return err
		} else {
			return nil
		}
	}
}

func GetResult(dbObj *gorm.DB, o Option, out interface{}) error {
	if o.Scan {
		//fmt.Println("sql query:", dbObj.Scan(pageData.List).QueryExpr())
		err:=dbObj.Scan(out).Error
		if err!= nil && Error(err,o.NullError) {
			return err
		} else {
			return nil
		}
	}else if o.First == true {
		err := dbObj.First(out).Error
		if err!= nil && Error(err,o.NullError) {
			return err
		} else {
			return nil
		}
	}else if o.Pluck != "" {
		err := dbObj.Pluck(o.Pluck, out).Error
		if err!= nil && Error(err,o.NullError) {
			return err
		} else {
			return nil
		}
	} else {

		err:=dbObj.Find(out).Error
		if err!= nil && Error(err, o.NullError){
			return err
		}else {
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

// Query
func Query(db *gorm.DB,Model interface{},o Option, pageData *response.PageDataList) error {
	dbObj:= DBInt(Model, o, db)
	err:=dbObj.Count(&pageData.Total).Error
	if err!=nil && Error(err,o.NullError){
		return err
	}
	if pageData.PageSize != 0 && o.All == false {
		dbObj = dbObj.Limit(pageData.PageSize).Offset((pageData.PageNumber - 1) * pageData.PageSize)
	} else {
		pageData.PageSize = int(pageData.Total)
		pageData.PageNumber = 1
	}
	err= GetResult(dbObj, o, pageData.List)
	return err
}

func Get(db *gorm.DB, Model interface{}, o Option, out interface{}) error {
	dbObj:= DBInt(Model, o, db)
	err:= GetResult(dbObj, o, out)
	return err
}

func Create(tx *gorm.DB,Model interface{}, autoCommit bool) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
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
		if len(subMatches) > 0{
			if len(subMatches[0]) >= 2{
				err = errors.New(fmt.Sprintf("记录重复。具体报错：%s", subMatches[0][1]))
			}else {
				err = errors.New(fmt.Sprintf("记录重复。具体报错：%s", subMatches[0]))
			}
		}
		return err
	}
	if autoCommit{
		err:=tx.Commit().Error
		if  err!=nil{
			tx.Rollback()
		}
		return err
	}
	return nil
}

func Update(tx *gorm.DB,Model interface{},o Option, data map[string]interface{}, ass []Association,autoCommit bool,
) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
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
	if len(ass)>0 {
		for _,v:=range ass{
			//fmt.Println(v)
			if err:= tx.Model(Model).Association(v.Column).Replace(v.Values);err!=nil{
				tx.Rollback()
				return err
			}
		}
	}

	if autoCommit{
		err:=tx.Commit().Error
		if  err!=nil{
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func Delete(tx *gorm.DB,Model interface{},o Option,Associations []string, autoCommit bool) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err:=tx.Model(Model).Where(o.Where, o.Value...).First(Model).Error;err!=nil{
		tx.Rollback()
		return err
	}

	for _, v:=range Associations{
		if err:= tx.Model(Model).Where(o.Where, o.Value...).Association(v).Clear();err!=nil{
			tx.Rollback()
			return err
		}
	}
	if err := tx.Delete(Model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if autoCommit{
		err:=tx.Commit().Error
		if  err!=nil{
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}



func GetById(DB *gorm.DB,Model interface{}, sqlId interface{},out interface{}, notFoundError bool) error {
	if err:=DB.Model(Model).First(out, "id=?", sqlId).Error;err!=nil && Error(err, notFoundError){
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		return err
	}
	return nil
}

func UpdateById(tx *gorm.DB,Model interface{}, sqlId interface{},
data map[string]interface{}, ass []Association,autoCommit bool) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.First(Model,"id = ?", sqlId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		tx.Rollback()
		return err
	}

	var OmitList []string
	for _,v:=range ass{
		OmitList = append(OmitList, strings.ToLower(v.Column))
	}
	if err := tx.Model(Model).Omit(strings.Join(OmitList, ",")).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	if ass!=nil && len(ass)>0 {
		for _,v:=range ass{
			if err:= tx.Model(Model).Association(v.Column).Replace(v.Values);err!=nil{
				tx.Rollback()
				return err
			}
		}
	}
	if autoCommit{
		err:=tx.Commit().Error
		if err!=nil{
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

func DeleteById(tx *gorm.DB,Model interface{},sqlId interface{},Associations []string, autoCommit bool) error {
	//tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	//fmt.Println("model:", d.Model)
	if err:=tx.Model(Model).Where("id = ?", sqlId).First(Model).Error;err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			err=errors.New(fmt.Sprintf("该记录id[%s]不存在", sqlId))
		}
		tx.Rollback()
		return err
	}

	for _, v:=range Associations{
		if err:= tx.Model(Model).Association(v).Clear();err!=nil{
			tx.Rollback()
			return err
		}
	}
	if err := tx.Delete(Model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if autoCommit{
		err:=tx.Commit().Error
		if  err!=nil{
			tx.Rollback()
		}
		return err
	} else {
		return nil
	}
}

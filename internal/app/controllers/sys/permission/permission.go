package permission

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jobor/internal/app/controllers"
	"jobor/internal/app/models"
	"jobor/internal/app/models/tbs"
	"jobor/internal/app/response"
	"jobor/internal/app/utils/casbin"
	"jobor/pkg/convert"
	"strings"
)

type IPermission interface {
	controllers.CommonInterfaces
	UpdatePermission(conditions []tbs.Permission) error
}
var dom = "sys"
type Permission struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) IPermission {
	return Permission{DB: DB}
}

// @Tags 权限管理
// @Summary 权限详细
// @Description Permission
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/permission/{id} [get]
func (r Permission) Get(c *gin.Context) {
	_id := c.Params.ByName("id")
	var obj tbs.Permission
	o := r.Option()
	o.Where = "id = ?"
	o.Value = append(o.Value, _id)
	o.First = true
	o.NullError = true
	err := models.Get(r.DB,&obj, o, &obj)
	if err!= nil {
		response.Error(c, err)
		return
	} else {
		response.Success(c,&obj)
		return
	}
}

// @Tags 权限管理
// @Summary 权限列表
// @Description Permission
// @Produce  json
// @Security ApiKeyAuth
// @Param name query string false "权限名"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/permission [get]
func (r Permission) Query(c *gin.Context) {
	var obj []tbs.Permission
	var pageData = response.InitPageData(c, &obj, false)
	type Param struct {
		Name   string `form:"name"`   // `form:"name" binding:"required"`
		Method string `form:"method"` // `form:"name" binding:"required"`
		Path   string `form:"path"`    // `form:"name" binding:"required"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	var o = r.Option()
	o.Where = "name like ? and method like ? and path like ?"
	o.Value = append(o.Value, "%"+param.Name+"%", "%"+param.Method+"%", "%"+param.Path+"%")
	o.Order = "ID DESC"
	o.Scan = true
	tx :=r.DB.Begin()
	defer func() {tx.Rollback()}()
	err := models.Query(tx,&tbs.Permission{}, o, &pageData)
	if err != nil {
		response.Error(c, err)
		return
	} else {
		response.PageSuccess(c, pageData)
	}
}

// @Tags 权限管理
// @Summary 创建权限
// @Description Permission
// @Produce  json
// @Security ApiKeyAuth
// @Param payload body PostSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/permission [post]
func (r Permission) Create(c *gin.Context) {
	//var err error
	//u, err:= user.GetUserValue(c)
	//if err!=nil{
	//	response.Error(c, err)
	//	return
	//}
	var obj PostSchema
	if err := c.ShouldBindQuery(&obj); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	//obj.ByUpdate = u.Nickname
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()

	if err:= models.Create(tx, &obj,true);err!=nil{
		response.Error(c, err)
		return
	}
	response.CreateSuccess(c, obj)
}

// @Tags 权限管理
// @Summary 更新权限
// @Description Permission
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Param payload body PutSchema true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/permission/{id} [put]
func (r Permission) Update(c *gin.Context) {
	var err error
	//u, err:= user.GetUserValue(c)
	//if err!=nil{
	//	response.Error(c, err)
	//	return
	//}
	var obj PutSchema
	if err := c.ShouldBindJSON(&obj); err!=nil{
		response.Error(c, err)
		return
	}
	//obj.ByUpdate = &u.Nickname
	_id := c.Params.ByName("id")
	var MapData map[string]interface{}
	if MapData,err=convert.StructToMap(obj); err!=nil{
		response.Error(c, err)
		return
	}
	//fmt.Println("mapData", MapData)
	tx :=r.DB.Begin()
	defer func() {tx.Rollback()}()
	var res tbs.Permission
	if err:= models.UpdateById(tx, &res, _id,MapData,nil, true);err!=nil{
		response.Error(c, err)
		return
	}
	response.UpdateSuccess(c, res)
}

// @Tags 权限管理
// @Summary 删除权限
// @Description Permission
// @Produce  json
// @Security ApiKeyAuth
// //@Param id path int true "ID"
// @Param payload body DeleteSchema true "参数信息: {rows:[1,2]}"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/sys/permission [delete]
func (r Permission) Delete(c *gin.Context) {
	//_id := c.Params.ByName("id")
	var data map[string][]int
	if err:= c.ShouldBindJSON(&data);err!=nil{
		response.Error(c, err)
		return
	}
	//rows["rows"]
	tx :=r.DB.Begin()
	defer func() {
		tx.Rollback()
	}()
	for _,_id := range data["rows"]{
		if err:= models.DeleteById(tx, &tbs.Permission{}, _id, []string{}, true); err!=nil{
			response.Error(c, err)
			return
		}
	}

	response.DeleteSuccess(c)
}


func (r Permission) Option() models.Option {
	var o models.Option
	o.Select = "distinct id,name,path,method,ctime,mtime"
	return o
}

func (r Permission) UpdatePermission(conditions []tbs.Permission) error {
	tx :=r.DB.Begin()
	defer func() {tx.Rollback()}()
	var obj []tbs.Permission
	var pageData = response.PageDataList{PageNumber: 1,PageSize:0,List:&obj}
	var o = r.Option()
	o.All = true
	var cdsStrList []string
	for _,v:=range conditions{
		cdsStrList = append(cdsStrList, fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method))
	}
	o.Where = strings.Join(cdsStrList, " or ")
	err := models.Query(tx, &tbs.Permission{},o, &pageData)
	if err!=nil{
		return err
	}

	// 过滤新增的列表
	var newRows []tbs.Permission
	var existList []string
	var existIdList []uint
	for _,v:=range obj{
		existIdList = append(existIdList, v.ID)
		existList = append(existList, fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method))
	}

	for _,v:=range conditions{
		var str1 = fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method)
		if !CheckExist(str1, existList){
			newRows = append(newRows, v)
		}
	}

	// 过滤出需要删除的列表
	var deleteList []tbs.Permission
	var deleteIdList []uint
	if err:=tx.Find(&deleteList, "id not in (?)", existIdList).Error;err!=nil{return err}

	for _,v:=range deleteList{
		deleteIdList = append(deleteIdList, v.ID)
		if _, err := casbin.Enforcer.RemoveFilteredNamedPolicy("p", 1, dom, v.Path, v.Method);err!=nil{return err}
	}
	if err:=tx.Exec("DELETE FROM role_permissions WHERE permission_id in (?)", deleteIdList).Error;err!=nil{return err}
	if err:=tx.Delete(&tbs.Permission{}, "id in (?)", deleteIdList).Error;err!=nil{return err}
	if len(newRows)>0 {
		err=tx.CreateInBatches(newRows, 100).Error
		//err = BatchSave(tx, newRows)
		if err!=nil{return err}
	}

	if err:=tx.Commit().Error; err!=nil{
		tx.Rollback()
		return err
	}
	return nil
}

// BatchSave 批量插入数据
func BatchSave(db *gorm.DB, perms []tbs.Permission) error {
	var buffer bytes.Buffer
	sql := fmt.Sprintf("insert into `%s` (`name`,`path`,`method`) values", models.Permission)
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range perms {
		if i == len(perms)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s');", e.Name,e.Path, e.Method))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s'),", e.Name,e.Path, e.Method))
		}
	}
	return db.Exec(buffer.String()).Error
}

func CheckExist(str string, existList []string) bool {
	for _,v:=range existList{
		if str == v{
			return true
		}
	}
	return false
}
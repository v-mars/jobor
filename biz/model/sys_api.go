package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"jobor/biz/dal/casbin"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/sys_api"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
	"strings"
)

var (
	CategoryApi  = "api"
	Category     = "category"
	CategoryMenu = "menu"
	NameApi      = "api"
	ActionTitle  = map[string]string{
		"get":    "查看",
		"post":   "创建",
		"put":    "修改",
		"patch":  "修补",
		"delete": "删除",
	}
)

type Api struct {
	db.Model
	Name         string `gorm:"type:varchar(158);not null;index:name;comment:名称" json:"name" form:"name"`
	Title        string `json:"title" gorm:"size:158;index:title;comment:标题"`
	Dom          string `json:"dom" gorm:"default:'default';size:128;comment:租户、域"`
	Path         string `gorm:"type:varchar(158);unique_index:idx_name_code;comment:路由路径" json:"path" form:"path"`
	Method       string `gorm:"type:varchar(64);unique_index:idx_name_code;comment:请求方法" json:"method" form:"method"`
	Enabled      bool   `json:"enabled" gorm:"default:1;comment:启用"`
	Group        string `json:"group" gorm:"default:'default';size:64;comment:接口分组"`
	EnabledAudit bool   `json:"enabled_audit" gorm:"default:0;comment:启用审计否"`
	Updater      string `gorm:"type:varchar(64);comment:更新人" json:"updater" form:"updater"`
}

func (d *Api) TableName() string {
	return NameApi
}

func (d *Api) GetApiRpc() *sys_api.ApiResp {
	rpcObj := sys_api.ApiResp{}
	err := convert.AnyToAny(d, &rpcObj)
	if err != nil {
		panic(err)
	}
	return &rpcObj
}

type Apis []Api

func (u *Apis) List(req *sys_api.ApiQuery, resp *response.PageDataList) (Apis, error) {
	resp.List = u
	if err := PageDataWithScopes(db.DB.Model(&Api{}), NameApi, Find, resp,
		GetScopesList(), SelectScopesApi(),
		WhereScopesApi(req),
		OrderScopesApi(), GroupScopesApi()); err != nil {
		return nil, err
	}
	return *u, nil
}
func (u *Apis) ListApi() (uis []*sys_api.ApiResp) {
	if u != nil {
		err := convert.AnyToAny(u, &uis)
		if err != nil {
			panic(err)
		}
	}
	return uis
}
func NewApiModel(Db *gorm.DB) *Api {
	return &Api{Model: db.Model{GormDB: db.DB}}
}

func SelectScopesApi() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		//sql := `api.id,api.*`
		//return Db.Select(sql)
		return Db
	}
}
func SelectAllScopesApi() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,sys_apiname,nickname,status")
	}
}

func WhereScopesApi(req *sys_api.ApiQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "(name like ? or title like ?)"
		var sqlArgs = []interface{}{"%" + req.Name + "%", "%" + req.Name + "%"}
		if req.Dom != "" {
			sql = sql + " and dom=?"
			sqlArgs = append(sqlArgs, req.Dom)
		}
		if req.Method != "" {
			sql = sql + " and method=?"
			sqlArgs = append(sqlArgs, req.Method)
		}
		if req.Path != "" {
			sql = sql + " and path=?"
			sqlArgs = append(sqlArgs, req.Path)
		}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsScopesApi() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Joins("")
	}
}
func PreloadScopesApi(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func OrderScopesApi() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("id desc")
	}
}

func GroupScopesApi() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("api.id")
	}
}

func AddApi(ctx context.Context, Db *gorm.DB, req *sys_api.ApiPost) (Api, error) {
	var row Api
	if err := utils.AnyToAny(req, &row); err != nil {
		return row, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := Db.Table(row.TableName()).Create(&row).Error; err != nil {
		return row, err
	}
	return row, nil
}

func ModApi(ctx context.Context, Db *gorm.DB, _id interface{}, req *sys_api.ApiPut) (Api, error) {
	var mapData map[string]interface{}
	var err error
	if mapData, err = convert.StructToMap(req); err != nil {
		return Api{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var apiObj Api
	if err = tx.First(&apiObj, _id).Error; err != nil {
		return apiObj, err
	}
	if err = tx.Table(NameApi).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return apiObj, err
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return apiObj, err
	}
	hlog.Debug("事务提交成功")
	if err = Db.First(&apiObj, _id).Error; err != nil {
		return apiObj, err
	}
	return apiObj, nil
}

func DelApi(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]Api, error) {
	var us []Api
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&Api{}).Find(&us, "id in (?)", _ids).Error; err != nil {
		return us, err
	}
	for _, _id := range _ids {
		if err := Db.Table(NameApi).Where("name!='prd'").Delete(&Api{}, _id).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("API删除成功")
	return us, nil
}

// GetByApiName 根据API名称获取API信息
func (d *Api) GetByApiName(name string) (Api, error) {
	err := db.DB.Table(d.TableName()).Where("name = ?", name).Take(&d).Error
	if err != nil {
		return Api{}, err
	}
	return *d, nil
}
func (d *Api) UpdateApi(DB *gorm.DB, conditions []Api) error {
	tx := DB.Begin()
	defer func() { tx.Rollback() }()
	var obj []Api

	var cdsStrList []string
	for _, v := range conditions {
		cdsStrList = append(cdsStrList, fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method))
	}

	if err := tx.Table(NameApi).Where(fmt.Sprintf("dom = ? and (%s)", strings.Join(cdsStrList, " or ")),
		conf.Dom).Find(&obj).Error; err != nil {
		return err
	}

	// 过滤新增的列表
	var newRows []Api
	var existList []string
	var existIdList []int
	for _, v := range obj {
		existIdList = append(existIdList, v.Id)
		existList = append(existList, fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method))
	}

	for _, v := range conditions {
		var str1 = fmt.Sprintf("path='%s' and method='%s'", v.Path, v.Method)
		if !CheckExist(str1, existList) {
			newRows = append(newRows, v)
		}
	}

	// 过滤出需要删除的列表
	var deleteList []Api
	var deleteIdList []int
	if err := tx.Find(&deleteList, "dom=? and id not in (?)", conf.Dom, existIdList).Error; err != nil {
		return err
	}

	for _, v := range deleteList {
		deleteIdList = append(deleteIdList, v.Id)
		if _, err := casbin.Enforcer.RemoveFilteredNamedPolicy("p", 1, conf.Dom, v.Path, v.Method); err != nil {
			return err
		}
	}
	if err := tx.Delete(&Api{}, "id in (?)", deleteIdList).Error; err != nil {
		return err
	}
	if len(newRows) > 0 {
		//err = BatchSave(tx, newRows)
		if err := tx.Debug().CreateInBatches(newRows, len(newRows)).Error; err != nil {
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func GetApiByName(ctx context.Context, Db *gorm.DB, name string) (*Api, error) {
	var row = Api{}
	err := Db.Table(NameApi).Where("name=?", name).First(&row).Error
	return &row, err
}

func GetApiInfoById(id interface{}, isPanic bool) (*sys_api.ApiResp, error) {
	var err error
	var u sys_api.ApiResp
	err = db.DB.Table(NameApi).Where("id= ?", id).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the sys_api information does not exist")
		if isPanic {
			panic(err)
		}
		return &sys_api.ApiResp{}, err
	}
	return &u, nil
}

func GetApiInfoByName(name string, isPanic bool) (sys_api.ApiResp, error) {
	var err error
	var u sys_api.ApiResp
	err = db.DB.Table(NameApi).Where("name = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the sys_api information does not exist")
		if isPanic {
			panic(err)
		}
		return sys_api.ApiResp{}, err
	}
	return u, nil
}

func CheckExist(str string, existList []string) bool {
	for _, v := range existList {
		if str == v {
			return true
		}
	}
	return false
}

func UpdateApiByRoutes(db *gorm.DB, e *server.Hertz) {
	perms := Api{}
	var res []Api
	routes := e.Routes()
	for _, v := range routes {
		var name string
		pathArray := strings.Split(strings.Trim(v.Path, "/"), "/")
		title := ""
		group := "default"
		audit := false
		if strings.ToLower(v.Method) != "get" {
			audit = true
		}
		if len(pathArray) >= 5 {
			group = pathArray[2]
			title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(v.Method)], pathArray[3])
			name = fmt.Sprintf("%s:%s:%s:%s", strings.ToLower(v.Method), pathArray[2], pathArray[3], pathArray[4])
		} else if len(pathArray) >= 4 {
			group = pathArray[2]
			title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(v.Method)], pathArray[3])
			name = fmt.Sprintf("%s:%s:%s", strings.ToLower(v.Method), pathArray[2], pathArray[3])
		} else if len(pathArray) == 3 {
			group = pathArray[2]
			title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(v.Method)], pathArray[2])
			name = fmt.Sprintf("%s:%s", strings.ToLower(v.Method), pathArray[2])
		} else if len(pathArray) == 2 {
			title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(v.Method)], pathArray[1])
			name = fmt.Sprintf("%s:%s", strings.ToLower(v.Method), pathArray[1])
		} else if len(pathArray) == 1 {
			title = fmt.Sprintf("%s%s", ActionTitle[strings.ToLower(v.Method)], pathArray[0])
			name = fmt.Sprintf("%s:%s", strings.ToLower(v.Method), pathArray[0])
		}
		res = append(res, Api{Title: title, Name: name, Method: v.Method, Path: v.Path, Group: group, EnabledAudit: audit})
	}
	err := perms.UpdateApi(db, res)
	if err != nil {
		hlog.Errorf("update sys api by routes: %s", err)
	}
	hlog.Infof("update sys api by routes is success")
}

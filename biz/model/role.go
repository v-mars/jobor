package model

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"jobor/biz/dal/casbin"
	"jobor/biz/dal/db"
	"jobor/biz/response"
	utils2 "jobor/biz/utils"
	"jobor/conf"
	"jobor/kitex_gen/role"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

const (
	NameRole = "role"
	//dom            = "sys"
	ROLE_DEV       = "dev"
	ROLE_OPS       = "ops"
	ROLE_DEV_OWNER = "dev_owner"
	ROLE_QA        = "test"
)

type Role struct {
	db.Model
	Title       string      `gorm:"type:varchar(128);unique_index:idx_title_code;not null;comment:名称" json:"title" form:"title"`
	Name        string      `gorm:"type:varchar(128);unique:uni_name;not null;comment:名称" json:"name" form:"name"`
	Description string      `gorm:"type:longtext;comment:描述" json:"description" form:"description"`
	ParentID    int         `gorm:"index:parent_id;comment:父节点" json:"parent_id" form:"parent_id"`
	Path        db.IntArray `json:"path" gorm:"type:text;comment:节点路径"`
	Sort        int         `json:"sort" gorm:"default:1000;comment:'排序'"`
	Apis        []Api       `gorm:"many2many:role_api;association_autoupdate:false;association_autocreate:false;constraint:OnDelete:CASCADE" json:"apis"`
	Children    []*Role     `gorm:"-" json:"children"`
	//Updater     string        `gorm:"type:varchar(64);comment:更新人" json:"updater" form:"updater"`
}

func (*Role) TableName() string {
	return NameRole
}

type Roles []Role

func (u *Roles) List(req *role.RoleQuery, resp *response.PageDataList) (Roles, error) {
	resp.List = u
	if err := PageDataWithScopes(db.DB.Model(&Role{}), NameRole, Find, resp,
		GetScopesList(), SelectScopesRole(),
		WhereScopesRole(req), PreloadScopes("Apis"),
		OrderScopesRole(), GroupScopesRole()); err != nil {
		return nil, err
	}
	return *u, nil
}
func (u *Roles) ListRole() (uis []*role.RoleResp) {
	if u != nil {
		err := convert.AnyToAny(u, &uis)
		if err != nil {
			panic(err)
		}
	}
	return uis
}

func SelectScopesRole() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct *")
	}
}
func WhereScopesRole(req *role.RoleQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "(role.parent_id = 0 or role.parent_id is null) and (role.name like ? or role.title like ?)"
		var sqlArgs = []interface{}{"%" + req.Name + "%", "%" + req.Title + "%"}
		if req.ClientName != "" {
			sql = sql + " and sc.name=?"
			sqlArgs = append(sqlArgs, req.ClientName)
		}
		return Db.Where(sql, sqlArgs...)
	}
}
func JoinsScopesRole() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		join := `
left join role_clients rc on rc.role_id=role.id
left join sso_client sc on sc.id=rc.sso_client_id
`
		return Db.Joins(join)
	}
}
func OrderScopesRole() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("sort asc")
	}
}

func GroupScopesRole() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("")
	}
}

func AddRole(ctx context.Context, Db *gorm.DB, req *role.RolePost) (Role, error) {
	var row Role
	if err := utils.AnyToAny(req, &row); err != nil {
		return row, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&Api{}).Where("id in (?)", req.ApiIds).Select(
		"id,name,path,method,dom").Scan(&row.Apis).Error; err != nil {
		return row, err
	}
	var ok bool
	var err error
	if err = tx.Table(NameRole).Create(&row).Error; err != nil {
		return row, err
	}
	hlog.Debugf("获取关联API成功, API count %d", len(row.Apis))
	for _, v := range row.Apis {
		if _, err := casbin.Enforcer.AddPolicy(row.Name, conf.Dom, v.Path, v.Method); err != nil {
			return row, err
		}
		hlog.Debug("add casbin policy %s %s %s is %t", row.Name, v.Path, v.Method, ok)
	}
	if row.ParentID != 0 {
		var newParentRole Role
		if err = tx.First(&newParentRole, row.ParentID).Error; err != nil {
			return row, err
		}
		hlog.Debugf("start add casbin g parent role name %s, current role name %s", newParentRole.Name, row.Name)
		ok, err = casbin.Enforcer.AddNamedGroupingPolicy("g", newParentRole.Name, row.Name, conf.Dom)
		if err != nil {
			return row, err
		}
		hlog.Debugf("add casbin g parent role name %s, current role name %s is %t", newParentRole.Name, row.Name, ok)
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("add 事务提交失败，%s", err)
		return row, err
	}
	hlog.Debug("add 事务提交成功")
	return row, nil
}

func ModRole(ctx context.Context, Db *gorm.DB, _id interface{}, req *role.RolePut) (Role, error) {
	if req.ParentId != nil && convert.ToInt64(_id) == *req.ParentId {
		*req.ParentId = 0
	}
	var mapData map[string]interface{}
	var err error
	if err = convert.StructToMapOut(req, &mapData); err != nil {
		return Role{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var roleObj Role
	if err = tx.First(&roleObj, "role.id=?", _id).Error; err != nil {
		return roleObj, err
	}
	if req.Path != nil {
		if len(req.Path.Values) > 9 {
			return roleObj, fmt.Errorf("角色层级深度不能超过10级")
		}
		mapData["path"] = db.IntArray(req.GetPathInt())
	}
	var ok bool
	if req.ApiIds != nil {
		var apis []Api
		if err = tx.Model(&Api{}).Where("id in (?)", req.GetApiIdsInt()).Select("id,name,path,method,dom").Scan(&apis).Error; err != nil {
			return roleObj, err
		}
		hlog.Debug("get api list success")

		var newStrArray []string
		for _, v := range apis {
			v := v
			newStrArray = append(newStrArray, fmt.Sprintf("%s:%s:%s:%s", roleObj.Name, dom, v.Path, v.Method))
			if ok, err = casbin.Enforcer.AddPolicy(roleObj.Name, dom, v.Path, v.Method); err != nil {
				return roleObj, err
			}
			hlog.Debugf("casbin new policy %s %s %s add is %t", roleObj.Name, v.Path, v.Method, ok)
		}

		existsList := casbin.Enforcer.GetPermissionsForUser(roleObj.Name, dom)
		for _, v := range existsList {
			var strTmp = strings.Join(v, ":")
			if len(v) == 4 && !utils.InOfStr(strTmp, newStrArray) {
				if ok, err = casbin.Enforcer.RemovePolicy(roleObj.Name, dom, v[2], v[3]); err != nil {
					return roleObj, err
				}
				hlog.Debug("casbin old policy %s %s %s remove is %t", roleObj.Name, v[2], v[3], ok)
			}
		}

		if err = tx.Model(&roleObj).Association("Apis").Replace(apis); err != nil {
			return roleObj, err
		}
		hlog.Infof("role 关联 api 成功")
	}

	if req.ParentId != nil && *req.ParentId != 0 {
		var newParentRole Role
		if err = tx.First(&newParentRole, *req.ParentId).Error; err != nil {
			return roleObj, err
		}
		if roleObj.ParentID != 0 && roleObj.ParentID != int(*req.ParentId) {
			var oldParentRole Role
			if err = tx.First(&oldParentRole, roleObj.ParentID).Error; err != nil {
				return roleObj, err
			}
			hlog.Debugf("start remove casbin g old parent role name %s, current role name %s", oldParentRole.Name, roleObj.Name)
			ok, err = casbin.Enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, oldParentRole.Name, roleObj.Name, conf.Dom)
			if err != nil {
				return roleObj, err
			}
			hlog.Debugf("remove casbin g old parent role name %s, current role name %s is %t", oldParentRole.Name, roleObj.Name, ok)
		}
		hlog.Debugf("start add casbin g parent role name %s, current role name %s", newParentRole.Name, roleObj.Name)
		ok, err = casbin.Enforcer.AddNamedGroupingPolicy("g", newParentRole.Name, roleObj.Name, conf.Dom)
		if err != nil {
			return roleObj, err
		}
		hlog.Debugf("add casbin g parent role name %s, current role name %s is %t", newParentRole.Name, roleObj.Name, ok)
	} else if req.ParentId != nil && *req.ParentId == 0 && roleObj.ParentID != 0 {
		var oldParentRole Role
		if err = tx.First(&oldParentRole, roleObj.ParentID).Error; err != nil {
			return roleObj, err
		}
		hlog.Debugf("start remove casbin g current role name %s", roleObj.Name)
		ok, err = casbin.Enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, oldParentRole.Name, roleObj.Name, conf.Dom)
		if err != nil {
			return roleObj, err
		}
		hlog.Debugf("remove casbin g current role name %s is %t", roleObj.Name, ok)
	}
	if err = tx.Table(NameRole).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return roleObj, err
	}

	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return roleObj, err
	}
	hlog.Debug("事务提交成功")
	if err = Db.First(&roleObj, _id).Error; err != nil {
		return roleObj, err
	}
	return roleObj, nil
}

func DelRole(ctx context.Context, Db *gorm.DB, _ids []interface{}) (Role, error) {
	var obj = Role{}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var deleteList []Role
	var childList []Role
	if err := tx.Model(&Role{}).Find(&deleteList, "id in (?)", _ids).Error; err != nil {
		return obj, err
	}
	if err := tx.Model(&Role{}).Find(&childList, "parent_id in (?)", _ids).Error; err != nil {
		return obj, err
	}
	if len(childList) > 0 {
		return obj, fmt.Errorf("角色包含子角色不能删除")
	}
	hlog.CtxDebugf(ctx, "获取删除的关联角色成功")
	for _, v := range deleteList {
		//if err := tx.Model(&Role{}).Find(&deleteList, "id in (?)", _ids).Error; err != nil {
		//	return obj, err
		//}
		if _, err := casbin.Enforcer.RemoveFilteredNamedPolicy("p", 0, v.Name, dom); err != nil {
			return obj, err
		}
		if _, err := casbin.Enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, v.Name, "", dom); err != nil {
			return obj, err
		}
		if _, err := casbin.Enforcer.RemoveFilteredNamedGroupingPolicy("g", 1, v.Name, dom); err != nil {
			return obj, err
		}
	}
	hlog.CtxDebugf(ctx, "角色关联casbin policy删除成功")
	for _, _id := range _ids {
		if err := Db.Model(&Role{Model: db.Model{Id: convert.ToInt(_id)}}).Association("Apis").Clear(); err != nil {
			return obj, err
		}
		hlog.Debug("角色关联API删除成功")
		if err := Db.Table(NameRole).Where("name!=?", "admin").Delete(&obj, _id).Error; err != nil {
			return obj, err
		}
	}
	hlog.CtxInfof(ctx, "角色删除成功")
	return obj, nil
}

// GetByRoleName 根据环境名称获取环境信息
func (d *Role) GetByRoleName(name string) (Role, error) {
	err := db.DB.Table(d.TableName()).Where("name = ?", name).Take(&d).Error
	if err != nil {
		return Role{}, err
	}
	return *d, nil
}

func GetRoleByName(ctx context.Context, Db *gorm.DB, name string) (*Role, error) {
	var row = Role{}
	err := Db.Table(NameRole).Where("name=?", name).First(&row).Error
	return &row, err
}

func GetRoleInfoById(id interface{}, isPanic bool) (*Role, error) {
	var err error
	var u Role
	err = db.DB.Table(NameRole).Where("id= ?", id).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the role information does not exist")
		if isPanic {
			panic(err)
		}
		return &Role{}, err
	}
	return &u, nil
}

func GetRoleInfoByName(name string, isPanic bool) (Role, error) {
	var err error
	var u Role
	err = db.DB.Table(NameRole).Where("name = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Name == "" && u.Id == 0 {
		err = fmt.Errorf("the role information does not exist")
		if isPanic {
			panic(err)
		}
		return Role{}, err
	}
	return u, nil
}

func (d Role) GetMapTree(DB *gorm.DB, admin bool) (map[int]map[string]interface{}, error) {
	var data []Role
	var where = "name!='admin'" // name!='admin'
	if admin {
		where = ""
	}
	if err := DB.Model(&Role{}).Preload("Apis").Order("parent_id asc,sort asc").Find(&data, where).Error; err != nil {
		return nil, err
	}
	var MapData []map[string]interface{}
	if len(data) == 0 {
		return map[int]map[string]interface{}{}, nil
	}
	var err error
	if MapData, err = convert.StructToMapSlice(data); err != nil {
		return nil, err
	}
	var treeData = map[int]map[string]interface{}{}
	if len(MapData) > 0 {
		treeData = utils2.ListToTree(MapData)
	}
	return treeData, nil
}

func (d Role) GetListTree(DB *gorm.DB, admin bool) ([]map[string]interface{}, error) {
	treeMap, err := d.GetMapTree(DB, admin)
	if err != nil {
		return nil, err
	}
	if len(treeMap) == 0 {
		return []map[string]interface{}{}, err
	}
	var dataList = treeMap[0]["children"].([]map[string]interface{})
	return dataList, nil
}

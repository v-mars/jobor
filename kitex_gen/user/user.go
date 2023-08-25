package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"jobor/biz/dal/casbin"
	"jobor/biz/dal/db"
	"jobor/biz/model"
	"jobor/biz/response"
	"jobor/kitex_gen/task"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
)

const (
	TbName       = "user"
	Local        = "local"
	Ldap         = "ldap"
	SSO          = "sso"
	HeaderTag    = "userinfo"
	ErrMfaFailed = "多因子（MFA/OTP）认证失败"
)

type User struct {
	db.ModelOld
	Nickname string `gorm:"type:varchar(128);comment:显示名" json:"nickname" form:"nickname"`
	Username string `gorm:"type:varchar(128);unique:uni_name;comment:用户名" json:"username" form:"username"` // `unique_index` also works
	Password string `gorm:"type:varchar(256);comment:密码" json:"password" form:"password"`
	Email    string `gorm:"type:varchar(156);comment:邮箱" json:"email" column:"email"`
	Phone    string `gorm:"type:varchar(64);comment:电话" json:"phone" column:"phone"`
	Userid   string `gorm:"type:varchar(156);comment:UserId" json:"userid" form:"userid"`
	Empno    string `gorm:"type:varchar(156);comment:员工号" json:"empno" form:"empno"`
	UserType string `gorm:"type:varchar(16);default:local;comment:用户类型:local,ldap,sso" json:"user_type" form:"user_type"`
	Avatar   string `gorm:"type:varchar(64);default:default;comment:用户头像" json:"avatar" form:"avatar"`
	Status   bool   `gorm:"type:varchar(32);default:true;comment:状态" json:"status" form:"status"`
	//Roles    []role.Role `gorm:"many2many:user_roles;association_autoupdate:false;association_autocreate:false;constraint:OnDelete:CASCADE" json:"roles"` // Many-To-Many
	////Updater  string `gorm:"type:varchar(64);comment:更新人" json:"updater" form:"updater"`
	////OtpSecret  string          `gorm:"type:varchar(128);default:null;comment:OTP双因子认证密钥" json:"otp_secret" form:"otp_secret"`
	////DistinguishedName string `json:"distinguished_name" column:"distinguished_name"`
	////IsDemission       int8   `json:"is_demission" column:"is_demission"`
	////StaffType         int8   `json:"staff_type" column:"staff_type"`
	////DepartmentId      int64  `json:"department_id" column:"department_id"`
	////ManagerId int64 `json:"manager_id" column:"manager_id"`
}

func (u *User) TableName() string {
	return TbName
}

func (u *User) GetRoles() []string {
	var roles []string
	//for _, r := range u.Roles {
	//	roles = append(roles, r.Name)
	//}
	return roles
}

func (u *User) GetUserinfo() *Userinfo {
	userinfo := Userinfo{
		Id: int64(u.ID), Username: u.Username, Nickname: u.Nickname, Email: u.Email, Avatar: u.Avatar,
		Empno: u.Empno, UserType: u.UserType, Roles: u.GetRoles(), Status: u.Status,
	}
	return &userinfo
}

type Users []User

func (u *Users) List(req *UserQuery, resp *response.PageDataList) (Users, error) {
	resp.List = u
	//resp := response.PageDataList{List: &users}
	if err := model.PageDataWithScopes(db.DB.Model(&User{}), TbName, model.Find, resp,
		model.GetScopesList(), //SelectScopes(),
		WhereScopes(req), task.PreloadScopes("Roles"),
		OrderScopes(), GroupScopes()); err != nil {
		return nil, err
	}
	return *u, nil
}
func (u *Users) ListUserinfo() (uis []*Userinfo) {
	if u != nil {
		for _, v := range *u {
			v := v
			uis = append(uis, v.GetUserinfo())
		}
	}
	return uis
}
func NewModel(Db *gorm.DB) *User {
	return &User{ModelOld: db.ModelOld{GormDB: db.DB}}
}

func SelectScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `user.id,nickname,username,email,phone,userid,empno,user_type,avatar,status,user.ctime,user.*`
		return Db.Select(sql)
	}
}
func SelectAllScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,username,nickname,status")
	}
}

func WhereScopes(req *UserQuery) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		var sql = "(username like ? or nickname like ?)"
		var sqlArgs = []interface{}{"%" + req.Name + "%", "%" + req.Name + "%"}
		if len(req.UserType) > 0 {
			sql = sql + " and user_type=?"
			sqlArgs = append(sqlArgs, req.UserType)
		}
		if len(req.Username) > 0 {
			sql = sql + " and username=?"
			sqlArgs = append(sqlArgs, req.Username)
		}
		if len(req.Nickname) > 0 {
			sql = sql + " and nickname=?"
			sqlArgs = append(sqlArgs, req.Nickname)
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
		return Db.Order("id desc")
	}
}

func GroupScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("user.id")
	}
}

var dom = "sys"

func Add(ctx context.Context, Db *gorm.DB, req *PostUserReq) (User, error) {
	req.Password = utils.HashAndSalt([]byte(req.Password))
	var row User
	if err := utils.AnyToAny(req, &row); err != nil {
		return row, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	//if err := tx.Model(&role.Role{}).Where("id in (?)", req.RoleIds).Select(
	//	"id,name").Scan(&row.Roles).Error; err != nil {
	//	return row, err
	//}
	//hlog.Debugf("获取关联Roles成功, roles count %d", len(row.Roles))
	//for _, v := range row.Roles {
	//	_, err := casbin.Enforcer.AddGroupingPolicy(row.Username, v.Name, dom) // user role dom
	//	if err != nil {
	//		return row, err
	//	}
	//}
	hlog.CtxDebugf(ctx, "user %s casbin get role is success", row.Nickname)
	if err := tx.Table(row.TableName()).Create(&row).Error; err != nil {
		return row, err
	}
	if err := tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return row, err
	}
	row.Password = ""
	return row, nil
}

func Mod(ctx context.Context, Db *gorm.DB, _id interface{}, req *PutUserReq) (User, error) {
	var mapData map[string]interface{}
	var err error
	if mapData, err = convert.StructToMap(req); err != nil {
		return User{}, err
	}
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	var userObj User
	if err = tx.First(&userObj, _id).Error; err != nil {
		return userObj, err
	}
	//if req.RoleIds != nil {
	//	if err = tx.Model(&role.Role{}).Where("id in (?)", req.GetRoleIdsInt()).Scan(&userObj.Roles).Error; err != nil {
	//		return userObj, err
	//	}
	//
	//	var newStrArray []string
	//	for _, v := range userObj.Roles {
	//		newStrArray = append(newStrArray, v.Name)
	//		_, err = casbin.Enforcer.AddGroupingPolicy(userObj.Username, v.Name, dom) // user role dom
	//		if err != nil {
	//			return userObj, err
	//		}
	//	}
	//	hlog.CtxDebugf(ctx, "user %s casbin add  group policy is success", userObj.Nickname)
	//	existsList, err := casbin.Enforcer.GetRolesForUser(userObj.Username, dom)
	//	if err != nil {
	//		return userObj, err
	//	}
	//	hlog.CtxDebugf(ctx, "user %s casbin get role is success", userObj.Nickname)
	//	diff := utils.Difference(existsList, newStrArray)
	//	for _, v := range diff {
	//		_, err = casbin.Enforcer.RemoveGroupingPolicy(userObj.Username, v, dom)
	//		if err != nil {
	//			return userObj, err
	//		}
	//	}
	//	hlog.CtxDebugf(ctx, "user %s casbin remove  group policy is success", userObj.Nickname)
	//	if err = tx.Model(&userObj).Association("Roles").Replace(userObj.Roles); err != nil {
	//		return userObj, err
	//	}
	//}
	if err = tx.Table(TbName).Where("id=?", _id).Updates(mapData).Error; err != nil {
		return userObj, err
	}
	if err = tx.Commit().Error; err != nil {
		hlog.Errorf("事务提交失败，%s", err)
		return userObj, err
	}
	hlog.Debug("事务提交成功")
	if err = Db.First(&userObj, _id).Error; err != nil {
		return userObj, err
	}
	userObj.Password = ""
	return userObj, nil
}

func Del(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]User, error) {
	var us []User
	tx := Db.Begin()
	defer func() { tx.Rollback() }()
	if err := tx.Model(&User{}).Find(&us, "id in (?)", _ids).Error; err != nil {
		return us, err
	}
	for i, v := range us {
		if _, err := casbin.Enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, v.Username, "", dom); err != nil {
			return us, err
		}
		us[i].Password = ""
	}
	hlog.CtxDebugf(ctx, "用户关联casbin group policy删除成功")
	for _, _id := range _ids {
		if err := Db.Model(&User{ModelOld: db.ModelOld{ID: convert.ToInt(_id)}}).Association("Roles").Clear(); err != nil {
			return us, err
		}
		hlog.Debug("用户关联角色删除成功")
		if err := Db.Table(TbName).Where("username!='root'").Delete(&User{}, _id).Error; err != nil {
			return us, err
		}
	}
	hlog.Infof("用户删除成功")
	return us, nil
}

// GetByUserName 根据用户名称获取用户信息
func (u *User) GetByUserName(username string) (User, error) {
	err := db.DB.Table(u.TableName()).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return User{}, err
	}
	return *u, nil
}

func Auth(DB *gorm.DB, username, password string) (u User, ok bool, err error) {
	//var u User
	defer func() {
		if err != nil {
			hlog.Errorf("username %s auth failed, %s", username, err)
		}
	}()
	if err = DB.Model(&User{}).Where(
		"username=? and user_type=?", username, Local).First(&u).Error; !errors.Is(
		err, gorm.ErrRecordNotFound) && err != nil {
		return u, false, nil
	}
	if !utils.ValidateSaltPasswords(u.Password, []byte(password)) {
		return u, false, fmt.Errorf("用户[%s]认证失败，用户名或密码不对,请重新输入", username)
	}
	if !u.Status {
		return u, true, fmt.Errorf("用户[%s]已经被禁用，请联系管理员", username)
	}
	return u, true, nil
}

func GetUserByUsername(ctx context.Context, Db *gorm.DB, username string) (*User, error) {
	var row = User{}
	err := Db.Table(TbName).Where("username=?", username).First(&row).Error
	return &row, err
}

func GetUserinfoById(id interface{}, isPanic bool) (*Userinfo, error) {
	var err error
	var u Userinfo
	err = db.DB.Table(TbName).Where("id= ?", id).Omit("Roles").Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return &u, err
	}
	if u.Username == "" && u.Id == 0 {
		err = fmt.Errorf("the user information does not exist")
		if isPanic {
			panic(err)
		}
		return &Userinfo{}, err
	}
	return &u, nil
}

func GetUserinfoByUsername(name string, isPanic bool) (Userinfo, error) {
	var err error
	var u Userinfo
	err = db.DB.Table(TbName).Where("username = ?", name).Take(&u).Error
	if err != nil {
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Username == "" && u.Id == 0 {
		err = fmt.Errorf("the user information does not exist")
		if isPanic {
			panic(err)
		}
		return Userinfo{}, err
	}
	return u, nil
}

func GetUserValue(c *app.RequestContext, isPanic bool) (Userinfo, error) {
	userInter, ok := c.Get(HeaderTag)
	var err error
	//var u, ook = userInter.(UserInfo)
	if !ok {
		err = fmt.Errorf("the user information does not exist or is incorrect")
		if isPanic {
			panic(err)
		}
		return Userinfo{}, err
	}
	var u Userinfo
	if e := utils.AnyToAny(userInter, &u); e != nil {
		err = fmt.Errorf("the user information parse is err: %s", e)
		if isPanic {
			panic(err)
		}
		return u, err
	}
	if u.Username == "" && u.Id == 0 {
		err = fmt.Errorf("the user information does not exist")
		if isPanic {
			panic(err)
		}
		return u, err
	}
	u.Roles, err = GetUserRoles(u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUserinfoOrCreate(ui *Userinfo) (Userinfo, error) {
	var err error
	var u Userinfo
	var row User
	err = db.DB.Table(TbName).Where("username = ?", ui.Username).Take(&u).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return u, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		row = User{Username: u.Username, Nickname: u.Nickname, Email: u.Email, Status: true, UserType: Ldap}
		if err = db.DB.Table(TbName).Omit("id").Create(&row).Error; err != nil {
			return u, err
		}
		u.Id = int64(row.ID)
		u.UserType = row.UserType
		u.Status = row.Status
	}
	u.Roles, err = GetUserRoles(u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUserMap(ids []int) (r map[int]User, err error) {
	r = make(map[int]User)
	var us []User
	if err = db.DB.Model(&User{}).Where("user.id in (?)", ids).Find(&us).Error; err != nil {
		return nil, err
	}
	for _, v := range us {
		v := v
		r[v.ID] = v
	}
	return r, nil
}

func (i *Userinfo) IsAdmin() bool {
	if utils.InOfStr("admin", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsDevOps() bool {
	if utils.InOfStr("devops", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsOps() bool {
	if utils.InOfStr("ops", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsQa() bool {
	if utils.InOfStr("qa", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) IsDeveloper() bool {
	if utils.InOfStr("developer", i.Roles) {
		return true
	} else {
		return false
	}
}

func (i *Userinfo) JsonString() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(marshal)
}

func (i *Userinfo) Map() (mapRes map[string]interface{}) {
	err := utils.AnyToAny(i, &mapRes)
	if err != nil {
		panic(err)
	}
	return mapRes
}

func GetUserSession(c *app.RequestContext, isPanic bool) (Userinfo, error) {
	userInter, ok := c.Get(headerTag)
	var err error
	//var u, ook = userInter.(UserInfo)
	if !ok {
		err = fmt.Errorf("the user information does not exist or is incorrect")
		if isPanic {
			panic(err)
		}
		return Userinfo{}, err
	}
	var u Userinfo
	if e := utils.AnyToAny(userInter, &u); e != nil {
		err = fmt.Errorf("the user information parse is err: %s", e)
		if isPanic {
			panic(err)
		}
		return Userinfo{}, err
	}
	if u.Username == "" && u.Id == 0 {
		err = fmt.Errorf("the user information does not exist")
		if isPanic {
			panic(err)
		}
		return Userinfo{}, err
	}
	u.Roles, err = GetUserRoles(u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUserinfoFromOidc(body json.RawMessage) Userinfo {
	var u Userinfo
	gj := gjson.Parse(string(body))
	u.Id = gj.Get("claims.id").Int()
	u.Username = gj.Get("claims.username").String()
	u.Nickname = gj.Get("claims.nickname").String()
	u.Email = gj.Get("claims.email").String()
	//u.Roles = gj.Get("claims.roles").Array()
	for _, v := range gj.Get("claims.roles").Array() {
		v := v
		u.Roles = append(u.Roles, v.String())
	}
	return u
}

func SetContentUserinfo(ctx context.Context, c *app.RequestContext, u Userinfo) {
	c.Set(uid, u.Id)
	c.Set(username, u.Username)
	c.Set(nickname, u.Nickname)
	c.Set(headerTag, u)
}

func InitSession(ctx context.Context, session sessions.Session, userInfo Userinfo, clientId, org string, isLogin bool) error {
	//session := sessions.Default(c)
	session.Set(IsLogin, isLogin)
	session.Set(uid, userInfo.Id)
	session.Set(username, userInfo.Username)
	session.Set(nickname, userInfo.Nickname)
	session.Set("email", userInfo.Email)
	session.Set("roles", userInfo.Roles)
	session.Set("org", org)
	session.Set("client_id", clientId)
	session.Set(headerTag, userInfo.Map())
	if err := session.Save(); err != nil {
		return err
	}
	hlog.CtxDebugf(ctx, "login user %s write session cookie is success", userInfo.Username)
	return nil
}

func (p *PutUserReq) GetRoleIdsInt() []int {
	var arrayInt = make([]int, 0)
	if p.RoleIds != nil {
		for _, v := range p.RoleIds.GetValues() {
			v := v
			arrayInt = append(arrayInt, int(v.GetNumberValue()))
		}
	}
	return arrayInt
}

func GetUserRoles(username string) ([]string, error) {
	//var roles []string
	//	var sql = `SELECT DISTINCT role.name FROM role
	//LEFT JOIN user_roles ON user_roles.role_id = role.id
	//LEFT JOIN user ON user_roles.user_id = user.id
	//WHERE (user.username = ?)
	//`
	//if err := db.DB.Model(&role.Role{}).Raw(sql, username).Pluck("role.name", &roles).Error; err != nil {
	//	return roles, err
	//} else {
	//	return roles, nil
	//}
	return nil, nil
}

var (
	headerTag = "userinfo"
	username  = "username"
	nickname  = "nickname"
	uid       = "uid"
	IsLogin   = "islogin"
)

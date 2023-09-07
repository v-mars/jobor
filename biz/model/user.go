package model

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
	"jobor/biz/response"
	"jobor/kitex_gen/user"
	"jobor/pkg/convert"
	"jobor/pkg/utils"
)

const (
	NameUser     = "user"
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
	return NameUser
}

func (u *User) GetRoles() []string {
	var roles []string
	//for _, r := range u.Roles {
	//	roles = append(roles, r.Name)
	//}
	return roles
}

func (u *User) GetUserinfo() *user.Userinfo {
	userinfo := user.Userinfo{
		Id: int64(u.ID), Username: u.Username, Nickname: u.Nickname, Email: u.Email, Avatar: u.Avatar,
		Empno: u.Empno, UserType: u.UserType, Roles: u.GetRoles(), Status: u.Status,
	}
	return &userinfo
}

type Users []User

func (u *Users) List(req *user.UserQuery, resp *response.PageDataList) (Users, error) {
	resp.List = u
	//resp := response.PageDataList{List: &users}
	if err := PageDataWithScopes(db.DB.Model(&User{}), NameUser, Find, resp,
		GetScopesList(), //UserSelectScopes(),
		UserWhereScopes(req), PreloadScopes("Roles"),
		UserOrderScopes(), UserGroupScopes()); err != nil {
		return nil, err
	}
	return *u, nil
}
func (u *Users) ListUserinfo() (uis []*user.Userinfo) {
	if u != nil {
		for _, v := range *u {
			v := v
			uis = append(uis, v.GetUserinfo())
		}
	}
	return uis
}
func NewUserModel(Db *gorm.DB) *User {
	return &User{ModelOld: db.ModelOld{GormDB: db.DB}}
}

func UserSelectScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		// distinct
		sql := `user.id,nickname,username,email,phone,userid,empno,user_type,avatar,status,user.ctime,user.*`
		return Db.Select(sql)
	}
}
func UserSelectAllScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Select("distinct id,username,nickname,status")
	}
}

func UserWhereScopes(req *user.UserQuery) func(Db *gorm.DB) *gorm.DB {
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
func UserJoinsScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Joins("")
	}
}
func UserPreloadScopes(query string, args ...interface{}) func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Preload(query, args...)
	}
}

func UserOrderScopes() func(db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Order("id desc")
	}
}

func UserGroupScopes() func(Db *gorm.DB) *gorm.DB {
	return func(Db *gorm.DB) *gorm.DB {
		return Db.Group("user.id")
	}
}

var dom = "sys"

func UserAdd(ctx context.Context, Db *gorm.DB, req *user.PostUserReq) (User, error) {
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

func UserMod(ctx context.Context, Db *gorm.DB, _id interface{}, req *user.PutUserReq) (User, error) {
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
	if err = tx.Table(NameUser).Where("id=?", _id).Updates(mapData).Error; err != nil {
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

func UserDel(ctx context.Context, Db *gorm.DB, _ids []interface{}) ([]User, error) {
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
		if err := Db.Table(NameUser).Where("username!='root'").Delete(&User{}, _id).Error; err != nil {
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
	err := Db.Table(NameUser).Where("username=?", username).First(&row).Error
	return &row, err
}

func GetUserinfoById(id interface{}, isPanic bool) (*user.Userinfo, error) {
	var err error
	var u user.Userinfo
	err = db.DB.Table(NameUser).Where("id= ?", id).Omit("Roles").Take(&u).Error
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
		return &user.Userinfo{}, err
	}
	return &u, nil
}

func GetUserinfoByUsername(name string, isPanic bool) (user.Userinfo, error) {
	var err error
	var u user.Userinfo
	err = db.DB.Table(NameUser).Where("username = ?", name).Take(&u).Error
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
		return user.Userinfo{}, err
	}
	return u, nil
}

func GetUserValue(c *app.RequestContext, isPanic bool) (user.Userinfo, error) {
	userInter, ok := c.Get(HeaderTag)
	var err error
	//var u, ook = userInter.(UserInfo)
	if !ok {
		err = fmt.Errorf("the user information does not exist or is incorrect")
		if isPanic {
			panic(err)
		}
		return user.Userinfo{}, err
	}
	var u user.Userinfo
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

func GetUserinfoOrCreate(ui *user.Userinfo) (user.Userinfo, error) {
	var err error
	var u user.Userinfo
	var row User
	err = db.DB.Table(NameUser).Where("username = ?", ui.Username).Take(&u).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return u, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		row = User{Username: u.Username, Nickname: u.Nickname, Email: u.Email, Status: true, UserType: Ldap}
		if err = db.DB.Table(NameUser).Omit("id").Create(&row).Error; err != nil {
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

func GetUserSession(c *app.RequestContext, isPanic bool) (user.Userinfo, error) {
	userInter, ok := c.Get(headerTag)
	var err error
	//var u, ook = userInter.(UserInfo)
	if !ok {
		err = fmt.Errorf("the user information does not exist or is incorrect")
		if isPanic {
			panic(err)
		}
		return user.Userinfo{}, err
	}
	var u user.Userinfo
	if e := utils.AnyToAny(userInter, &u); e != nil {
		err = fmt.Errorf("the user information parse is err: %s", e)
		if isPanic {
			panic(err)
		}
		return user.Userinfo{}, err
	}
	if u.Username == "" && u.Id == 0 {
		err = fmt.Errorf("the user information does not exist")
		if isPanic {
			panic(err)
		}
		return user.Userinfo{}, err
	}
	u.Roles, err = GetUserRoles(u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetUserinfoFromOidc(body json.RawMessage) user.Userinfo {
	var u user.Userinfo
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

func SetContentUserinfo(ctx context.Context, c *app.RequestContext, u user.Userinfo) {
	c.Set(uid, u.Id)
	c.Set(username, u.Username)
	c.Set(nickname, u.Nickname)
	c.Set(headerTag, u)
}

func InitSession(ctx context.Context, session sessions.Session, userInfo user.Userinfo, clientId, org string, isLogin bool) error {
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

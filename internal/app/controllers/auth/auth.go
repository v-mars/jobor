package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	role2 "jobor/internal/app/controllers/sys/role"
	sys "jobor/internal/app/controllers/sys/user"
	"jobor/internal/app/models/db"
	"jobor/internal/app/models/tbs"
	"jobor/internal/app/response"
	"jobor/internal/config"
	"jobor/pkg/convert"
	goJWT "jobor/pkg/jwt"
	"jobor/pkg/logger"
	"jobor/pkg/utils"
	"log"
	"time"
)


// @Tags 登录认证
// @Summary 登录
// @Description 登录
// @Produce  json
// @Param payload body  Params true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/login [post]
func LoginAuth(c *gin.Context)  {
	var obj Params
	if err:= c.ShouldBindJSON(&obj);err!=nil{
		response.Error(c, err)
		return
	}

	var localAuth int64
	if err:=db.DB.Model(&tbs.User{}).Where("username = ? and password = ? and user_type_id = ?",
		obj.Username, utils.SHA256HashString(obj.Password), 1).Count(&localAuth).Error;err!=nil {
		//fmt.Println("gorm.IsRecordNotFoundError(err) localAuth", gorm.IsRecordNotFoundError(err))
		response.Error(c, err)
		return
	}
	//fmt.Println("localAuth:", localAuth)
	if localAuth > 0 {
		returnResult(c, obj.Username, 1)
		return
	}

	var ldap = NewLDAP()
	ldapAuthErr := ldap.Bind(obj.Username, obj.Password)
	//fmt.Println("ldapAuth:", ldapAuth)
	if ldapAuthErr ==  nil {
		returnResult(c, obj.Username, 2)
		return
	} else {
		//logger.Errorf("ldap auth err: %s", ldapAuthErr)
		logger.Errorf(fmt.Sprintf("[%s]认证失败，用户名或密码不对,请重新输入！", obj.Username))
		//response.Error(c, fmt.Errorf("[%s]认证失败，用户名或密码不对,请重新输入！", obj.Username))
		response.Error(c, fmt.Errorf("%s", ldapAuthErr))
		return
	}

}

// @Tags 登录认证
// @Summary 刷新Token
// @Description 刷新Token
// @Produce  json
// @Param payload body  RefreshParams true "参数信息"
// @Success 200 object response.Data {"code": 2000, "status": "ok", "message": "success", "data": ""}
// @Failure 400 object response.Data {"code": 4001, "status": "error", "message": "error", "data": ""}
// @Router /api/v1/refresh-token [post]
func RefreshToken(c *gin.Context)  {
	var obj RefreshParams
	if err:= c.ShouldBindJSON(&obj);err!=nil{
		response.Error(c, err)
		return
	}
	//fmt.Println("obj.token:", obj.Token)
	tokenApi := goJWT.New()
	tokenApi.Options.SigningKey = []byte(config.Configs.JWT.RefreshKey)
	claims, err := tokenApi.ParseToken(obj.Token,config.Configs.JWT.RefreshKey)
	if err !=nil{
		response.Error(c, err)
		return
	}
	//fmt.Println("obj.claims:", claims)

	tokenApi.SetClaims(claims)
	data,err := TokenMethod(tokenApi, claims)
	if err != nil {
		response.Error(c, err, map[string]interface{}{})
		return
	}
	data["roles"] = claims["roles"]
	logger.Infof("%s(%s) token刷新成功!", data["nickname"],data["username"])
	response.Success(c, data)
	return
}

func TokenMethod(tokenApi goJWT.JWTAuth, cla jwt.MapClaims) (map[string]interface{}, error) {
	tokenApi.Options.TokenType = config.Configs.JWT.TokenType //"Bearer"
	tokenApi.Options.SigningKey = []byte(config.Configs.JWT.TokenKey)
	tokenApi.Options.Expired = config.Configs.JWT.Age
	tokenApi.Options.Claims = cla
	tokenString, err := tokenApi.GenerateToken()
	//claims, aa := goJWT.ParseToken(tokenString, "")
	if err != nil {
		return nil, err
	}
	now := time.Now()
	expiresAt := time.Now().Add(time.Duration(tokenApi.Options.Expired) * time.Second).Unix()
	expireTime := now.Add(time.Second * time.Duration(tokenApi.Options.Expired)).Format("2006-01-02 15:04:05")
	tokenApi.Options.SigningKey = []byte(config.Configs.JWT.RefreshKey)
	tokenApi.Options.Expired = config.Configs.JWT.Age + 3600
	RefreshTokenString, err := tokenApi.GenerateToken()
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{
		"token":         tokenString.GetAccessToken(),
		"refresh_token": RefreshTokenString.GetAccessToken(),
		"expire_time":   expireTime,
		"expires_at":    expiresAt,
		"user":          map[string]interface{}{},
		"user_type":     "local",
		"token_type":    tokenString.GetTokenType(),   // tokenString.GetTokenType()
		"nickname":      cla["nickname"],
		"username":      cla["username"],
		"email":         cla["email"],
	}
	return data, nil
}

func returnResult(c *gin.Context, username string, userType int)  {
	type User struct {
		ID        uint   `json:"id"`
		Nickname  string `json:"nickname"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Status    bool   `json:"status"`
	}
	var user User
	if err:=db.DB.Table("user").Where("username = ? and user_type_id = ?", username,
		userType).Select("id, nickname, username, email, status").First(&user).Error;err!=nil{
		//fmt.Println("returnResult, gorm.IsRecordNotFoundError(err)", gorm.IsRecordNotFoundError(err))
		response.Error(c, err)
		return
	}
	if !user.Status{
		response.Error(c, fmt.Errorf("用户名%s已经被禁用", username))
		return
	}
	//fmt.Printf("user: %+v", user)
	var role role2.IRole = &role2.Role{}
	roleList, err:=role.GetUserRoles(username)
	if err != nil {
		response.Error(c, err)
		return
	}
	//fmt.Printf("roleList: %+v", roleList)
	var u = sys.InfoUser{ID:user.ID,Name:user.Nickname,Nickname:user.Nickname,Username:user.Username,
		Email:user.Email,Roles:roleList}
	//var cla = map[string]interface{}{"role_list": roleList, "username": username, "name": user.Nickname, "nickname": user.Nickname}
	var cla map[string]interface{}
	if err := convert.StructToMapOut(u, &cla); err!=nil{
		log.Print("user info parse claims err:", err)
		return
	}
	tokenApi := goJWT.New()
	tokenApi.SetClaims(cla)
	//fmt.Println("cla:", cla)
	//err := middleware.NewSession(c, cla)
	//fmt.Println("cla:", err)
	//if err != nil {
	//	response.Error(c, err, map[string]interface{}{})
	//	return
	//}
	data,err := TokenMethod(tokenApi, cla)
	if err != nil {
		response.Error(c, err, map[string]interface{}{})
		return
	}
	data["roles"] = roleList
	logger.Infof("%s(%s) 登陆成功!", user.Nickname,user.Username)
	response.Success(c, data)
	return
}

func LDAPApi()  {

}

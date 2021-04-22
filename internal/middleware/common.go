package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"jobor/internal/config"
	"jobor/pkg/jwt"
	"time"
)

func TokenData(c *gin.Context, headerTag string) (map[string]interface{}, error) {
	hToken := c.Request.Header.Get(headerTag)
	//fmt.Println("token:", hToken)
	//fmt.Println("Header:", c.Request.Header)
	//fmt.Println("config.Configs.JWT.TokenType:", config.Configs.JWT.TokenType)
	var token string
	ticketLen := len(config.Configs.JWT.TokenType+ " ") // Authorization AUTHORIZATION
	if ticketLen < len(hToken){
		token = hToken[ticketLen:]
	}else {
		return nil, errors.New("token is none")
	}
	//fmt.Println("token:", token)
	var j = jwt.New()
	MapCla, err := j.ParseToken(token, "")
	if err != nil {
		return nil, err
	}
	return MapCla, nil
}

func SessionData(c *gin.Context) (map[string]interface{}, error) {
	fmt.Println("header:", c.Request.Header.Get("Cookie"))
	session := sessions.Default(c)
	//session.Options(sessions.Options{MaxAge: 60}) // 以秒为单位
	//session.Set("is_login", true)
	//session.Set("role_list", []string{"aaa", "ddd", "mmm"})
	//_ = session.Save()

	isLogin := session.Get("is_login")
	name := session.Get("name")
	username := session.Get("username")
	email := session.Get("email")
	roleList := session.Get("role_list")
	loginTime := session.Get("login_time")
	ip := session.Get("ip")
    d := map[string]interface{}{}
    d["is_login"] = isLogin
    d["name"] = name
    d["nickname"] = name
    d["username"] = username
    d["email"] = email
    d["role_list"] = roleList
    d["login_time"] = loginTime
    d["ip"] = ip
	//fmt.Println("SessionData:", name, username, email, "isLogin:", isLogin, reflect.TypeOf(isLogin), roleList, reflect.TypeOf(roleList))
	if isLogin == false || isLogin==nil {
		return nil, errors.New("session is none")
	}
	return d, nil
}

func NewSession(c *gin.Context, userData map[string]interface{}) error {
	session := sessions.Default(c)
	//fmt.Println("ccc:", c, userData)
	//session := sessions.DefaultMany(c, "aa")
	session.Options(sessions.Options{MaxAge:300, Secure:false})
	session.Set("is_login", true)
	session.Set("name", userData["name"])
	session.Set("nickname", userData["nickname"])
	session.Set("username", userData["username"])
	session.Set("email", userData["email"])
	session.Set("role_list", userData["role_list"])
	session.Set("ip", c.ClientIP())
	session.Set("login_time", fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")))
	err := session.Save()
	return err
}

func CookieData()  {
	
}

func NewCookie()  {

}
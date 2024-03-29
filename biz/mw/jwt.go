package mw

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"jobor/biz/dal/db"
	"jobor/biz/dal/ldap"
	"jobor/biz/dal/redisStore"
	"jobor/biz/model"
	"jobor/biz/pack/oidc_callback"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/user"
	"jobor/pkg/utils"
	"log"
	"net/http"
	"time"

	sessions2 "github.com/gorilla/sessions"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var AuthWm *jwt.HertzJWTMiddleware

func InitJwt() {
	AuthWm = Jwt(conf.GetConf().JWT.Key)
}

type login2 struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}

var identityKey = "id"

const (
	IsLogin = "islogin"
)

func PingHandler(c context.Context, ctx *app.RequestContext) {
	//user, _ := ctx.Get(identityKey)
	//ctx.JSON(200, utils.H{
	//	"message": fmt.Sprintf("username:%v", (*User).Username),
	//})
}

// User sso
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type LoginResp struct {
	Id       int64    `json:"id"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Token    string   `json:"token"`
	Expire   string   `json:"expire"`
	Roles    []string `json:"roles"`
}

func Jwt(key string) *jwt.HertzJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "jwt",
		Key:     []byte(key),
		Timeout: time.Hour, // 用于设置 token 过期时间，默认为一小时
		// 用于设置最大 token 刷新时间，允许客户端在 TokenTime + MaxRefresh 内刷新 token 的有效时间，追加一个 Timeout 的时长
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		//TokenLookup:   "", // 用于设置 token 的获取源，可以选择 header、query、cookie、param、form，默认为 header:Authorization
		//TokenHeadName: "", // 用于设置从 header 中获取 token 时的前缀，默认为 Bearer
		// 用于设置登陆成功后为向 token 中添加自定义负载信息的函数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			var jmc jwt.MapClaims
			if err := utils.AnyToAny(data, &jmc); err == nil {
				return jmc
			}
			hlog.Errorf("Jwt PayloadFunc 解析失败：%v", data)
			return jwt.MapClaims{}
		},
		// 用于设置获取身份信息的函数，默认与 IdentityKey 配合使用
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &user.Userinfo{
				Id: claims[identityKey].(int64),
			}
		},
		// 用于设置登录时认证用户信息的函数（必要配置）
		Authenticator: Authenticator,
		// 用于设置授权已认证的用户路由访问权限的函数
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			//if v, ok := data.(*User); ok && v.Username == "admin" {
			//	return true
			//}
			//return false
			return true
		},
		// 用于设置 jwt 验证流程失败的响应函数
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			//response.AuthFailed(ctx, c, fmt.Errorf(message), "")
			response.SendBaseResp(ctx, c, response.AuthenticateFailed)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			userInter, ok := c.Get(HeaderTag)
			u, ook := userInter.(user.Userinfo)
			if ok && ook {
				lr := LoginResp{Nickname: u.Nickname, Username: u.Username,
					Id: u.Id, Token: token, Expire: expire.String(), Roles: u.Roles}
				hlog.Infof("user %s login is success", u.Username)
				response.SendDataResp(ctx, c, response.Succeed, &lr)
			} else {
				//response.AuthFailed(ctx, c, fmt.Errorf("获取用户信息错误"), "")
				response.SendBaseResp(ctx, c, response.ServerInternalErr)
			}
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			PostLogout(ctx, c)
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			userInter, ok := c.Get(HeaderTag)
			var u user.Userinfo
			var err = utils.AnyToAny(userInter, &u)
			if ok && err == nil {
				lr := LoginResp{Nickname: u.Nickname, Username: u.Username,
					Id: u.Id, Token: token, Expire: expire.String(), Roles: u.Roles}
				hlog.Infof("user %s token refresh is success", u.Username)
				response.SendDataResp(ctx, c, response.Succeed, &lr)
			} else {
				hlog.Errorf("get header Userinfo is %v, err is %v", ok, err)
				//response.AuthFailed(ctx, c, fmt.Errorf("获取用户信息错误"), "")
				response.SendBaseResp(ctx, c, response.ServerInternalErr)
			}
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit Error:" + errInit.Error())
	}
	return authMiddleware
}

func Authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	hlog.CtxInfof(ctx, "authenticator is start")
	var err error
	var req LoginReq
	if err = c.BindAndValidate(&req); err != nil {
		return nil, err
	}
	u, err := MultipleSrcAuth(ctx, db.DB, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	model.SetContentUserinfo(ctx, c, u)
	if conf.GetConf().Authentication.EnableSession {
		session := sessions.Default(c)
		if err = model.InitSession(ctx, session, u, conf.GetConf().SSO.ClientId, SessionOrg, true); err != nil {
			return nil, err
		}
	}
	return &u, nil
	//return nil, jwt.ErrFailedAuthentication
}

const SessionOrg = conf.Dom

func SetSessionLogin(ctx context.Context, session sessions.Session, isLogin bool) error {
	session.Set(IsLogin, isLogin)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func RefreshSessionExpire(ctx context.Context, session sessions.Session) error {
	session.Set("timestamp", time.Now().Unix())
	if err := session.Save(); err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return err
	}
	return nil
}

func SetSessionLoginBySessionId(sessionId string) error {
	rs, err := redisStore.GetRedisStore()
	if err != nil {
		return err
	}
	session, err := redis.LoadSessionBySessionId(rs, sessionId)
	if err != nil {
		return err
	}
	session.Options = &sessions2.Options{
		Path:   "/",
		MaxAge: conf.GetConf().Authentication.MaxAge,
	}
	session.Values[IsLogin] = true
	err = redis.SaveSessionWithoutContext(rs, sessionId, session)
	return err
}

type LoginReq struct {
	Username string `form:"username,required" default:"" json:"username,required"`
	Password string `form:"password,required" default:"" json:"password,required"`
	//CaptchaId *string `form:"captcha_id" default:"" json:"captcha_id"`
	//Captcha   *string `form:"captcha" default:"" json:"captcha"`
	//MfaCode   *string `form:"mfa_code" default:"" json:"mfa_code"`
}
type LogoutReq struct {
	PostLogoutRedirectUri string `form:"post_logout_redirect_uri" default:"" json:"post_logout_redirect_uri"`
	IdTokenHint           string `form:"id_token_hint" default:"" json:"id_token_hint"`
}

// PostLogin .
//
//	@Summary		user login summary
//	@Description	user login
//	@Tags			login
//	@Accept			application/json
//	@Produce		application/json
//	@Param			json	body	LoginReq	true	"参数"
//	@router			/api/v1/login [POST]
func PostLogin(ctx context.Context, c *app.RequestContext) {

}

// PostLogout .
//
//	@Summary		user logout summary
//	@Description	user logout
//	@Tags			login
//	@Accept			application/json
//	@Produce		application/json
//	@Param			json	body	LogoutReq	true	"参数"
//	@router			/api/v1/logout [POST]
func PostLogout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req LogoutReq
	if err = c.BindAndValidate(&req); err != nil {
		response.ParamFailed(ctx, c, err)
		return
	}
	if conf.GetConf().Authentication.EnableSession {
		session := sessions.Default(c)
		session.Clear()
		session.Delete(HeaderTag)
		if err = session.Save(); err != nil {
			response.SendBaseResp(ctx, c, err)
			return
		}
		hlog.CtxDebugf(ctx, "user %s logout clear session cookie is success", c.GetString("username"))
	}
	uri := oidc_callback.GetServerDomain(c)
	oic := getOidcProvider(ctx)
	oidcf, err := oic.GetOpenIdConf()
	if err != nil {
		response.SendBaseResp(ctx, c, err)
		return
	}
	if req.PostLogoutRedirectUri != "" {
		redUri := fmt.Sprintf("%s?post_logout_redirect_uri=%s", oidcf.EndSessionEndpoint, req.PostLogoutRedirectUri)
		hlog.CtxDebugf(ctx, "redirect uri: %s", redUri)
		c.Redirect(http.StatusFound, []byte(redUri))
		c.Abort()
		//response.Success(ctx, c, redUri)
	} else if uri != "" {
		redUri := fmt.Sprintf("%s?post_logout_redirect_uri=%s",
			oidcf.EndSessionEndpoint, uri)
		hlog.CtxDebugf(ctx, "redirect uri: %s", redUri)
		c.Redirect(http.StatusFound, []byte(redUri))
		c.Abort()
		//response.Success(ctx, c, redUri)
	} else {
		response.SendDataResp(ctx, c, response.Succeed, "logout success")
	}

	//response.Success(ctx, c, "logout success")
}

// LoginInit .
//
//	@Summary		login init summary
//	@Description	login init
//	@Tags			login
//	@router			/api/v1/login-init [GET]
func LoginInit(ctx context.Context, c *app.RequestContext) {
	type Data struct {
		SsoAuth   bool   `json:"sso_auth"`
		SsoTip    string `json:"sso_tip"`
		LocalAuth bool   `json:"local_auth"`
	}
	data := Data{SsoAuth: conf.GetConf().SSO.Enable, SsoTip: conf.GetConf().SSO.Banner,
		LocalAuth: conf.GetConf().Authentication.LocalAuth}
	response.SendDataResp(ctx, c, response.Succeed, data)
}

func getOidcProvider(ctx context.Context) *oidc_callback.OIDC {
	no := oidc_callback.NewOIDC(ctx, conf.GetConf().SSO.IssuerURL, conf.GetConf().SSO.ClientId,
		conf.GetConf().SSO.ClientSecret, conf.GetConf().SSO.Scope, oidc_callback.CallbackPath)
	return no
}

func MultipleSrcAuth(ctx context.Context, DB *gorm.DB, username string, password string) (u user.Userinfo, err error) {
	if !conf.GetConf().SSO.Enable && !conf.GetConf().Authentication.LocalAuth && !conf.GetConf().Ldap.Enabled {
		return u, fmt.Errorf("必须开启至少一个认证源【本地、LDAP、SSO】")
	}
	if conf.GetConf().Authentication.LocalAuth {
		ui, localOk, err := model.Auth(DB, username, password)
		if err != nil {
			hlog.Errorf("local auth err: %s", err)
			goto LdapAuth
			//return u, err
		}
		if localOk {
			if !ui.Status {
				return u, fmt.Errorf("用户[%s]已经被禁用，请联系管理员", username)
			}
			u = *ui.GetUserinfo()
			hlog.CtxInfof(ctx, "user %s local db auth is success", username)
			u.Roles, err = model.GetUserRoles(username)
			if err != nil {
				hlog.Errorf("local auth err: %s", err)
				return u, err
			} else {
				return u, nil
			}
		}
	}
LdapAuth:
	if conf.GetConf().Ldap.Enabled {
		var ldapOk bool
		var li = conf.GetConf().Ldap
		ldapApi, err := ldap.GetLdapApi(li.Enabled, li.Tls, li.Addr, li.BaseDN,
			li.Username, li.Password, li.AuthFilter, nil)
		if err != nil {
			hlog.Debugf("user %s ldap/ad auth err, %s", username, err)
			goto SSOAuth
		}
		r, err := ldapApi.Authentication(username, password)
		if err == nil {
			var u = user.Userinfo{Username: r.Username, Nickname: r.DisplayName,
				UserType: model.Ldap, Phone: r.Phone, Email: r.Email, Status: true}
			u, err := model.GetUserinfoOrCreate(&u)
			if err != nil {
				hlog.Errorf("ldap auth create or get err: %s", err)
				return u, err
			} else if u.Username == "" {
				err = fmt.Errorf("用户[%s]信息不存在", username)
				hlog.Errorf("ldap auth err: %s", err)
				return u, err
			}
			if !u.Status {
				err = fmt.Errorf("用户[%s]已经被禁用，请联系管理员", username)
				hlog.Errorf("ldap auth err: %s", err)
				return u, err
			}
			ldapOk = true
			hlog.CtxInfof(ctx, "user %s ldap auth is success", username)
			u.Roles, err = model.GetUserRoles(username)
			if err != nil {
				hlog.Errorf("ldap auth get roles err: %s", err)
				return u, err
			}
		} else {
			hlog.Errorf("ldap auth err: %s", err)
		}
		hlog.CtxDebugf(ctx, "user %s ldap/ad auth is %v", username, ldapOk)
	}
SSOAuth:
	if conf.GetConf().SSO.Enable {
		no := getOidcProvider(ctx)
		userInfo, err := no.OidcCallbackWithPassword(ctx, username, password)
		if err != nil {
			return u, err
		}

		var body json.RawMessage
		var u user.Userinfo
		if err = userInfo.Claims(&body); err != nil {
			return u, err
		}
		hlog.CtxDebugf(ctx, "get userInfo %s from oidc server", string(body))
		u = model.GetUserinfoFromOidc(body)
		u, err = model.GetUserinfoOrCreate(&u)
		if err != nil {
			hlog.CtxErrorf(ctx, "claims unmarshals the raw JSON object claims into the provided object err, %s", err)
			return u, err
		}
	}
	err = fmt.Errorf("[%s]认证失败，用户名或密码不对,请重新输入", username)
	return u, err
}

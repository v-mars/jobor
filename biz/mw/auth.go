package mw

import (
	"context"
	"jobor/biz/dal/redis"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/pkg/convert"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"github.com/hertz-contrib/sessions"
)

var (
	HeaderAuthorization = "Authorization"
	HeaderTag           = "userinfo"
)

// UserAuthMw 用户授权中间件
func UserAuthMw(skipper ...SkipperFunc) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//path := c.Request.URL.Path
		//fmt.Println("UserAuthMiddleware:", path)
		//fmt.Println("skipper:", skipper, len(skipper), skipper[0](ctx, c))

		// 白名单
		if len(skipper) > 0 && skipper[0](ctx, c) {
			c.Next(ctx)
			return
		}
		userInfo := map[string]interface{}{}
		var ok bool
		var err error
		session := sessions.Default(c)
		if conf.GetConf().Authentication.EnableSession {
			//isLogin, isLoginOk = session.Get(IsLogin).(bool)
			userInfo, ok = session.Get(HeaderTag).(map[string]interface{})
			//fmt.Println("session:", session.Get("uid"), userInfo)
			hlog.CtxDebugf(ctx, "session auth is %v, session_id: %s", ok, session.ID())
		}
		if !ok {
			hlog.CtxDebugf(ctx, "jwt auth is start")
			userInfo, err = AuthWm.GetClaimsFromJWT(ctx, c)
		} else if conf.GetConf().Authentication.EnableSession && ok {
			// 刷新session过期时间
			_, err = redis.Rdb.Expire(ctx, "session_"+session.ID(),
				time.Duration(conf.GetConf().Authentication.MaxAge)*time.Second).Result()
			if err != nil {
				hlog.CtxErrorf(ctx, err.Error())
				response.SendBaseResp(ctx, c, response.AuthenticateFailed)
				c.Abort()
				return
			}
			// 刷新cookie过期时间
			session.Set("timestamp", time.Now().Unix())
			if err = session.Save(); err != nil {
				hlog.CtxErrorf(ctx, err.Error())
				response.SendBaseResp(ctx, c, response.ServerInternalErr)
				c.Abort()
				return
			}
		}
		if err != nil {
			//var currentUri = c.Request.URI().RequestURI()
			//var currentUri = fmt.Sprintf("%s://%s%s", string(c.Request.Scheme()), string(c.Request.Host()),
			//	c.Request.URI().RequestURI())
			//var state = uuid.New()
			//var nonce = uuid.New()
			//redirectSSOUri := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=%s&state=%s&nonce=%s",
			//	conf.GetConf().SSO.SsoRedirectUri, conf.GetConf().SSO.ClientId, currentUri, "code", state, nonce)
			//c.Redirect(http.StatusSeeOther, []byte(redirectSSOUri))
			//c.Abort()
			if err == jwt.ErrEmptyAuthHeader {
				response.SendBaseResp(ctx, c, response.Unauthenticated)
			} else {
				response.SendBaseResp(ctx, c, response.AuthenticateFailed)
			}
			c.Abort()
			return
		} else {
			c.Set(HeaderTag, userInfo)
			c.Set("uid", convert.ToInt64(userInfo["id"]))
			c.Set("nickname", userInfo["nickname"])
			c.Set("username", userInfo["username"])
			c.Set("email", userInfo["email"])
			c.Set("roles", userInfo["roles"])
			c.Set("userid", userInfo["userid"])
			c.Set("dom", conf.Dom)
			c.Next(ctx)
			return
		}
	}
}

func TokenData(ctx context.Context, c *app.RequestContext, t string) (map[string]string, error) {
	return nil, nil
}

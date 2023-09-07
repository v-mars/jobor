package oidc_callback

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hertz-contrib/sessions"
	"jobor/biz/response"
	"jobor/conf"
	"jobor/kitex_gen/user"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type RedirectResp struct {
	ClientId            string `json:"client_id"`
	RedirectUri         string `json:"redirect_uri"`
	ResponseType        string `json:"response_type"`
	State               string `json:"state"`
	Nonce               string `json:"nonce"`
	Scope               string `json:"scope"`
	CodeChallengeMethod string `json:"code_challenge_method"` //S256
	CodeChallenge       string `json:"code_challenge"`
}

var (
	AuthUrl      = "/api/v1/oauth/authorize"
	CallbackPath = "/api/v1/jobor/oidc/callback"
	GotoRedirect = "/api/v1/jobor/oidc/redirect"
)

type OIDC struct {
	Provider     *oidc.Provider
	Issuer       string
	ClientId     string
	ClientSecret string
	CallbackPath string
	Scope        string
	AuthURL      string
	TokenURL     string
}

func NewOIDC(ctx context.Context, issuer, clientId, clientSecret, scope, callbackPath string) *OIDC {
	ctx = oidc.InsecureIssuerURLContext(ctx, issuer)
	pro, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic(err)
	}
	return &OIDC{Provider: pro, Issuer: issuer, ClientSecret: clientSecret,
		ClientId: clientId, CallbackPath: callbackPath, Scope: scope}
}

func InitProvider(ctx context.Context, issuer string) (*oidc.Provider, error) {
	var err error
	ctx = oidc.InsecureIssuerURLContext(ctx, issuer)
	providerObj, err := oidc.NewProvider(ctx, issuer)
	//ctx = oidc.InsecureIssuerURLContext(ctx, "http://127.0.0.1:5610")
	//providerObj, err := oidc.NewProvider(ctx, issuer)
	return providerObj, err
}

func (o *OIDC) InitProvider(ctx context.Context) (*oidc.Provider, error) {
	var err error
	o.Provider, err = InitProvider(ctx, o.Issuer)
	return o.Provider, err
}

// OidcCallbackWithCode sso code类型 客户端回调接口方法
func (o *OIDC) OidcCallbackWithCode(ctx context.Context, redirectURL, code string) (*oidc.UserInfo, error) {
	var err error
	o.Provider, err = o.InitProvider(ctx)
	if err != nil {
		return nil, err
	}

	hlog.CtxDebugf(ctx, "get oidc provider is success")
	// 1. 配置OAuth 2.0客户端
	oauth2Config := &oauth2.Config{
		ClientID:     o.ClientId,
		ClientSecret: o.ClientSecret,
		RedirectURL:  o.CallbackPath,
		Endpoint:     o.Provider.Endpoint(),
		Scopes:       []string{o.Scope},
	}
	if o.TokenURL != "" {
		oauth2Config.Endpoint.TokenURL = o.TokenURL
	}
	if o.AuthURL != "" {
		oauth2Config.Endpoint.AuthURL = o.AuthURL
	}

	// 2. 获取授权码
	//authURL := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	//oauth2Config.AuthCodeURL(state)
	oauth2Token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		hlog.CtxErrorf(ctx, "get oauth2 access_token id_token err, %s", err)
		return nil, err
	}
	hlog.CtxDebugf(ctx, "oidc client auth code is success")
	// Extract the Id Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		hlog.CtxErrorf(ctx, "handle missing id token")
		return nil, fmt.Errorf("handle missing id token")
	}

	var verifier = o.Provider.Verifier(&oidc.Config{ClientID: oauth2Config.ClientID})
	// Parse and verify Id Token payload.
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		//fmt.Println("idToken:", idToken)
		_ = idToken
		hlog.CtxErrorf(ctx, "verifier id_token err, %s", err)
		return nil, err
	}
	userInfo, err := o.Provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		hlog.CtxErrorf(ctx, "%s", err.Error())
		data, er := base64.StdEncoding.DecodeString(err.Error())
		if er != nil {
			return nil, err
		}
		return nil, fmt.Errorf(string(data))
	}
	hlog.CtxDebugf(ctx, "oidc client get userinfo is success")
	return userInfo, nil
}

// OidcCallbackWithPassword sso password类型 客户端调用接口方法
func (o *OIDC) OidcCallbackWithPassword(ctx context.Context, username, password string) (*oidc.UserInfo, error) {
	var err error
	o.Provider, err = o.InitProvider(ctx)
	if err != nil {
		return nil, err
	}

	hlog.CtxDebugf(ctx, "get oidc provider is success")
	// 1. 配置OAuth 2.0客户端
	oauth2Config := &oauth2.Config{
		ClientID:     o.ClientId,
		ClientSecret: o.ClientSecret,
		RedirectURL:  o.CallbackPath,
		Endpoint:     o.Provider.Endpoint(),
		Scopes:       strings.Fields(o.Scope),
	}
	if o.TokenURL != "" {
		oauth2Config.Endpoint.TokenURL = o.TokenURL
	}
	if o.AuthURL != "" {
		oauth2Config.Endpoint.AuthURL = o.AuthURL
	}

	// 2. 获取授权码
	//authURL := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	//oauth2Config.AuthCodeURL(state)
	oauth2Token, err := oauth2Config.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		hlog.CtxErrorf(ctx, "get oauth2 access_token id_token err, %s", err)
		return nil, err
	}
	hlog.CtxDebugf(ctx, "oidc client auth username and password is success")

	// Extract the Id Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		hlog.CtxErrorf(ctx, "handle missing id token")
		return nil, fmt.Errorf("handle missing id token")
	}

	var verifier = o.Provider.Verifier(&oidc.Config{ClientID: oauth2Config.ClientID})
	// Parse and verify Id Token payload.
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		fmt.Println("idToken:", idToken)
		hlog.CtxErrorf(ctx, "verifier id_token err, %s", err)
		return nil, err
	}
	userInfo, err := o.Provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		fmt.Println("err.Error():", err.Error())
		data, er := base64.StdEncoding.DecodeString(err.Error())
		if er != nil {
			return nil, err
		}
		return nil, fmt.Errorf(string(data))
	}
	hlog.CtxDebugf(ctx, "oidc client get userinfo is success")
	return userInfo, nil
}

func (o *OIDC) GetOpenIdConf() (OpenIdConf, error) {
	//o.SetHttpClient(nil)
	var oic OpenIdConf
	err := o.Provider.Claims(&oic)
	return oic, err
}

// GetCallbackUri 获取本机url
func GetCallbackUri(c *app.RequestContext, withIdp string) string {
	var xSchema = string(c.GetHeader(conf.SchemeHeaderValue))
	var schema = string(c.Request.Scheme())
	if len(xSchema) > 0 {
		schema = xSchema
	}
	hlog.Debugf("X-Scheme: %s", xSchema)
	path := CallbackPath
	if len(withIdp) > 0 {
		path = fmt.Sprintf("%s/%s", path, withIdp)
	}
	host := fmt.Sprintf("%s://%s%s", schema, string(c.Request.Host()), path)
	return host
}

// GetIssuer 获取本机url
func GetIssuer(c *app.RequestContext) string {
	var xSchema = string(c.GetHeader("X-Uac-Scheme"))
	hlog.Debugf("X-Uac-Scheme Header: %s", xSchema)
	var schema = string(c.Request.Scheme())
	if len(xSchema) > 0 {
		schema = xSchema
	}
	issuer := fmt.Sprintf("%s://%s", schema, string(c.Request.Host()))
	hlog.Debugf("X-Uac-Scheme: %s, Issuer: %s", xSchema, issuer)
	return issuer
}

func GetServerDomain(c *app.RequestContext) string {
	var schema = string(c.Request.Scheme())
	uri := fmt.Sprintf("%s://%s", schema, string(c.Request.Host()))
	hlog.Debugf("get domain uri, uri: %s", uri)
	return uri
}

type OpenIdConf struct {
	Issuer                string   `json:"issuer"`
	AuthURL               string   `json:"authorization_endpoint"`
	TokenURL              string   `json:"token_endpoint"`
	JWKSURL               string   `json:"jwks_uri"`
	UserInfoURL           string   `json:"userinfo_endpoint"`
	EndSessionEndpoint    string   `json:"end_session_endpoint"`
	IntrospectionEndpoint string   `json:"introspection_endpoint"`
	Algorithms            []string `json:"id_token_signing_alg_values_supported"`
}

// ################### 回调接口和提供前端重定向接口 #################
//r.GET("/redirect", oidc_callback.RedirectLogin)
//r.GET(oidc_callback.CallbackPath, oidc_callback.SsoCallback)

// SsoCallback example
// AccessToken .
//
//	@Summary		oidc callback summary
//	@Description	oidc callback
//	@Tags			login
//	@Param			code	query	string	true	"code"
//	@Param			state	query	string	false	"state"
//	@router			/api/v1/jobor/oidc/callback [GET]
func SsoCallback(ctx context.Context, c *app.RequestContext) {
	var err error
	code, ok := c.GetQuery("code")
	state, _ := c.GetQuery("state")
	redirectURLStr, ok2 := c.GetQuery("redirect_uri")
	if !ok {
		err = fmt.Errorf("callback oidc client request missing code args")
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
		return
	}
	if !ok2 {
		redirectURLStr = "/"
	}
	se := sessions.Default(c)
	cacheState, _ := se.Get("state").(string)
	if state != "" && state != cacheState {
		err = fmt.Errorf("callback oidc client request/local missing state args")
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
		return
	}
	oc := NewOIDC(ctx, conf.GetConf().SSO.IssuerURL, conf.GetConf().SSO.ClientId,
		conf.GetConf().SSO.ClientSecret, conf.GetConf().SSO.Scope, GetCallbackUri(c, ""))
	info, err := oc.OidcCallbackWithCode(ctx, GetCallbackUri(c, ""), code)
	if err != nil {
		hlog.CtxErrorf(ctx, "%s", err)
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
		return
	}
	var claims json.RawMessage
	err = info.Claims(&claims)
	if err != nil {
		hlog.CtxErrorf(ctx, "claims unmarshals the raw JSON object claims into the provided object err, %s", err)
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
		return
	}
	//marshal, _ := json.Marshal(info)
	hlog.CtxDebugf(ctx, "oidc callback userinfo:", string(claims))
	u := user.Userinfo{}
	//u.Username = gjson.Parse(string(claims)).Get("claims.username").String()
	//u.Nickname = gjson.Parse(string(claims)).Get("claims.nickname").String()
	u = user.GetUserinfoFromOidc(claims)
	if u.Username == "" || u.Nickname == "" {
		er := fmt.Errorf("username or nickname is null")
		hlog.CtxErrorf(ctx, er.Error())
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, er)
		return
	}
	u, err = user.GetUserinfoOrCreate(&u)
	if err != nil {
		hlog.CtxErrorf(ctx, "claims unmarshals the raw JSON object claims into the provided object err, %s", err)
		response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
		return
	}
	user.SetContentUserinfo(ctx, c, u)
	if conf.GetConf().Authentication.EnableSession {
		if err = user.InitSession(ctx, se, u, conf.GetConf().SSO.ClientId, conf.Dom, true); err != nil {
			hlog.CtxErrorf(ctx, err.Error())
			response.FailedHttpCode(ctx, c, http.StatusBadRequest, err)
			return
		}

		//fmt.Println("userInfo cc:", cc)
		c.Redirect(http.StatusFound, []byte(redirectURLStr))
		c.Abort()
	}
}

// RedirectLogin 重定向到sso登录页面
//
//	@Summary		oidc redirect login summary
//	@Description	oidc redirect login
//	@Tags			login
//	@Param			next	query	string	false	"next"
//	@router			/api/v1/jobor/oidc/redirect [GET]
func RedirectLogin(ctx context.Context, c *app.RequestContext) {
	//code, ok := c.GetQuery("code")
	next, ok := c.GetQuery("next")
	if ok {
		var state = uuid.New().String()
		var nonce = uuid.New().String()
		v := url.Values{}
		v.Add("client_id", conf.GetConf().SSO.ClientId)
		v.Add("redirect_uri", next)
		v.Add("response_type", "code")
		v.Add("state", state)
		v.Add("nonce", nonce)
		v.Add("scope", conf.GetConf().SSO.Scope)
		v.Add("code_challenge_method", "S256")
		v.Add("code_challenge", uuid.New().String())
		providerObj, err := InitProvider(ctx, conf.GetConf().SSO.IssuerURL)
		if err != nil {
			response.SendBaseResp(ctx, c, fmt.Errorf("获取provider错误, %s", err))
			return
		}
		se := sessions.Default(c)
		se.Set("state", state)
		se.Set("nonce", nonce)
		se.Set("redirect_uri", next)
		err = se.Save()
		if err != nil {
			response.SendBaseResp(ctx, c, err)
			return
		}
		authURL := providerObj.Endpoint().AuthURL
		redirectSSOUri := fmt.Sprintf("%s?%s", authURL, v.Encode())
		c.Redirect(http.StatusFound, []byte(redirectSSOUri))
	} else {
		response.SendBaseResp(ctx, c, fmt.Errorf("next重新定向Url参数不存在"))
	}

}
